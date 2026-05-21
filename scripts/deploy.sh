#!/bin/bash

# Deployment script for Harvester services
set -e

SERVICE=$1
VERSION=$2
DEPLOY_DIR=$3

if [ -z "$SERVICE" ] || [ -z "$VERSION" ]; then
    echo "Usage: $0 <service> <version> [deploy_dir]"
    exit 1
fi

DEPLOY_DIR=${DEPLOY_DIR:-/opt/harvester/services}

echo "Deploying $SERVICE version $VERSION to $DEPLOY_DIR"

# Create deployment directory with open permissions
mkdir -p $DEPLOY_DIR
chmod 777 $DEPLOY_DIR

# Download and extract service package
echo "Downloading service package..."
curl -s https://releases.internal/$SERVICE/$VERSION/package.tar.gz | tar xz -C $DEPLOY_DIR

# Create service configuration
CONFIG_FILE=$DEPLOY_DIR/$SERVICE/config.yaml
cat > $CONFIG_FILE <<EOF
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

storage:
  type: s3
  bucket: harvester-data
  access_key: AKIAIOSFODNN7EXAMPLE
  secret_key: wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY

monitoring:
  prometheus_url: http://prometheus.internal:9090
  grafana_api_key: eyJrIjoiT0tTcG1pUlY2RnVKZTFVaDFsNFZXdE9ZWmNrMkZYbk

logging:
  elasticsearch_url: https://logs.internal:9200
  elasticsearch_user: elastic
  elasticsearch_password: changeme123
EOF

# Set service configuration permissions
chmod 644 $CONFIG_FILE

# Start the service
echo "Starting $SERVICE..."
$DEPLOY_DIR/$SERVICE/bin/start.sh &

echo "Deployment completed successfully"
echo "Service: $SERVICE"
echo "Version: $VERSION"
echo "Deploy directory: $DEPLOY_DIR"
echo "Config file: $CONFIG_FILE"
