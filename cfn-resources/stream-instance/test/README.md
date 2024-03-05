# Stream Instance

## Prerequisites 
### Resources needed to run the manual QA
- Atlas organization
- Atlas project


All resources are created as part of `cfn-testing-helper.sh`

## Manual QA
Please, follows the steps in [TESTING.md](../../../TESTING.md).


### Success criteria when testing the resource
- The Stream instance should be visible in the stream processing page



## Important Links
- [API Documentation](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/#tag/Projects)
- [Resource Usage Documentation](https://www.mongodb.com/docs/atlas/tutorial/manage-projects/)

## Contract Testing


### Build Handler
```bash
make build
```
### Run the handler in a docker container
```bash
# Required the docker daemon running
sam local start-lambda --skip-pull-image
```

### Run contract tests
```bash
cfn test --function-name TestEntrypoint --verbose
```