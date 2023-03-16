## Description

Please include a summary of the fix/feature/change, including any relevant motivation and context.

Link to any related issue(s):

## Type of change:

- [ ] Bug fix (non-breaking change which fixes an issue)
- [ ] New feature (non-breaking change which adds functionality)
- [ ] Breaking change (fix or feature that would cause existing functionality to not work as
  expected)
- [ ] This change requires a documentation update

## Manual QA performed:

- [ ] cfn invoke for each of CRUDL/cfn test
- [ ] Updated resource in  [example](https://github.com/mongodb/mongodbatlas-cloudformation-resources/tree/master/examples)
- [ ] Published to AWS private registry
- [ ] Created and updated resource in AWS
- [ ] Deleted stack to ensure resources are deleted
- [ ] Created multiple resources in same stack
- [ ] Validated in Atlas UI
- [ ] Included screenshots

## Required Checklist:

- [ ] I have signed the [MongoDB CLA](https://www.mongodb.com/legal/contributor-agreement)
- [ ] I have added tests that prove my fix is effective or that my feature works
- [ ] I have added any necessary documentation (if appropriate)
- [ ] I have run `make fmt` and formatted my code
- [ ] For CFN Resources: I have released by changes in the private registry and proved by change
  works in Atlas

## Further comments

