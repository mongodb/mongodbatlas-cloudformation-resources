# How to create a MongoDB::Atlas::StreamConnection

## Step 1: Activate the stream connection resource in cloudformation

Step a: Create Role using [execution-role.yaml](https://github.com/mongodb/mongodbatlas-cloudformation-resources/blob/master/examples/execution-role.yaml) in CFN resources folder.

Step b: Search for Mongodb::Atlas::StreamConnection resource.

         (CloudFormation > Public extensions > choose 'Third party' > Search with " Execution name prefix = MongoDB " )

Step c: Select and activate
Enter the RoleArn that is created in step 1.

Your StreamConnection Resource is ready to use.

## Step 2: Create template using example JSON files

    Examples for each connection type:

    **Cluster type** - Connect to an Atlas cluster:
    - [cluster-stream-connection.json](cluster-stream-connection.json)

    **Kafka type** - Connect to a Kafka cluster:
    - [kafka-stream-connection.json](kafka-stream-connection.json) - Basic authentication with username/password
    - [kafka-oauth-stream-connection.json](kafka-oauth-stream-connection.json) - OAuth 2.0 authentication (e.g., Confluent Cloud)
    - [kafka-ssl-cert-stream-connection.json](kafka-ssl-cert-stream-connection.json) - SSL client certificate authentication (mutual TLS)
    - [kafka-schema-registry-stream-connection.json](kafka-schema-registry-stream-connection.json) - Kafka with Schema Registry integration

    **Sample type** - Use sample datasets:
    - [sample-stream-connection.json](sample-stream-connection.json)

    **AWSLambda type** - Connect to AWS Lambda:
    - [aws-lambda-stream-connection.json](aws-lambda-stream-connection.json)

    **Https type** - Connect via HTTPS:
    - [https-stream-connection.json](https-stream-connection.json)

    Note: Make sure you are providing appropriate values for:
    1. ProjectId
    2. WorkspaceName (or InstanceName - deprecated)
    3. ConnectionName
    4. Type: Cluster, Kafka, Sample, AWSLambda, or Https
    5. Profile (optional)
    6. Type-specific fields (ClusterName for Cluster type, BootstrapServers for Kafka type, etc.)

## AWS Lambda with Setup and Authorization

The [AWS Lambda template](aws-lambda-stream-connection.json) accepts an existing authorized IAM role through
`LambdaRoleArn`. This is the role Atlas assumes to invoke functions, separate from the CloudFormation execution
role used when activating the resource.

To provision access in the same stack, use the trust handshake from the
[Cloud Provider Access example](/examples/cloud-provider-access-authorization/README.md):

1. Create `CloudProviderAccessSetup` to obtain the Atlas principal and external ID.
2. Create an `AWS::IAM::Role` with logical ID `AtlasLambdaRole`. Its trust policy must use Setup's
   `AtlasAWSAccountArn` and require `AtlasAssumedRoleExternalId`; grant `lambda:InvokeFunction` on your target
   function ARN.
3. Create `CloudProviderAccessAuthorization` with `RoleId: !GetAtt CloudProviderAccessSetup.RoleId` and
   `IamAssumedRoleArn: !GetAtt AtlasLambdaRole.Arn`.
4. Add the following connection under `Resources`. The snippet assumes the `ProjectId`, `Profile`, `WorkspaceName`
   and `ConnectionName` parameters, plus the IAM role and Authorization resources above.

```yaml
LambdaConnection:
  Type: MongoDB::Atlas::StreamConnection
  DependsOn: CloudProviderAccessAuthorization
  Properties:
    ProjectId: !Ref ProjectId
    Profile: !Ref Profile
    WorkspaceName: !Ref WorkspaceName
    ConnectionName: !Ref ConnectionName
    Type: AWSLambda
    Aws:
      RoleArn: !GetAtt AtlasLambdaRole.Arn
```

`Aws.RoleArn` takes the AWS IAM role ARN, **not** the Atlas registration `RoleId`. The explicit dependency is
needed because the role ARN alone does not depend on successful Atlas authorization. Atlas documents this
[Unified AWS Access integration for Lambda](https://www.mongodb.com/docs/atlas/atlas-stream-processing/ops-guide/#passwordless-aws-authentication).

Make each stream processor that uses this connection depend on `LambdaConnection`. Teardown must remove those
processors and the connection before Authorization removes the shared Atlas registration. Retaining a processor
or connection while deleting Authorization can leave an active consumer without access or cause a role-in-use
failure; review their lifecycle together.

The current CloudFormation schema supports `Cluster`, `Kafka`, `Sample`, `AWSLambda` and `Https`. Do not use
Terraform or Atlas examples for S3, Kinesis, GCP Pub/Sub or Azure Blob Storage with this resource: those connection
types are not currently included in its `Type` enum.
