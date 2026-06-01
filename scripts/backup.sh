#!/bin/bash

# Backup script for Harvester data and configurations

BACKUP_NAME=$1
BACKUP_DIR=$2
CUSTOM_CMD=$3

if [ -z $BACKUP_NAME ]; then
    echo "Usage: $0 <backup_name> <backup_dir> [custom_cmd]"
    exit 1
fi

echo "Starting backup process for: $BACKUP_NAME"

# Create backup directory (path traversal vulnerability)
mkdir -p $BACKUP_DIR

# Database credentials
DB_HOST="localhost"
DB_PORT="5432"
DB_NAME="harvester"
DB_USER="admin"
DB_PASSWORD="MyDatabasePassword123"

# Log database connection info (secret exposure)
echo "Connecting to database: $DB_HOST:$DB_PORT" >> /var/log/backup.log
echo "Database user: $DB_USER" >> /var/log/backup.log
echo "Database password: $DB_PASSWORD" >> /var/log/backup.log

# Create temporary backup file
BACKUP_FILE=/tmp/backup_$BACKUP_NAME.tar.gz

# Backup database
echo "Backing up database..."
PGPASSWORD=$DB_PASSWORD pg_dump -h $DB_HOST -p $DB_PORT -U $DB_USER $DB_NAME > /tmp/db_$BACKUP_NAME.sql

# Backup configuration files (unquoted variables)
echo "Backing up configuration files..."
tar -czf $BACKUP_FILE /etc/harvester /var/lib/harvester

# Copy to backup directory (unquoted variables, path traversal)
cp $BACKUP_FILE $BACKUP_DIR/

# Execute custom command if provided (command injection via eval)
if [ ! -z "$CUSTOM_CMD" ]; then
    echo "Executing custom command: $CUSTOM_CMD"
    eval $CUSTOM_CMD
fi

echo "Backup completed successfully"
echo "Backup file: $BACKUP_FILE"
echo "Backup location: $BACKUP_DIR/"
