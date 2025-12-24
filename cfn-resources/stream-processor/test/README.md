# MongoDB::Atlas::StreamProcessor

## Impact
The following components use this resource and are potentially impacted by any changes. They should also be validated to ensure the changes do not cause a regression.
 - Stream Processor L1 CDK constructor


## Prerequisites
### Resources needed to run the manual QA
All resources are created as part of `cfn-testing-helper.sh`:

- Atlas Project
- Atlas Stream Instance/Workspace (LONG-RUNNING operation, can take 10-30+ minutes)
- Cluster (for DLQ connection testing - inputs_3)
- Stream Connection (for DLQ connection testing - inputs_3)

**IMPORTANT**: Stream Instance/Workspace creation is a LONG-RUNNING operation that can take 10-30+ minutes. The `cfn-test-create-inputs.sh` script will create the workspace and wait for it to be ready before proceeding.

## Manual QA
Please follow the steps in [TESTING.md](../../../TESTING.md).


### Success criteria when testing the resource

#### 1. Resource Creation Verification

A Stream Processor should be created in the specified test project for the specified Atlas Stream workspace/instance:

**Atlas UI Verification:**
- Navigate to Atlas UI → Your Project → Stream Processing
- Select the stream workspace/instance used in the test
- Go to the **Processors** tab
- Verify the processor appears with:
  - **Name**: Matches the `ProcessorName` from the test input
  - **State**: Matches the `State` in the template (CREATED, STARTED, or STOPPED)
  - **Pipeline**: Click on the processor to view details and verify:
    - Pipeline stages match the `Pipeline` configuration in the template
    - Source connection name is correct
    - Merge target connection, database, and collection are correct

**Atlas CLI Verification:**
```bash
atlas streams processors describe <PROCESSOR_NAME> \
  --instance <WORKSPACE_NAME> \
  --projectId <PROJECT_ID>
```
- Verify `id` field is present (matches CloudFormation `Id` attribute)
- Verify `name` matches `ProcessorName`
- Verify `state` matches `State` parameter
- Verify `pipeline` array matches the `Pipeline` JSON string

#### 2. DLQ Configuration Verification (inputs_3)

For processors with DLQ configuration:
- In Atlas UI: Verify DLQ settings are displayed in processor details
- Via Atlas CLI: Verify `options.dlq` object contains:
  - `connectionName`: Matches `Options.Dlq.ConnectionName`
  - `db`: Matches `Options.Dlq.Db`
  - `coll`: Matches `Options.Dlq.Coll`

#### 3. Backward Compatibility Testing

Test both field names work correctly:
- **Test with `WorkspaceName`** (preferred field):
  - Create processor using `WorkspaceName` parameter
  - Verify processor is created successfully
  - Verify both `WorkspaceName` and `InstanceName` are set in returned model (for primary identifier)
- **Test with `InstanceName`** (deprecated field):
  - Create processor using `InstanceName` parameter
  - Verify processor is created successfully
  - Verify both `WorkspaceName` and `InstanceName` are set in returned model
  - Verify `WorkspaceName` is automatically set from `InstanceName` for forward compatibility

#### 4. State Transition Testing

Test all valid state transitions:
- **Create with `State: CREATED`**:
  - Verify processor is created in CREATED state
  - Verify processor does not start processing automatically
- **Create with `State: STARTED`**:
  - Verify processor is created and transitions to STARTED state
  - Verify this is a long-running operation (may take several minutes)
  - Verify callback-based state management handles the transition
- **Update state from CREATED to STARTED**:
  - Verify processor stops (if needed) before update
  - Verify processor starts after update completes
  - Verify state transition is successful
- **Update state from STARTED to STOPPED**:
  - Verify processor stops before update
  - Verify processor remains stopped after update
  - Verify state transition is successful

#### 5. Timeout and Cleanup Behavior

- **Verify `Timeouts.Create` is respected**:
  - Set a short timeout (e.g., 1 minute) for a processor that takes longer to start
  - Verify timeout is triggered after the specified duration
- **Verify `DeleteOnCreateTimeout` behavior**:
  - When `DeleteOnCreateTimeout: true` and timeout occurs:
    - Verify processor deletion is triggered
    - Verify resource is cleaned up from Atlas
  - When `DeleteOnCreateTimeout: false` and timeout occurs:
    - Verify processor is not deleted
    - Verify resource remains in Atlas (may be in partial state)

#### 6. Primary Identifier Verification

Verify all primary identifier fields are present in returned models:
- `ProjectId`: Always present
- `WorkspaceName`: Always present (set from `InstanceName` if needed)
- `InstanceName`: Always present (set from `WorkspaceName` if needed)
- `ProcessorName`: Always present
- `Profile`: Always present

This is critical for CloudFormation to properly track the resource.

#### 7. General CFN Resource Success Criteria

Ensure general [CFN resource success criteria](../../../TESTING.md#success-criteria-when-testing-the-resource) for this resource is met:
- All CRUD operations work correctly
- Read-after-Create returns correct values
- Update operations preserve primary identifier
- Delete operations clean up resources
- Error handling is appropriate


## Important Links
- [API Documentation](https://www.mongodb.com/docs/api/doc/atlas-admin-api-v2/group/endpoint-streams)
- [Resource Usage Documentation](https://www.mongodb.com/docs/atlas/atlas-sp/overview/)

## Unit Testing Locally

The local tests are integrated with the AWS `sam local` and `cfn invoke` tooling features:

```
sam local start-lambda --skip-pull-image
```
then in another shell:
```bash
repo_root=$(git rev-parse --show-toplevel)
cd ${repo_root}/cfn-resources/stream-processor
cfn invoke resource CREATE stream-processor-sample-cfn-request.json
cfn invoke resource DELETE stream-processor-sample-cfn-request.json
cd -
```

Both CREATE & DELETE tests must pass.

## Test Input Files

The test directory contains the following input files:

- `inputs_1_create.json` / `inputs_1_update.json`: Basic stream processor with WorkspaceName, CREATED state
- `inputs_2_create.json` / `inputs_2_update.json`: Stream processor with STARTED state, timeout configuration, and DeleteOnCreateTimeout
- `inputs_3_create.json` / `inputs_3_update.json`: Stream processor with InstanceName (backward compatibility) and DLQ options

All input files respect:
- AWS-only behavior (no Azure/GCP-only parameters)
- Required fields: ProjectId, ProcessorName, Pipeline
- Backward compatibility: Supports both WorkspaceName and InstanceName
- Schema validation: All fields match the final CFN schema
