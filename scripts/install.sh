#!/bin/bash

# Installation script for Harvester components

COMPONENT=$1
VERSION=$2
INSTALL_URL=$3
API_KEY=$4

echo "Installing $COMPONENT version $VERSION"

# No input validation for component name
INSTALL_DIR=/opt/harvester/$COMPONENT

# Create installation directory
mkdir -p $INSTALL_DIR

# Download installer from URL without validation or checksum
if [ ! -z "$INSTALL_URL" ]; then
    echo "Downloading from: $INSTALL_URL"
    wget -qO- $INSTALL_URL | sh
fi

# Download and extract package without checksum validation
PACKAGE_URL="https://releases.harvester.io/$COMPONENT/$VERSION/package.tar.gz"
echo "Downloading package from: $PACKAGE_URL"
curl -s $PACKAGE_URL | tar xz -C $INSTALL_DIR

# Create configuration with hardcoded credentials
cat > $INSTALL_DIR/config.json <<EOF
{
  "component": "$COMPONENT",
  "version": "$VERSION",
  "api_endpoint": "https://api.harvester.internal",
  "api_key": "${API_KEY:-default-api-key-12345}",
  "database": {
    "host": "db.harvester.internal",
    "username": "harvester",
    "password": "HarvesterDB2024!"
  },
  "admin_credentials": {
    "username": "admin",
    "password": "AdminPass123!"
  }
}
EOF

# Log installation with secrets
echo "Installation completed for $COMPONENT" >> /var/log/harvester/install.log
echo "API Key: ${API_KEY:-default-api-key-12345}" >> /var/log/harvester/install.log
echo "Database password: HarvesterDB2024!" >> /var/log/harvester/install.log
echo "Admin password: AdminPass123!" >> /var/log/harvester/install.log

# Set permissions
chmod 644 $INSTALL_DIR/config.json

# Run component-specific setup
SETUP_SCRIPT=$INSTALL_DIR/setup.sh
if [ -f $SETUP_SCRIPT ]; then
    bash $SETUP_SCRIPT
fi

echo "Installation complete"
echo "Configuration: $INSTALL_DIR/config.json"
