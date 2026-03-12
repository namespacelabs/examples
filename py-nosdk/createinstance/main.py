# This example demonstrates how to create a Namespace instance with an nginx
# container using only httpx and the Connect protocol (JSON).
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

    from datetime import datetime, timedelta, timezone

    deadline = (datetime.now(timezone.utc) + timedelta(hours=1)).isoformat()

    with httpx.Client(timeout=300) as client:
        # Create an instance with an nginx container.
        create_resp = connect_call(
            client,
            COMPUTE_API,
            SERVICE,
            "CreateInstance",
            token,
            {
                "shape": {"virtualCpu": 2, "memoryMegabytes": 4096, "machineArch": "amd64"},
                "documentedPurpose": "createinstance example",
                "deadline": deadline,
                "containers": [
                    {
                        "name": "nginx",
                        "imageRef": "nginx",
                        "args": [],
                        "exportPorts": [
                            {"name": "nginx", "containerPort": 80, "proto": "TCP"},
                        ],
                    },
                ],
            },
        )

        instance_id = create_resp["metadata"]["instanceId"]
        print(f"[namespace] Instance: {create_resp['instanceUrl']}", file=sys.stderr)
        print(json.dumps(create_resp, indent=2), file=sys.stderr)

        # Wait until the instance is ready.
        wait_resp = connect_call(
            client,
            COMPUTE_API,
            SERVICE,
            "WaitInstanceSync",
            token,
            {"instanceId": instance_id},
        )

        print(json.dumps(wait_resp, indent=2), file=sys.stderr)


if __name__ == "__main__":
    main()
