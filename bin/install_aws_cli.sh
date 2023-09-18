#!/usr/bin/env bash

cd /workspace

FILENAME="awscliv2.zip"

rm -f "/workspace/$FILENAME"
rm -rf "/workspace/aws"


curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "$FILENAME"
unzip $FILENAME
sudo ./aws/install

aws sts get-caller-identity

cd $PROJECT_ROOT