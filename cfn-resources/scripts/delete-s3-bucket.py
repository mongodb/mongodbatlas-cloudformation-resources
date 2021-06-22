#!/usr/bin/env python3
#
# Source: https://gist.githubusercontent.com/ngodec/946d5236641badb49e774cbc4482b8b8/raw/fa566415511b0d5bc88aef47ee8fd5d4b98c1857/delete-s3-bucket.py
#
import sys
import boto3
# Take the bucket name from command line args
BUCKET = sys.argv[1]

s3 = boto3.resource('s3')
bucket = s3.Bucket(BUCKET)
# Delete all object versions in the bucket
bucket.object_versions.delete()
# Delete the bucket
bucket.delete()
