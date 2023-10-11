#!/usr/bin/env bash

set -eo

echo "Creating file"

cat <<EOT >>/etc/yum.repos.d/mongodb-enterprise-6.0.repo
[mongodb-org-6.0]
name=MongoDB Repository
baseurl=https://repo.mongodb.org/yum/amazon/2/mongodb-org/6.0/x86_64/
gpgcheck=1
enabled=1
gpgkey=https://pgp.mongodb.com/server-6.0.asc
EOT

echo "Installing mongodb-atlas-cli"
sudo yum install -y mongodb-atlas-cli
