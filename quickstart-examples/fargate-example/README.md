# MEAN Stack Example - Steps to Launch the Solution

## Prerequisites

- Docker up and running
- AWS CLI with a profile configured with API keys

## Build and Push Client and Server Containers

**Note**: Make sure *Docker is up and running*. Update `profile` and `region` in `fargate-example/client/build.sh` and `fargate-example/server/build.sh` files if you are not using AWS `default` profile.

### Client Application

`build.sh` file in `fargate-example/client` directory does the following things:

- builds the container image using docker
- logs in into AWS ECR using the profile
- creates an ECR repository if not exists
- and finally pushes the container image into ECR repository.

```bash
$ cd fargate-example/client
$ ./build.sh

# Output
Login Succeeded
sha256:12f77cbd4b6e64a463d9c06b277578b6098bb6083c00e611f89e652a8337a6ea
The push refers to repository [816546967292.dkr.ecr.ca-central-1.amazonaws.com/partner-meanstack-atlas-fargate-client]
278dd0d9c609: Pushed
b6831859228a: Pushed
f1ecd284fd96: Pushed
6d626da635fc: Pushed
latest: digest: sha256:d8c3c25cb3034583332bf6b3ffd5b913f95606873bcec8672e0f52fa28b73f0c size: 1155

ClientServiceECRImageURI: 816546967292.dkr.ecr.ca-central-1.amazonaws.com/partner-meanstack-atlas-fargate-client:latest
```

Use the `ClientServiceECRImageURI` from above output as the input to `ClientServiceECRImageURI` field in CloudFormation template or `.taskcat.yml` file.

### Server Application

`build.sh` file in `fargate-example/server` directory does the following things:

- builds the container image using docker
- logs in into AWS ECR using the profile
- creates an ECR repository if not exists
- and finally pushes the container image into ECR repository.

```bash
$ cd fargate-example/server
$ ./build.sh

# Output
Login Succeeded
sha256:9fad7918f466c52c6d8d8c7d1aa383dec051c33fe98685788f28cd41fbab3dff
The push refers to repository [816546967292.dkr.ecr.ca-central-1.amazonaws.com/partner-meanstack-atlas-fargate-server]
4ffd8190a67a: Pushed
59e988eae20b: Pushed
1f22723773f1: Pushed
a51f61eea2cf: Pushed
3bcec8753323: Pushed
2d0a2eb257ad: Pushed
1508433ef884: Pushed
bc519b60e13f: Pushed
9d9e2362ae3e: Pushed
4d31756873fb: Pushed
latest: digest: sha256:82f1953535d8a454371e0091a52659dce71bf599ff78416ece2a6757900907b3 size: 2414

ServerServiceECRImageURI: 816546967292.dkr.ecr.ca-central-1.amazonaws.com/partner-meanstack-atlas-fargate-server:latest
```

Use the `ServerServiceECRImageURI` from above output as the input to `ServerServiceECRImageURI` field in CloudFormation template or `.taskcat.yml` file.

## Launch the Solution

Finally execute the CloudFormation template `templates/mongodb-atlas-mean-stack.template.yaml` with the valid parameters.
