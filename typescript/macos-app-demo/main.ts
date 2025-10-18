import { loadDefaults } from "@namespacelabs/sdk/auth";
import { createComputeClient } from "@namespacelabs/sdk/api/compute";
import { WaitInstanceResponse } from "@namespacelabs/sdk/proto/namespace/cloud/compute/v1beta/compute_pb";
import { Timestamp } from "@bufbuild/protobuf";

void main();

async function main() {
    const tokenSource = await loadDefaults();
    const api = createComputeClient({ tokenSource });

    const created = await api.compute.createInstance({
        shape: {
            os: "macos",
            machineArch: "arm64",
            virtualCpu: 6,
            memoryMegabytes: 14336,
            selectors: [
                { name: "macos.version", value: "26.x" },
            ],
        },
        deadline: Timestamp.fromDate(new Date(Date.now() + 30 * 60 * 1000)),
        applications: [
            {
                name: "demo",
                imageRef: "nscr.io/01gr1g2rpb7ahzddy3f227exq9/demo@sha256:fa0e6861e56ac5520e55df1b5bf29cf8a070c8dd64c3121d02f5c039db968da0",
                command: "./imagetool",
                args: ["--help"],
                environment: { "SHELL": "/usr/bin/bash" },
                experimental: { includeLogs: ["/var/log/wifi.log"] },
            },
        ],
        experimental: {
            preStartHook: [{
                command: {
                    command: "/bin/bash",
                    args: ["-l", "-c", "ls -la $HOME"],
                },
            }],
        },
    });
    console.log("created", created.instanceUrl);

    const instanceId = created.metadata.instanceId;

    const waited = await api.compute.waitInstanceSync({ instanceId: instanceId });
    console.log("started in", waited.metadata.hwDeployment?.geoContinent);

    for await (const x of api.observability.streamInstanceLogs({ instanceId: instanceId, follow: true })) {
        for (const l of x.lines) {
            console.log(l.timestamp.toDate(), x.labels, l.content, l.stream);
        }
    }


    let finished: WaitInstanceResponse;
    for await (const x of api.compute.waitInstance({ instanceId: instanceId, destroyedOk: true })) {
        finished = x;
    }
    console.log("finished");
}
