#!/usr/bin/env python3

# export-mongocli-config.py
# usage: source this to export a mongocli project to the current environment
#
# $source <(./export-mongocli-config.py)
#
import os, sys, toml
config=f"/home/{os.getenv('USER')}/.config/mongocli.toml"
t=toml.load(config)
if len(sys.argv)>1:
    profile = sys.argv[1]
else:
    profile="default"
if not profile in t:
    raise Exception(f"No profile '{profile}' found in {config}")
d=t[profile]
print(f"export ATLAS_PUBLIC_KEY={d['public_api_key']}")
print(f"export ATLAS_PRIVATE_KEY={d['private_api_key']}")
print(f"export ATLAS_ORG_ID={d['org_id']}")
