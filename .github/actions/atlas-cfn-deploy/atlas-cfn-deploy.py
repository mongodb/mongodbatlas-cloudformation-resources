#!/usr/bin/env python3
from datetime import datetime
from loguru import logger
import os
import subprocess
import sys
import yaml

tier_one_resources = ["project","cluster","database-user","ip-whitelist"]

tier_two_resources = [ "cloud-provider-snapshots", "network-peering",
                "cloud-provider-snapshot-restore-jobs",
               "encryption-at-rest",
               "network-container" ]

def try_log_error(res, ok_msg="OK"):
    if res:
        logger.error(res)
    else:
        logger.info(ok_msg)

def main(parameters: ('resource parameters','option','p')
         , region: ('region','option','r')
         , noun):
    """
    atlas-cfn-cli.py
    MongoDB Atlas AWS CloudFormation Custom Resource CLI
    This tool helps you deploy Atlas resources for your CF templates
    Set LOGURU_LEVEL to higher for more logs
     """

    TOOL = os.path.realpath(__file__)
    TOOL_HOME = os.path.dirname( TOOL )
    logger.info( f"TOOL_HOME listdir: {os.listdir(TOOL_HOME)}" )
    ATLAS_CFN_HOME = os.path.join( TOOL_HOME, "cfn-resources")
    if TOOL_HOME == "/github/workspace/.github/actions/atlas-cfn-deploy":
        logger.info(f"Did not detect \"cfn-resources\" folder in {ATLAS_TOOL_HOME}, checking running from Github Action.")
        ATLAS_CFN_HOME="/github/workspace/cfn-resources"
    logger.info(f"TOOL: {TOOL}")
    logger.info(f"TOOL_HOME: {TOOL_HOME}")
    logger.info(f"ATLAS_CFN_HOME: {ATLAS_CFN_HOME}")
    logger.info(f"atlas-cfn-deploy deploying MongoDB Atlas CFN Custom Resources to AWS region {region}.")
    logger.info("WARNING: This tool is in active development. Swim at your own risk.")

    check_res = os.listdir(ATLAS_CFN_HOME)
    if not len(check_res):
        logger.error(f"os.listdir('{ATLAS_CFN_HOME}'): {check_res}")
        logger.fatal("Can't find any cfn-resources to deploy, unable to process.")
        sys.exit(1)

    if parameters:
        params=dict( kvp.strip().split('=') for kvp in parameters.strip().split(',') )
    else:
        params={}

    noun=noun.lower()
    logger.info(f"noun:{noun}")
    logger.info(f"parameters:{parameters}")
    logger.info(f"params:{params}")
    res = os.chdir(f'{ATLAS_CFN_HOME}')
    try_log_error(res,f"Changed to {ATLAS_CFN_HOME}")


    if noun.startswith("all"):
        logger.debug("noun was all")
        res = tier_one_resources
        if noun=="all+":
            for r in tier_two_resources:
                res.append(r)
        if noun=="all-":
            res = tier_two_resources
        for resource in res:
            command = ["python3"
                       ,TOOL
                       ,f"{resource}"
                       ,"--region"
                       ,f"{region}"]
            if parameters:
                command.append(["--parameters",f"{parameters}"])
            logger.debug(f"resource:{resource}")
            logger.debug(f"command:{command}")
            logger.debug(f"res:{res}")
            res = subprocess.run(command)
            process = subprocess.Popen(command, stdout=subprocess.PIPE, stderr=subprocess.STDOUT)
            output, error = process.communicate()
            if error:
                logger.error(error)
            logger.info(output)
    else:
        submit_cmd=["cfn","submit", "-v", "--region", f"{region}","--set-default"]
        cwd = os.getcwd()
        try:
            res_path = os.path.join(f"{ATLAS_CFN_HOME}",noun)
            res = os.chdir(f'{res_path}')
            try_log_error(res,f"Changed to resource: {noun} in {res_path}")
            process = subprocess.Popen(["make"], stdout=subprocess.PIPE, stderr=subprocess.STDOUT)
            output, error = process.communicate()
            if error:
                logger.error(error)
            logger.info(output)

            process = subprocess.Popen(submit_cmd, stdout=subprocess.PIPE, stderr=subprocess.STDOUT)
            output, error = process.communicate()
            if error:
                logger.error(error)
            logger.info(output)

            res = os.chdir(f'{cwd}')
            try_log_error(res,f"Changed to directory: {cwd}")
            print(res)
        except Exception as f:
            logger.error(f)

if __name__ == '__main__':
    import plac; plac.call(main)

