# atlas-cfn-deploy

Credit to: https://github.com/marketplace/actions/cfn-deploy-action 

## Usage

An example workflow for deploying a all the MongoDB Atlas AWS CloudFormation Custom Resources

```
    - name: Run the atlas-cfn-deploy custom action.
      uses: ./.github/actions/atlas-cfn-deploy
      with:
        AWS_ACCESS_KEY_ID: ${{ secrets.AWS_ACCESS_KEY_ID }}
        AWS_SECRET_ACCESS_KEY: ${{ secrets.AWS_SECRET_ACCESS_KEY }}
        AWS_REGION_INPUT: ${{ github.event.inputs.region }}
        AWS_REGION_SECRET: ${{ secrets.AWS_REGION }}

```

## Secrets

 - `AWS_ACCESS_KEY_ID` – (Required) The AWS access key part of your credentials [more info](https://help.github.com/en/actions/automating-your-workflow-with-github-actions/creating-and-using-encrypted-secrets)
 
 - `AWS_SECRET_ACCESS_KEY` – (Required) The AWS secret access key part of your credentials [more info](https://help.github.com/en/actions/automating-your-workflow-with-github-actions/creating-and-using-encrypted-secrets)

## Environment variables

All environment variables listed in the official documentation are supported.

The custom env variables to be added are:

`AWS_REGION` - Region to which you need to deploy your app<br>
`CAPABLITIES` - IAM capablities for the cloudformation stack<br>

## atlas-cfn-deploy.py

Helper script to publish cfn resources.

```
usage: atlas-cfn-deploy.py [-h] [-p PARAMETERS] [-r REGION] noun

    atlas-cfn-cli.py
    MongoDB Atlas AWS CloudFormation Custom Resource CLI
    This tool helps you deploy Atlas resources for your CF templates
    Set LOGURU_LEVEL to higher for more logs
     

positional arguments:
  noun

optional arguments:
  -h, --help            show this help message and exit
  -p PARAMETERS, --parameters PARAMETERS
                        resource parameters
  -r REGION, --region REGION
                        region

```

Examples:

```bash
➜  atlas-cfn-deploy git:(master) ✗ ./atlas-cfn-deploy.py --region=us-east-2 all+ 
```
