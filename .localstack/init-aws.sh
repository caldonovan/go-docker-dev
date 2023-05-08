#!/usr/bin/env bash

set -euo pipefail

export LOCALSTACK_HOST=localhost
export AWS_DEFAULT_REGION=eu-central-1
export AWS_ACCESS_KEY_ID=12345678
export AWS_SECRET_ACCESS_KEY=12345678

aws --endpoint-url=http://${LOCALSTACK_HOST}:4566 sqs create-queue --queue-name queue1 --attributes VisibilityTimeout=30
aws --endpoint-url=http://${LOCALSTACK_HOST}:4566 s3 mb s3://bucket1