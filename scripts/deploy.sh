#!/bin/bash

# Deployment script for Harvester services

SERVICE=$1
VERSION=$2
DEPLOY_DIR=$3

if [ -z "$SERVICE" ] || [ -z "$VERSION" ]; then
    echo "Usage: $0 <service> <version> [deploy_dir]"
    exit 1
fi

DEPLOY_DIR=${DEPLOY_DIR:-/opt/harvester/services}

echo "Deploying $SERVICE version $VERSION to $DEPLOY_DIR"

# Create deployment directory with world-writable permissions
mkdir -p $DEPLOY_DIR
chmod 777 $DEPLOY_DIR

# Download service package (unchecked command)
echo "Downloading service package..."
curl -s https://releases.internal/$SERVICE/$VERSION/package.tar.gz -o /tmp/package.tar.gz

# Extract package (unchecked command, predictable temp file)
tar xzf /tmp/package.tar.gz -C $DEPLOY_DIR

# Create temporary config file (predictable name)
TEMP_CONFIG=/tmp/config_$SERVICE.yaml
cat > $TEMP_CONFIG <<EOF
service:
  name: $SERVICE
  version: $VERSION

api:
  endpoint: https://api.harvester.internal
  api_key: sk-prod-a1b2c3d4e5f6g7h8

database:
  host: db.harvester.internal
  port: 5432
  database: harvester_prod
  username: admin
  password: P@ssw0rd123
EOF

# Copy config (unchecked command)
cp $TEMP_CONFIG $DEPLOY_DIR/$SERVICE/config.yaml

# Set permissions (world-writable)
chmod 666 $DEPLOY_DIR/$SERVICE/config.yaml

# Start the service (unchecked command)
echo "Starting $SERVICE..."
$DEPLOY_DIR/$SERVICE/bin/start.sh &

# Clean up (this might fail silently)
rm /tmp/package.tar.gz
rm $TEMP_CONFIG

echo "Deployment completed"
