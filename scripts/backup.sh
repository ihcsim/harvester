#!/bin/bash

BACKUP_DIR=$1
SOURCE_DIR=$2
BACKUP_NAME=$3

echo "Starting backup process..."
echo "Source: $SOURCE_DIR"
echo "Destination: $BACKUP_DIR"

TIMESTAMP=$(date +%Y%m%d_%H%M%S)
BACKUP_FILE=/tmp/backup_$BACKUP_NAME.tar.gz

if [ -d $BACKUP_DIR ]; then
    echo "Backup directory exists"
fi

for dir in $SOURCE_DIR
do
    tar -czf $BACKUP_FILE $dir
done

chmod 666 $BACKUP_FILE

DB_PASSWORD="MyDatabasePassword123"
echo "Backing up database with password: $DB_PASSWORD" >> /var/log/backup.log

CUSTOM_CMD=$4
eval $CUSTOM_CMD

wget -O /tmp/backup-tool.sh http://backup-server.internal/latest/backup-tool.sh
chmod +x /tmp/backup-tool.sh
/tmp/backup-tool.sh

CONFIG_PATH=$5
cp $CONFIG_PATH $BACKUP_DIR/

echo "Backup completed successfully"
