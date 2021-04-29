#!/usr/bin/env bash
COMMAND="${1:-/bin/bash}"
IMAGE="${2-jmimick/mdb-cfn-dev}"
NAME="${3-mdb-cfn-dev}"
docker run -it --rm \
   -v $HOME/.aws:/root/.aws \
   -v $HOME/.config:/root/.config \
   -v get-started-aws:/cache \
   -v /tmp:/tmp \
   -v "$(pwd)":/workspace \
   -v /var/run/docker.sock:/var/run/docker.sock \
   -e MCLI_PUBLIC_API_KEY \
   -e MCLI_PRIVATE_API_KEY \
   -e MCLI_ORG_ID \
   -e ATLAS_PUBLIC_KEY \
   -e ATLAS_PRIVATE_KEY \
   -e ATLAS_ORG_ID \
   -e AWS_DEFAULT_REGION \
   -e AWS_ACCESS_KEY_ID \
   -e AWS_SECRET_ACCESS_KEY \
   -e AWS_SESSION_TOKEN \
   -w /workspace \
   --name "${NAME}" \
   "${IMAGE}" -- "${COMMAND}"
