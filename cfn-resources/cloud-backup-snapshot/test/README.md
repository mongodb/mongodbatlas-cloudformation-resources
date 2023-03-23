# Cloud backup snapshot

## Impact 
The following components use this resource and are potentially impacted by any changes. They should also be validated to ensure the changes do not cause a regression.
 - Cloud backup snapshot L1 CDK constructor


## Prerequisites 
### Resources needed to run the manual QA
- Atlas Project
- Cluster with backup enabled



All resources are created as part of `cfn-testing-helper.sh`

## Manual QA
Please, follows the steps in [TESTING.md](../../../TESTING.md.md).


### Success criteria when testing the resource
- Backup snapshot for the Atlas Cluster should be shown in the "Snapshots" page:
![image](https://user-images.githubusercontent.com/5663078/227233348-ea32d93a-bfc6-468a-b111-fb12bc0a50ec.png)
## Important Links
- [API Documentation](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/#tag/Alert-Configurations/operation/listAlertConfigurations)
- [Resource Usage Documentation](https://www.mongodb.com/docs/atlas/configure-alerts/#configure-an-alert)