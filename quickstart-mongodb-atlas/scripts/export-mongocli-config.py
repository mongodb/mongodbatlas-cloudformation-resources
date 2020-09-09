#!/usr/bin/env python3

# Export environment variables for a given mongocli
# profile or "default"
# To run this, source it:
# $ source <(./export-mongocli-config.py)
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
print(f"echo \"Exporting mongocli config for '{profile}'\"")
print(f"export ATLAS_PUBLIC_KEY={d['public_api_key']}")
print(f"export ATLAS_PRIVATE_KEY={d['private_api_key']}")
print(f"export ATLAS_ORG_ID={d['org_id']}")
print(f"export PUBLIC_KEY={d['public_api_key']}")
print(f"export PRIVATE_KEY={d['private_api_key']}")
print(f"export ORG_ID={d['org_id']}")
print("env | grep -E 'PUBLIC_KEY|PRIVATE_KEY|ORG_ID'")
