#!/bin/bash
URL="https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip"
if [ $1 == "linux/arm64" ]; then
    URL="https://awscli.amazonaws.com/awscli-exe-linux-aarch64.zip"
fi
curl ${URL} -o "awscliv2.zip" && unzip awscliv2.zip && ./aws/install