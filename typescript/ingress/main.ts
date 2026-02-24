import { loadDefaults } from "@namespacelabs/sdk/auth";
import { createComputeClient } from "@namespacelabs/sdk/api/compute";
import { Timestamp } from "@bufbuild/protobuf";
import { ContainerRequest_Network } from "@namespacelabs/sdk/proto/namespace/cloud/compute/v1beta/compute_pb";

const deadlineMinutes = parseInt(process.argv[2] || "5", 10);

void main();

async function main() {
	const tokenSource = await loadDefaults();

	const api = createComputeClient({ tokenSource });

	// Create an instance with an nginx container on the host network, listening on port 8080.
	const resp = await api.compute.createInstance({
		shape: { virtualCpu: 2, memoryMegabytes: 4096 },
		documentedPurpose: "ingress example",
		deadline: Timestamp.fromDate(
			new Date(Date.now() + deadlineMinutes * 60 * 1000)
		),
		containers: [
			{
				name: "nginx",
				imageRef: "nginx",
				args: [
					"sh",
					"-c",
					`echo 'server { listen 8080; location / { default_type text/plain; return 200 "hello from nginx\\n"; } }' > /etc/nginx/conf.d/default.conf && nginx -g 'daemon off;'`,
				],
				network: ContainerRequest_Network.HOST,
			},
		],
	});

	const instanceId = resp.metadata!.instanceId;

	console.error(`[namespace] Instance: ${resp.instanceUrl}`);
	console.error(`[Waiting until instance becomes ready]`);

	// Wait until the instance is ready.
	const waitStream = api.compute.waitInstance({ instanceId });
	for await (const _ of waitStream) {
	}

	// Expose port 8080 via public ingress.
	const ingressResp = await api.compute.createIngress({
		instanceId,
		ingresses: [
			{
				name: "nginx",
				exportedPortBackend: { port: 8080 },
			},
		],
	});

	for (const ingress of ingressResp.allocatedIngresses) {
		console.log(`https://${ingress.fqdn}`);
	}
}
