# Templates used for generating the required parameters for SSM Automation

## locations.json 
    To specify the regions and AWS accounts

## params.json 
    To specify the parameters requried for running the SSM Automation Document
## <resource_name>.json 
    To specify the parameters requried for running the SSM Automation Document for a specific resource

# for local dev testing
run the below command to run the workflow locally

For dev testing using act tool. For more info refer https://github.com/nektos/act

```
cd .githib/workflows
act -W ./cfn-publish-resource.yaml -P ubuntu-latest=ubuntu:latest --secret-file templates/params.secrets -e templates/inputs.json --pull=false
```

## inputs.json 
    inputs.json file is used to specify the input parameters for the workflow. i.e. github.event.inputs.<varName>
    eg:
```
        {
            "inputs": {
            "regions": "me-south-1,ap-south-2",
            "resourceNames": "ldap-configuration,federated-settings-org-role-mapping",
            "otherParams": ""
            }
        }
```

## add .secret file in the templates directory and initialize all the required secrets for the workflow
```
        ATLAS_PUBLIC_KEY="xxxxxxxx"
        ATLAS_PRIVATE_KEY="xxxxx-xxxxx-xxxxx-xxxxx-xxxxx"
        ATLAS_ORG_ID="xxxxxxxxxxxxxxxx"
        ATLAS_FEDERATED_SETTINGS_ID=""
        
        AWS_ACCESS_KEY_ID=""
        AWS_SECRET_ACCESS_KEY=""
        AWS_ACCOUNT_ID=""
        
        LDAP_BIND_PASSWORD=""
        LDAP_BIND_USER_NAME="CN=XXXXXXXX,CN=XXXXX,DC=XXXXXXXXXXXXXX,DC=XXXXXXXXX,DC=com"
        LDAP_HOST_NAME=""
        
        WEBHOOK_CREATE_URL=
        WEBHOOK_UPDATE_URL=
        WEBHOOK_UPDATE_SECRET=
        PROMETHEUS_USER_NAME=
        PROMETHEUS_PASSWORD_NAME=
        PAGER_DUTY_CREATE_SERVICE_KEY=
        PAGER_DUTY_UPDATE_SERVICE_KEY=
        DATA_DOG_CREATE_API_KEY=
        DATA_DOG_UPDATE_API_KEY=
        OPS_GENIE_API_KEY=
        MICROSOFT_TEAMS_WEBHOOK_CREATE_URL=
        MICROSOFT_TEAMS_WEBHOOK_UPDATE_URL= 
```
