// This example demonstrates how to use the Namespace Compute API without the
// @namespacelabs/sdk package, using only fetch and the Connect protocol (JSON).
//
// WARNING: This example loads a bearer token directly from NSC_TOKEN_FILE.
// This is NOT recommended for production use. Bearer tokens expire and need to
// be refreshed. Use the @namespacelabs/sdk package instead, which handles token
// lifecycle (session tokens, caching, refresh) automatically.

import * as fs from "fs/promises";

const REGION = "us";
const COMPUTE_API = `https://${REGION}.compute.namespaceapis.com`;

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

	const start = Date.now();

	// Create an instance with an ubuntu container.
	const createResp = (await connectCall(
		COMPUTE_API,
		"namespace.cloud.compute.v1beta.ComputeService",
		"CreateInstance",
		token,
		{
			shape: { virtualCpu: 2, memoryMegabytes: 4096, machineArch: "amd64" },
			documentedPurpose: "exec example (no-sdk)",
			deadline: new Date(Date.now() + 10 * 60 * 1000).toISOString(),
			containers: [
				{
					name: "ubuntu",
					imageRef: "ubuntu:latest",
					args: ["sleep", "600"],
				},
			],
		}
	)) as {
		instanceUrl: string;
		metadata: { instanceId: string };
		extendedMetadata?: { commandServiceEndpoint?: string };
	};

	const instanceId = createResp.metadata.instanceId;
	console.error(`[namespace] Instance: ${createResp.instanceUrl}`);
	console.error(JSON.stringify(createResp, null, 2));

	const endpoint = createResp.extendedMetadata?.commandServiceEndpoint;
	if (!endpoint) {
		throw new Error("command service endpoint not available");
	}

	console.error(`[namespace] Command service endpoint: ${endpoint}`);

	// Run a command in the container via the CommandService.
	const result = (await connectCall(
		endpoint,
		"namespace.cloud.compute.v1beta.CommandService",
		"RunCommandSync",
		token,
		{
			instanceId,
			targetContainerName: "ubuntu",
			command: {
				command: ["uname", "-a"],
			},
		}
	)) as {
		stdout?: string; // base64-encoded
		stderr?: string; // base64-encoded
		exitCode?: number;
	};

	const elapsed = Date.now() - start;
	console.error(
		`[namespace] Total time from CreateInstance to command result: ${elapsed}ms`
	);

	if (result.stdout) {
		process.stdout.write(Buffer.from(result.stdout, "base64"));
	}
	if (result.stderr) {
		process.stderr.write(Buffer.from(result.stderr, "base64"));
	}

	if (result.exitCode && result.exitCode !== 0) {
		throw new Error(`command exited with code ${result.exitCode}`);
	}
}
