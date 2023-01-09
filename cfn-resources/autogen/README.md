# Auto Generate CloudFormation template and generic GO Code

## How to generate Resource setup and Respective Go template for a new Resource

### Step 1:

Update mapping.json in schema folder with new resource details

For Example:
``` json
{
    "resources" : [
        {
            "typeName" :"cloudProviderAccess",
            "openApiPath" : [
            "/api/atlas/v1.0/groups/{groupId}/clusters/{clusterName}/fts/indexes",
            "/api/atlas/v1.0/groups/{groupId}/clusters/{clusterName}/fts/indexes/{databaseName}/{collectionName}",
            "/api/atlas/v1.0/groups/{groupId}/clusters/{clusterName}/fts/indexes/{indexId}"
            ]
        }
    ]
}
```


### Step 2:
Run below command to generate schema file that is required for ``cfn generagte`` in the next step.
```bash
    make schema
``` 
### Step 3:
Run below command to generate Resource template which reads
/configs/<resource>-schema.json and /configs/<resource>-req.json

For example: to create search-index resource run below command.
search-indexes is the directory name for resource to create
indexes is the ResourceName that can be found in the end of the OpenAPI URL.
For Example: /api/atlas/v1.0/groups/{groupId}/clusters/{clusterName}/fts/indexes
```bash
    make create dirName=search-indexes resourceName=indexes
``` 

