#!/bin/bash

# Backup data script
# This script is intentionally vulnerable for testing purposes

# Missing set -e - commands can fail silently
# Missing set -u - undefined variables won't error
# Missing set -o pipefail - pipeline failures can be hidden

BACKUP_DIR=$1
SOURCE_DIR=$2

# Missing argument validation
# $1 and $2 could be empty or malicious

# Predictable temp file names - race condition vulnerability
TEMP_LIST=/tmp/backup-list.txt
TEMP_ARCHIVE=/tmp/backup.tar.gz
LOG_FILE=/tmp/backup.log

# Unquoted variables
echo Backing up from $SOURCE_DIR to $BACKUP_DIR

# Create world-writable directory
mkdir -p $BACKUP_DIR
chmod 777 $BACKUP_DIR

# Unquoted variable in find command
find $SOURCE_DIR -type f > $TEMP_LIST

# Command injection through filename
while read filename
do
    # Unquoted variable in command
    cp $filename $BACKUP_DIR/
done < $TEMP_LIST

# Unsafe download without verification
BACKUP_TOOL_URL=$3
wget $BACKUP_TOOL_URL -O /tmp/backup-tool.sh
chmod +x /tmp/backup-tool.sh
/tmp/backup-tool.sh

# Using variable in tar command without quotes
tar -czf $TEMP_ARCHIVE $BACKUP_DIR

# Database password in plaintext
DB_PASSWORD="MySecretPassword123"
echo "Backing up database with password: $DB_PASSWORD" >> $LOG_FILE

# Command substitution without quotes
TIMESTAMP=$(date +%Y%m%d)
BACKUP_NAME=backup-$TIMESTAMP

# Missing validation - path traversal possible
RESTORE_PATH=$4
tar -xzf $TEMP_ARCHIVE -C $RESTORE_PATH

# Unsafe eval
BACKUP_OPTS=$5
eval tar $BACKUP_OPTS

# Writing credentials to world-readable file
echo "db_user=admin" > /tmp/db-config
echo "db_pass=$DB_PASSWORD" >> /tmp/db-config
chmod 644 /tmp/db-config

# Using external input in command without sanitization
EXCLUDE_PATTERN=$6
find $SOURCE_DIR -name $EXCLUDE_PATTERN -delete

# Missing error checking
rsync -a $SOURCE_DIR/ $BACKUP_DIR/
rm -rf $TEMP_LIST

# Unsafe use of variables in conditions
if [ $BACKUP_DIR = "/data" ]; then
    echo "Backing up production data"
fi

# No cleanup on error
# If script fails, temp files remain

echo "Backup complete"
exit 0
