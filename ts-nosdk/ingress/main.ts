// This example demonstrates how to use the Namespace Compute API without the
// @namespacelabs/sdk package, using only fetch and the Connect protocol (JSON).
//
// It creates an nginx container on the host network listening on port 8080,
// then exposes it via CreateIngress and outputs the public URL.
//
// WARNING: This example loads a bearer token directly from NSC_TOKEN_FILE.
// This is NOT recommended for production use. Bearer tokens expire and need to
// be refreshed. Use the @namespacelabs/sdk package instead, which handles token
// lifecycle (session tokens, caching, refresh) automatically.

import * as fs from "fs/promises";

const REGION = "us";
const COMPUTE_API = `https://${REGION}.compute.namespaceapis.com`;
const SERVICE = "namespace.cloud.compute.v1beta.ComputeService";

const deadlineMinutes = parseInt(process.argv[2] || "5", 10);

interface TokenJson {
	bearer_token: string;
	session_token?: string;
}

async function loadBearerToken(): Promise<string> {
	const tokenFile = process.env.NSC_TOKEN_FILE;
	if (!tokenFile) {
		throw new Error(
			"NSC_TOKEN_FILE environment variable is not set. " +
				"Point it to a token.json file (e.g. from `nsc login`)."
		);
	}

	const content = await fs.readFile(tokenFile, "utf8");
	const tokenJson: TokenJson = JSON.parse(content);
	if (!tokenJson.bearer_token) {
		throw new Error("Token file does not contain a bearer_token");
	}
	return tokenJson.bearer_token;
}

// Make a Connect (JSON) unary RPC call using fetch.
async function connectCall(
	baseUrl: string,
	service: string,
	method: string,
	token: string,
	body: unknown
): Promise<unknown> {
	const url = `${baseUrl}/${service}/${method}`;
	const res = await fetch(url, {
		method: "POST",
		headers: {
			"Content-Type": "application/json",
			Authorization: `Bearer ${token}`,
		},
		body: JSON.stringify(body),
	});

	if (!res.ok) {
		const text = await res.text();
		throw new Error(`${method} failed (${res.status}): ${text}`);
	}

	return await res.json();
}

void main();

async function main() {
	const token = await loadBearerToken();

	// Create an instance with an nginx container on the host network, listening on port 8080.
	const createResp = (await connectCall(
		COMPUTE_API,
		SERVICE,
		"CreateInstance",
		token,
		{
			shape: { virtualCpu: 2, memoryMegabytes: 4096, machineArch: "amd64" },
			documentedPurpose: "ingress example (no-sdk)",
			deadline: new Date(
				Date.now() + deadlineMinutes * 60 * 1000
			).toISOString(),
			containers: [
				{
					name: "nginx",
					imageRef: "nginx",
					args: [
						"sh",
						"-c",
						`echo 'server { listen 8080; location / { default_type text/plain; return 200 "hello from nginx\\n"; } }' > /etc/nginx/conf.d/default.conf && nginx -g 'daemon off;'`,
					],
					network: "HOST",
				},
			],
		}
	)) as {
		instanceUrl: string;
		metadata: { instanceId: string };
	};

	const instanceId = createResp.metadata.instanceId;
	console.error(`[namespace] Instance: ${createResp.instanceUrl}`);
	console.error(`[Waiting until instance becomes ready]`);

	// WaitInstanceSync is a unary call that blocks until the instance is ready.
	await connectCall(COMPUTE_API, SERVICE, "WaitInstanceSync", token, {
		instanceId,
	});

	// Expose port 8080 via public ingress.
	const ingressResp = (await connectCall(
		COMPUTE_API,
		SERVICE,
		"CreateIngress",
		token,
		{
			instanceId,
			ingresses: [
				{
					name: "nginx",
					exportedPortBackend: { port: 8080 },
				},
			],
		}
	)) as {
		allocatedIngresses: { name: string; fqdn: string }[];
	};

	for (const ingress of ingressResp.allocatedIngresses) {
		console.log(`https://${ingress.fqdn}`);
	}
}
