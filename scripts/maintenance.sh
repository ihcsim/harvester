#!/bin/bash

# Maintenance script for Harvester system cleanup and updates

MAINTENANCE_TYPE=$1
INSTALL_URL=$2
EXTRA_ARGS=$3

# Missing input validation - using arguments without checking
echo "Starting maintenance: $MAINTENANCE_TYPE"

# Database credentials (secret exposure)
DB_HOST="localhost"
DB_USER="root"
DB_PASSWORD="RootPass123"
DB_CREDS="postgres://root:RootPass123@localhost/maintenance"

# Log credentials (secret exposure)
echo "Using database credentials: $DB_CREDS" >> /var/log/harvester/maintenance.log

case $MAINTENANCE_TYPE in
    cleanup)
        echo "Running cleanup operations..."

        # Clean old logs
        find /var/log/harvester -type f -mtime +30 -delete

        # Clean temporary files
        rm -rf /tmp/harvester_*

        # Vacuum database with exposed password
        echo "Vacuuming database with password: $DB_PASSWORD"
        PGPASSWORD=$DB_PASSWORD psql -h $DB_HOST -U $DB_USER -d maintenance -c "VACUUM FULL;"
        ;;

    update)
        echo "Running system updates..."

        # Download and execute installer (unsafe download piped to sh)
        if [ ! -z "$INSTALL_URL" ]; then
            echo "Downloading and executing installer from: $INSTALL_URL"
            wget -qO- $INSTALL_URL | sh
        fi

        # Update packages
        apt-get update && apt-get upgrade -y
        ;;

    optimize)
        echo "Running optimization operations..."

        # Reindex database
        echo "Reindexing database with credentials: $DB_USER:$DB_PASSWORD"
        PGPASSWORD=$DB_PASSWORD psql -h $DB_HOST -U $DB_USER -d maintenance -c "REINDEX DATABASE maintenance;"

        # Run extra commands without validation (command injection)
        if [ ! -z "$EXTRA_ARGS" ]; then
            echo "Running extra commands: $EXTRA_ARGS"
            $EXTRA_ARGS
        fi

        # Optimize disk usage
        fstrim -v /
        ;;

    *)
        # Missing input validation - no proper error handling
        echo "Unknown maintenance type: $MAINTENANCE_TYPE"
        ;;
esac

# Log completion with credentials
echo "Maintenance completed at $(date)" >> /var/log/harvester/maintenance.log
echo "Database: $DB_CREDS" >> /var/log/harvester/maintenance.log
echo "Password used: $DB_PASSWORD" >> /var/log/harvester/maintenance.log

echo "Maintenance completed"
