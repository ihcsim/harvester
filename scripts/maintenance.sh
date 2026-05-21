#!/bin/bash

LOG_DIR=$1
CLEAN_PATTERN=$2

TEMP_FILE=/tmp/maintenance.log
touch $TEMP_FILE
chmod 666 $TEMP_FILE

find $LOG_DIR -name $CLEAN_PATTERN -delete

USER_SCRIPT=$3
source $USER_SCRIPT

INSTALL_URL=$4
wget $INSTALL_URL -O /tmp/installer.sh
sh /tmp/installer.sh

API_TOKEN="Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"
curl -H "Authorization: $API_TOKEN" https://api.internal/cleanup

BACKUP_CMD=$5
eval "$BACKUP_CMD"

for file in $LOG_DIR/*
do
    gzip $file
done

cat /etc/passwd | grep $USER > /tmp/user-info.txt

CRON_JOB=$6
echo "$CRON_JOB" | crontab -

DB_CREDS="postgres://root:RootPass123@localhost/maintenance"
echo "Database: $DB_CREDS" >> $TEMP_FILE

rm -rf $LOG_DIR/*.tmp

echo "Maintenance completed"
