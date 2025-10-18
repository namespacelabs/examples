import * as fs from "fs/promises";

import { loadDefaults } from "@namespacelabs/sdk/auth";
import { createComputeClient } from "@namespacelabs/sdk/api/compute";
import { CreateInstanceRequest_Feature } from "@namespacelabs/sdk/proto/namespace/cloud/compute/v1beta/compute_pb";
import { Timestamp } from "@bufbuild/protobuf";

void main();

async function main() {
	const tokenSource = await loadDefaults();

	// Configure client stub.
	const computeClient = createComputeClient({ tokenSource });
	const client = computeClient.compute;

	// Create instance.
	const resp = await client.createInstance({
		shape: { virtualCpu: 2, memoryMegabytes: 4096 },
		// Run the instance for 30 mins.
		deadline: Timestamp.fromDate(new Date(Date.now() + 30 * 60 * 1000)),
		// By default the VM is created with only containerd in it and not K8s.
		features: [
			CreateInstanceRequest_Feature.KUBERNETES,
			CreateInstanceRequest_Feature.KUBERNETES_INGRESS_MANAGER,
		],
	});
	const instanceId = resp.metadata.instanceId;

	console.log("Kubernetes Cluster created.");
	console.log("   - ID:  ", instanceId);
	console.log("   - URL: ", resp.instanceUrl);
	console.log("   - Deadline: ", resp.metadata.deadline.toDate());
	console.log();

	// Wait for the instance to boot up and K8s to initialize.
	console.log("Waiting for the cluster to initialize...");
	const waitStream = client.waitInstance({ instanceId });
	for await (const _ of waitStream);
	console.log("   - cluster initialized.");

	// Get kubeconfig.
	console.log();
	console.log("Getting kubeconfig.yaml...");
	const kubeconfig = await client.getKubernetesConfig({ instanceId });
	const kubeconfigPath = `${instanceId}.yaml`;
	await fs.writeFile(kubeconfigPath, kubeconfig.kubeconfig);
	console.log("   - got: ", kubeconfigPath);

	console.log();
	console.log("Try:");
	console.log(`   $ kubectl --kubeconfig=${kubeconfigPath} get all`);

	if (false) {
		await client.destroyInstance({
			instanceId,
			reason: "testing",
		});
		console.log("Kubernetes Cluster destroyed");
	}
}
