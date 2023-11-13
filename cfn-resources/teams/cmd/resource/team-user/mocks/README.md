## Introduction

These mocks are generated through the `mockgen` command.

## HOW-TO: Update mocks

1. In order to update these mocks, you need to have `mockgen` **installed** on your machine: `go install github.com/golang/mock/mockgen@latest`.
1. `cd` in the root of the `mongodbatlas-clouformation-resources` repo
1. After that, you can run the `mockgen` command as following:

```
mockgen -source cfn-resources/teams/cmd/resource/team-user/team_user.go -destination /PATH/TO/REPOSITORY/mongodbatlas-cloudformation-resources/cfn-resources/teams/cmd/resource/team-user/mocks/mocks.go -package mocks
```