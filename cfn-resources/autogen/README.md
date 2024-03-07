# Using autogen 

## Autogenerating CFN resource schema and generic Go code for a new resource:

### Step 1:

Update `schema-gen/mapping.json` with the new resource details (for this example we will use Stream Connection). The typeName must be in CamelCase, as is a CloudFormation standard.

For example:
``` json
{
  "resources" : [
     {
      "typeName": "StreamConnection",
      "openApiPath": [
        "/api/atlas/v2/groups/{groupId}/streams/{tenantName}/connections",
        "/api/atlas/v2/groups/{groupId}/streams/{tenantName}/connections/{connectionName}"
      ],
      "contentType": "application/vnd.atlas.2023-02-01+json"
    }
  ]
}
```

### Step 2:
Run below command to generate schema file that is required for ``cfn generate`` in the next step.
```bash
    make schema
``` 

this procedure should generate 2 files in the schema folder (custom db role for this example)

![img.png](https://github.com/mongodb/terraform-provider-mongodbatlas/assets/122359335/c5d8f2b8-6e7c-4a28-b205-059e69327051)

- mongodb-atlas-stream-connection.json : is the actual schema used to generate the resource
- mongodb-atlas-stream-connection-req.json : contains metadata related to the code generation, like potentially required fields, and fields that might be needed by the Atlas go Client

### Step 3:
Run make create command to generate Resource template which reads the files generated in Step #2
`/schemas/<resource>-schema.json` and `/schemas/<resource>-req.json`

#### the make create requires 3 inputs:

- schemaFileName = is the name of the resource in the file schema, check the file created in the /schemas folder, for this example the schemaFileName is "customdbrole" (mongodb-atlas-**streamconnection**.json)
- dirName = is the name of the directory that will be used to store the resource, (dash case is the convention for the folder name) for this example we are using "custom-db-role"
- typeName = is the partial resource type name, used by cloud formation, (MongoDB::Atlas::{TypeName}). This  MUST be the same as in the schema generated (mongodb-atlas-streamconnection.json) in step 2.

```bash
        make create dirName=stream-connection resourceName=streamconnection typeName=StreamConnection
```


After executing the command you need to provide the Go import path, which will be prompted on the terminal, for example, `github.com/mongodb/mongodbatlas-cloudformation-resources/stream-connection`


```
> make create dirName=stream-connection resourceName=streamconnection typeName=StreamConnection
Enter the go import path as: github.com/mongodb/mongodbatlas-cloudformation-resources/stream-connection 
cd .. && \
        mkdir stream-connection && cp autogen/schemas/mongodb-atlas-streamconnection.json stream-connection/mongodb-atlas-streamconnection.json && \
        cd stream-connection && \
        cfn init -t MongoDB::Atlas::StreamConnection -a r && cfn generate
Initializing new project
One language plugin found, defaulting to go

> Enter the GO Import path
>> github.com/mongodb/mongodbatlas-cloudformation-resources/stream-connection
```
