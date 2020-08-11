#!/usr/bin/env python3
from datetime import datetime
import subprocess
import os

def main(parameters: ('resource parameters','option','p')
         , region: ('region','option','r')
         , noun):
    """
    atlas-cfn-cli.py
    MongoDB Atlas AWS CloudFormation Custom Resource CLI
    This tool helps you deploy Atlas resources for your CF templates
     """

    cli_path = os.path.dirname(os.path.realpath(__file__))
    ## HOME is parent of `util` folder, my grandma
    ATLAS_CFN_HOME = os.path.dirname(os.path.join(cli_path,"..",".."))
    print(f"ATLAS_CFN_HOME=${ATLAS_CFN_HOME}")
    res = os.chdir(f'{ATLAS_CFN_HOME}')
    print(res)
    if parameters:
        params=dict( kvp.strip().split('=') for kvp in parameters.strip().split(',') )
    else:
        params={}
    print(noun, parameters, params)
    noun=noun.lower()
    supported_resources = ["project","cluster","database-user","ip-whitelist"]
    tier2_resources = [ "cloud-provider-snapshots", "network-peering",
                        "cloud-provider-snapshot-restore-jobs",
                       "encryption-at-rest",
                       "network-container" ]

    if noun.startswith("all"):
        print("all")
        res = supported_resources
        if noun=="all+":
            for r in tier2_resources:
                res.append(r)
        if noun=="all-":
            res = tier2_resources
        for resource in res:
            command = ["python3"
                       ,os.path.join(f"{ATLAS_CFN_HOME}","util","atlas-cfn-cli.py")
                       ,f"{resource}"
                       ,"--region"
                       ,f"{region}"]
            if parameters:
                command.append(["--parameters",f"{parameters}"])
            res = subprocess.run(command)

            print("---------------------------------------------------------")
            print(f"Resource: {resource}\nCommand: {command}\bResponse:{res}")
            print("---------------------------------------------------------")

    else:
        submit_cmd=["cfn","submit", "-v", "--region", f"{region}","--set-default"]
        cwd = os.getcwd()
        print(f"cwd===> {cwd}")
        try:
            res = os.chdir(f'{noun}')
            print(res)
            make_response = subprocess.run(["make"])
            print(make_response)
            submit_response = subprocess.run(submit_cmd)
            print(submit_response)
            res = os.chdir(f'{cwd}')
            print(res)
        except Exception as f:
            print("++++++++++++++++++++++++++++++++++++++++++")
            print(f"{f}")
            print("++++++++++++++++++++++++++++++++++++++++++")

if __name__ == '__main__':
    import plac; plac.call(main)

