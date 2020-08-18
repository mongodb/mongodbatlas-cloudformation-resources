# cfn-deploy
<img src="https://github.com/intuit/cfn-deploy/blob/master/.github/cfn-deploy-logo.png" align=right alt="cfn-deploy" width="250"> 
A simple github action to deploy cloudformation yaml files to AWS

## Usage

An example workflow for deploying a cloudformation template follows.

```
 - uses: intuit/cfn-deploy@master
      env:
        AWS_REGION: us-east-2
        STACK_NAME: cfn-deploy
        TEMPLATE_FILE: ec2.yml
        PARAMETERS_FILE: parameter.json
        CAPABLITIES: CAPABILITY_IAM
        AWS_ACCESS_KEY_ID: ${{secrets.AWS_ACCESS_KEY_ID}}
        AWS_SECRET_ACCESS_KEY: ${{secrets.AWS_SECRET_ACCESS_KEY}}

```

## Secrets

 - `AWS_ACCESS_KEY_ID` – (Required) The AWS access key part of your credentials [more info](https://help.github.com/en/actions/automating-your-workflow-with-github-actions/creating-and-using-encrypted-secrets)
 
 - `AWS_SECRET_ACCESS_KEY` – (Required) The AWS secret access key part of your credentials [more info](https://help.github.com/en/actions/automating-your-workflow-with-github-actions/creating-and-using-encrypted-secrets)

## Environment variables

All environment variables listed in the official documentation are supported.

The custom env variables to be added are:

`AWS_REGION` - Region to which you need to deploy your app<br>
`STACK_NAME` - Cloudformation Stack Name <br>
`TEMPLATE_FILE` - Cloudformation template yaml file<br>
`PARAMETERS_FILE` - Input parameters to the cloudformation stack as json file<br>
`CAPABLITIES` - IAM capablities for the cloudformation stack<br>

## Contributing

See [Contributing](https://github.com/intuit/cfn-deploy/blob/master/.github/CONTRIBUTING.md)
