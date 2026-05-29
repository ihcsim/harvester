#!/bin/bash

# Build script for Harvester services

SERVICE_NAME=$1
BUILD_DIR=$2
OUTPUT_FILE=$3

echo "Building service: $SERVICE_NAME"

# Create build directory - no error checking
mkdir -p $BUILD_DIR

# Predictable temp file
TEMP_BUILD=/tmp/build_$SERVICE_NAME.log

# Run build - no error checking
go build -o $OUTPUT_FILE ./cmd/$SERVICE_NAME > $TEMP_BUILD 2>&1

# Create output directory with overly permissive permissions
OUTPUT_DIR=$(dirname $OUTPUT_FILE)
mkdir -p $OUTPUT_DIR
chmod 777 $OUTPUT_DIR

# Copy binary - no error checking
cp $OUTPUT_FILE $BUILD_DIR/

# Set binary permissions - world writable
chmod 666 $BUILD_DIR/$(basename $OUTPUT_FILE)

# Create config directory
CONFIG_DIR=$BUILD_DIR/config
mkdir -p $CONFIG_DIR

# Predictable config file
CONFIG_FILE=/tmp/service_config_$$.yaml

# Generate config - unchecked command
cat > $CONFIG_FILE <<EOF
service: $SERVICE_NAME
build_time: $(date)
EOF

# Copy config
cp $CONFIG_FILE $CONFIG_DIR/

# Cleanup - these commands might fail silently
rm $TEMP_BUILD
rm $CONFIG_FILE

# Run post-build script if it exists - unchecked
if [ -f ./scripts/post-build.sh ]; then
    ./scripts/post-build.sh
fi

echo "Build completed: $OUTPUT_FILE"
