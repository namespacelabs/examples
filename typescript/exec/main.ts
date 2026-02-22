import { createClient } from "@connectrpc/connect";
import { createConnectTransport } from "@connectrpc/connect-node";
import { Timestamp } from "@bufbuild/protobuf";
import { loadDefaults } from "@namespacelabs/sdk/auth";
import { createComputeClient } from "@namespacelabs/sdk/api/compute";
import { bearerAuthInterceptor } from "@namespacelabs/sdk/api";
import { CommandService } from "@namespacelabs/sdk/proto/namespace/cloud/compute/v1beta/command_connect";

void main();

async function main() {
	const tokenSource = await loadDefaults();

	const api = createComputeClient({ tokenSource });

	const start = Date.now();

	// Create an instance with an ubuntu container.
	const resp = await api.compute.createInstance({
		shape: { virtualCpu: 2, memoryMegabytes: 4096, machineArch: "amd64" },
		documentedPurpose: "exec example",
		deadline: Timestamp.fromDate(new Date(Date.now() + 10 * 60 * 1000)),
		containers: [
			{
				name: "ubuntu",
				imageRef: "ubuntu:latest",
				args: ["sleep", "600"],
			},
		],
	});

	const instanceId = resp.metadata!.instanceId;
	console.error(`[namespace] Instance: ${resp.instanceUrl}`);
	console.error(JSON.stringify(resp.toJson(), null, 2));

	const endpoint = resp.extendedMetadata?.commandServiceEndpoint;
	if (!endpoint) {
		throw new Error("command service endpoint not available");
	}

	console.error(`[namespace] Command service endpoint: ${endpoint}`);

	// Connect to the CommandService on the instance.
	const cmdTransport = createConnectTransport({
		httpVersion: "1.1",
		baseUrl: endpoint,
		useBinaryFormat: false,
		interceptors: [bearerAuthInterceptor(tokenSource)],
	});

	const cmdClient = createClient(CommandService, cmdTransport);

	const result = await cmdClient.runCommandSync({
		instanceId,
		targetContainerName: "ubuntu",
		command: {
			command: ["uname", "-a"],
		},
	});

	const elapsed = Date.now() - start;
	console.error(
		`[namespace] Total time from CreateInstance to command result: ${elapsed}ms`
	);

	const decoder = new TextDecoder();
	process.stdout.write(decoder.decode(result.stdout));
	if (result.stderr.length > 0) {
		process.stderr.write(decoder.decode(result.stderr));
	}

	if (result.exitCode !== 0) {
		throw new Error(`command exited with code ${result.exitCode}`);
	}
}
