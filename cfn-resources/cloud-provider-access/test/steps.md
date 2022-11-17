```
cat <<EOD >> role-trust-policy.json
{
"Version": "2012-10-17",
"Statement": [
{
"Effect": "Allow",
"Principal": {
"AWS": "arn:aws:iam::536727724300:root"
},
"Action": "sts:AssumeRole",
"Condition": {
"StringEquals": {
"sts:ExternalId": "b794cde3-22e9-4be2-ad3f-dcca4693bef4"
}
}
}
]
}
EOD

```

```
aws iam create-role \
 --role-name mongorole1 \
 --assume-role-policy-document file://role-trust-policy.json
```