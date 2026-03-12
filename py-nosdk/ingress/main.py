# This example demonstrates how to use the Namespace Compute API without an SDK,
# using only httpx and the Connect protocol (JSON).
#
# It creates an nginx container on the host network listening on port 8080,
# then exposes it via CreateIngress and outputs the public URL.
#
# Authentication: set one of the following environment variables:
#   NSC_TOKEN      - a bearer token, e.g. from `nsc auth generate-dev-token`
#   NSC_TOKEN_FILE - path to a token.json file; set automatically inside a
#                    Namespace instance
#
# WARNING: Bearer tokens expire and need to be refreshed. This example does not
# handle token lifecycle.

from __future__ import annotations

import json
import os
import sys

import httpx

REGION = "us"
COMPUTE_API = f"https://{REGION}.compute.namespaceapis.com"
SERVICE = "namespace.cloud.compute.v1beta.ComputeService"


def load_bearer_token() -> str:
    # Prefer NSC_TOKEN if set (e.g. from `nsc auth generate-dev-token`).
    token = os.environ.get("NSC_TOKEN")
    if token:
        return token

    # Fall back to NSC_TOKEN_FILE (set automatically inside a Namespace instance).
    token_file = os.environ.get("NSC_TOKEN_FILE")
    if not token_file:
        raise RuntimeError(
            "Set NSC_TOKEN (e.g. from `nsc auth generate-dev-token`) "
            "or NSC_TOKEN_FILE (path to a token.json)."
        )

    with open(token_file) as f:
        token_json = json.load(f)

    bearer = token_json.get("bearer_token")
    if not bearer:
        raise RuntimeError("Token file does not contain a bearer_token")
    return bearer


def connect_call(
    client: httpx.Client,
    base_url: str,
    service: str,
    method: str,
    token: str,
    body: dict,
) -> dict:
    """Make a Connect (JSON) unary RPC call."""
    url = f"{base_url}/{service}/{method}"
    resp = client.post(
        url,
        json=body,
        headers={
            "Content-Type": "application/json",
            "Authorization": f"Bearer {token}",
        },
    )
    if resp.status_code != 200:
        raise RuntimeError(f"{method} failed ({resp.status_code}): {resp.text}")
    return resp.json()


def main() -> None:
    token = load_bearer_token()

    deadline_minutes = int(sys.argv[1]) if len(sys.argv) > 1 else 5

    from datetime import datetime, timedelta, timezone

    deadline = (datetime.now(timezone.utc) + timedelta(minutes=deadline_minutes)).isoformat()

    with httpx.Client(timeout=300) as client:
        # Create an instance with an nginx container on the host network, listening on port 8080.
        create_resp = connect_call(
            client,
            COMPUTE_API,
            SERVICE,
            "CreateInstance",
            token,
            {
                "shape": {"virtualCpu": 2, "memoryMegabytes": 4096, "machineArch": "amd64"},
                "documentedPurpose": "ingress example (no-sdk)",
                "deadline": deadline,
                "containers": [
                    {
                        "name": "nginx",
                        "imageRef": "nginx",
                        "args": [
                            "sh",
                            "-c",
                            "echo 'server { listen 8080; location / { default_type text/plain; return 200 \"hello from nginx\\n\"; } }' > /etc/nginx/conf.d/default.conf && nginx -g 'daemon off;'",
                        ],
                        "network": "HOST",
                    },
                ],
            },
        )

        instance_id = create_resp["metadata"]["instanceId"]
        print(f"[namespace] Instance: {create_resp['instanceUrl']}", file=sys.stderr)
        print("[Waiting until instance becomes ready]", file=sys.stderr)

        # WaitInstanceSync blocks until the instance is ready.
        connect_call(client, COMPUTE_API, SERVICE, "WaitInstanceSync", token, {
            "instanceId": instance_id,
        })

        # Expose port 8080 via public ingress.
        ingress_resp = connect_call(
            client,
            COMPUTE_API,
            SERVICE,
            "CreateIngress",
            token,
            {
                "instanceId": instance_id,
                "ingresses": [
                    {
                        "name": "nginx",
                        "exportedPortBackend": {"port": 8080},
                    },
                ],
            },
        )

        for ingress in ingress_resp["allocatedIngresses"]:
            print(f"https://{ingress['fqdn']}")


if __name__ == "__main__":
    main()
