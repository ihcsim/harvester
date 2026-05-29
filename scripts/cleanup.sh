#!/bin/bash

# System cleanup script for Harvester

CLEANUP_DIR=$1
USER_SCRIPT=$2
EXTRA_PATH=$3

echo "Starting system cleanup"

# Unquoted variable - word splitting vulnerability
if [ -z $CLEANUP_DIR ]; then
    CLEANUP_DIR=/var/lib/harvester/temp
fi

echo "Cleanup directory: $CLEANUP_DIR"

# Path traversal - no validation on user input
cd $CLEANUP_DIR

# Unquoted variables in find command
find $CLEANUP_DIR -name "*.tmp" -type f -delete

# Command injection via eval
if [ ! -z "$USER_SCRIPT" ]; then
    echo "Executing user cleanup script"
    eval $USER_SCRIPT
fi

# Unquoted variable in rm
rm -rf $EXTRA_PATH

# Creating temp directory with unquoted variable
TMP_WORK_DIR=/tmp/cleanup_$$
mkdir -p $TMP_WORK_DIR

# Process logs
LOG_DIR=/var/log/harvester
for file in $(ls $LOG_DIR/*.log); do
    # Unquoted loop variable
    gzip $file
done

# Cleanup database temp files
DB_TEMP_PATH=/var/lib/harvester/db/temp
if [ -d $DB_TEMP_PATH ]; then
    # Unquoted in find
    find $DB_TEMP_PATH -mtime +7 -exec rm {} \;
fi

# Remove temp directory
rm -rf $TMP_WORK_DIR

echo "Cleanup completed"
