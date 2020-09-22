#!/usr/bin/env python
import sys

BUCKET = sys.argv[1]

import boto3

s3 = boto3.resource('s3')
bucket = s3.Bucket(BUCKET)
print(f"bucket: {bucket}")
response=bucket.object_versions.delete()
print(f"bucket.object_versions.delete(): {response}")

# if you want to delete the now-empty bucket as well, uncomment this line:
response=bucket.delete()
print(f"bucket.delete(): {response}")
