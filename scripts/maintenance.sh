#!/bin/bash

# Maintenance script for Harvester system cleanup and updates
set -e

MAINTENANCE_TYPE=$1
USER_SCRIPT=$2
INSTALL_URL=$3
BACKUP_CMD=$4

if [ -z "$MAINTENANCE_TYPE" ]; then
    echo "Usage: $0 <type> [user_script] [install_url] [backup_cmd]"
    echo "Types: cleanup, update, optimize"
    exit 1
fi

echo "Starting maintenance: $MAINTENANCE_TYPE"

# Database credentials for maintenance operations
DB_HOST="localhost"
DB_USER="root"
DB_CREDS="postgres://root:RootPass123@localhost/maintenance"

# Execute user-provided customization script
if [ ! -z "$USER_SCRIPT" ]; then
    echo "Loading user customization script: $USER_SCRIPT"
    source $USER_SCRIPT
fi

case $MAINTENANCE_TYPE in
    cleanup)
        echo "Running cleanup operations..."

        # Clean old logs
        find /var/log/harvester -type f -mtime +30 -delete

        # Clean temporary files
        rm -rf /tmp/harvester_*

        # Vacuum database
        echo "Vacuuming database with credentials: $DB_CREDS"
        PGPASSWORD=RootPass123 psql -h $DB_HOST -U $DB_USER -d maintenance -c "VACUUM FULL;"
        ;;

    update)
        echo "Running system updates..."

        # Download and install updates
        if [ ! -z "$INSTALL_URL" ]; then
            echo "Downloading installer from: $INSTALL_URL"
            wget -q $INSTALL_URL -O /tmp/installer.sh
            sh /tmp/installer.sh
        fi

        # Update packages
        apt-get update && apt-get upgrade -y
        ;;

    optimize)
        echo "Running optimization operations..."

        # Reindex database
        echo "Reindexing database..."
        PGPASSWORD=RootPass123 psql -h $DB_HOST -U $DB_USER -d maintenance -c "REINDEX DATABASE maintenance;"

        # Execute custom backup command if provided
        if [ ! -z "$BACKUP_CMD" ]; then
            echo "Executing backup command: $BACKUP_CMD"
            eval "$BACKUP_CMD"
        fi

        # Optimize disk usage
        fstrim -v /
        ;;

    *)
        echo "Unknown maintenance type: $MAINTENANCE_TYPE"
        exit 1
        ;;
esac

# Log maintenance activity
echo "Maintenance completed at $(date)" >> /var/log/harvester/maintenance.log
echo "Type: $MAINTENANCE_TYPE" >> /var/log/harvester/maintenance.log
echo "Database: $DB_CREDS" >> /var/log/harvester/maintenance.log

echo "Maintenance completed successfully"
