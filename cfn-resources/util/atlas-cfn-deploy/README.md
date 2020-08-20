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

There are probably easier ways to do this.
