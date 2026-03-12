// This example demonstrates how to create a Namespace instance with an nginx
// container using only fetch and the Connect protocol (JSON).
//
// Authentication: set one of the following environment variables:
//   NSC_TOKEN      - a bearer token, e.g. from `nsc auth generate-dev-token`
//   NSC_TOKEN_FILE - path to a token.json file; set automatically inside a
//                    Namespace instance
//
// WARNING: Bearer tokens expire and need to be refreshed. This example does not
// handle token lifecycle.

import * as fs from "fs/promises";

const REGION = "us";
const COMPUTE_API = `https://${REGION}.compute.namespaceapis.com`;
const SERVICE = "namespace.cloud.compute.v1beta.ComputeService";

interface TokenJson {
	bearer_token: string;
	session_token?: string;
}

async function loadBearerToken(): Promise<string> {
	// Prefer NSC_TOKEN if set (e.g. from `nsc auth generate-dev-token`).
	const envToken = process.env.NSC_TOKEN;
	if (envToken) {
		return envToken;
	}

	// Fall back to NSC_TOKEN_FILE (set automatically inside a Namespace instance).
	const tokenFile = process.env.NSC_TOKEN_FILE;
	if (!tokenFile) {
		throw new Error(
			"Set NSC_TOKEN (e.g. from `nsc auth generate-dev-token`) " +
				"or NSC_TOKEN_FILE (path to a token.json)."
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

	// Create an instance with an nginx container.
	const createResp = (await connectCall(
		COMPUTE_API,
		SERVICE,
		"CreateInstance",
		token,
		{
			shape: { virtualCpu: 2, memoryMegabytes: 4096, machineArch: "amd64" },
			documentedPurpose: "createinstance example",
			deadline: new Date(Date.now() + 60 * 60 * 1000).toISOString(),
			containers: [
				{
					name: "nginx",
					imageRef: "nginx",
					args: [],
					exportPorts: [
						{ name: "nginx", containerPort: 80, proto: "TCP" },
					],
				},
			],
		}
	)) as {
		instanceUrl: string;
		metadata: { instanceId: string };
	};

	const instanceId = createResp.metadata.instanceId;
	console.error(`[namespace] Instance: ${createResp.instanceUrl}`);
	console.error(JSON.stringify(createResp, null, 2));

	// Wait until the instance is ready.
	const waitResp = await connectCall(
		COMPUTE_API,
		SERVICE,
		"WaitInstanceSync",
		token,
		{ instanceId }
	);

	console.error(JSON.stringify(waitResp, null, 2));
}
