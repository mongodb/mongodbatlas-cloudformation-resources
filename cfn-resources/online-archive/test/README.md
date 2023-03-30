# MongoDB::Atlas::OnlineArchive

## Impact 
The following components use this resource and are potentially impacted by any changes. They should also be validated to ensure the changes do not cause a regression.
 - Online archive L1 CDK constructor


## Prerequisites 
### Resources needed to run the manual QA
All resources are created as part of `cfn-testing-helper.sh`:

- Atlas Project
- Cluster with sample data
- Database name and collection name (using sample_airbnb)

## Manual QA
Please follow the steps in [TESTING.md](../../../TESTING.md.md).


### Success criteria when testing the resource
1. Online archive should be created in the specified test project:

  ![image](https://user-images.githubusercontent.com/122359335/227655088-8c1d44d3-da02-4413-af2a-5d814ab113a8.png)

2. Ensure general [CFN resource success criteria](../../../TESTING.md#success-criteria-when-testing-the-resource) for this resource is met.


## Important Links
- [API Documentation](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/#tag/Online-Archive)
- [Resource Usage Documentation](https://www.mongodb.com/docs/atlas/online-archive/connect-to-online-archive/)
