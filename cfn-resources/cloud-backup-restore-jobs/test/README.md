# Cloud backup restore jobs 

## Impact 
The following components use this resource and are potentially impacted by any changes. They should also be validated to ensure the changes do not cause a regression.
 - Cloud backup restore jobs  L1 CDK constructor


## Prerequisites 
### Resources needed to run the manual QA
- Atlas Project
- Cluster with backup enabled
- SnapshotId


All resources are created as part of `cfn-testing-helper.sh`

## Manual QA
Please, follows the steps in [TESTING.md](../../../TESTING.md.md).


### Success criteria when testing the resource
- Backup restore job for the Atlas Cluster should be shown in "Restores & Downloads" page:
![image](https://user-images.githubusercontent.com/5663078/227225795-0f1b6650-95fe-40ca-942d-99902b747aa2.png)
## Important Links
- [API Documentation](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/#tag/Alert-Configurations/operation/listAlertConfigurations)
- [Resource Usage Documentation](https://www.mongodb.com/docs/atlas/configure-alerts/#configure-an-alert)