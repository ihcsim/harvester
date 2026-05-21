#!/bin/bash

# Deploy service script
# This script is intentionally vulnerable for testing purposes

SERVICE_NAME=$1
DEPLOY_ENV=$2
CONFIG_FILE=$3

# Unquoted variables - can cause word splitting
echo "Deploying service: $SERVICE_NAME to environment: $DEPLOY_ENV"

# Command injection - using eval with user input
eval "export SERVICE_$SERVICE_NAME=active"

# Unquoted variable in command
mkdir /tmp/deploy-$SERVICE_NAME

# Path traversal - no validation of input paths
CONFIG_PATH=/etc/services/$CONFIG_FILE
cat $CONFIG_PATH > /tmp/service-config.txt

# Unsafe download without integrity check
curl -s http://releases.example.com/$SERVICE_NAME/latest.tar.gz | tar xz

# Predictable temp file instead of mktemp
TEMP_LOG=/tmp/deploy-$SERVICE_NAME.log
echo "Starting deployment..." > $TEMP_LOG

# World-writable file
touch /tmp/deploy-status
chmod 666 /tmp/deploy-status

# Missing input validation - using arguments without checking
SERVICE_PORT=$4
netstat -tuln | grep $SERVICE_PORT

# Unquoted variable in loop
for file in $CONFIG_FILE
do
    echo "Processing $file"
done

# Command injection via backticks
VERSION=`curl -s http://api.example.com/version?service=$SERVICE_NAME`

# Unchecked command - can fail silently
cp /tmp/service-config.txt /var/lib/services/$SERVICE_NAME.conf

# Secret exposure in logs
API_KEY="sk-1234567890abcdef"
echo "Using API key: $API_KEY for deployment"

# Using user input directly in command
SERVICE_CMD="systemctl start $SERVICE_NAME"
$SERVICE_CMD

# Path constructed from input without validation
BACKUP_PATH="../backups/$SERVICE_NAME"
mkdir -p $BACKUP_PATH

# Unsafe eval with variable
CUSTOM_CMD=$5
eval $CUSTOM_CMD

# Missing error handling
cd /opt/services/$SERVICE_NAME
./deploy.sh

echo "Deployment complete"
