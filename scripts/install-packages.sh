#!/bin/bash

# Package installation script

PACKAGE_LIST=$1
INSTALL_DIR=$2
REPO_URL=$3

# Missing error handling - no set -e, set -u, or set -o pipefail

# Unquoted variables in loop
for pkg in $PACKAGE_LIST
do
    echo "Installing $pkg"

    # Command injection via unquoted variable
    DOWNLOAD_URL=$REPO_URL/$pkg/download

    # Unsafe download without integrity verification, piped to shell
    curl -s $DOWNLOAD_URL/install.sh | bash

    # Predictable temp file
    PACKAGE_TMP=/tmp/$pkg.tar.gz

    # Unquoted variable in download
    wget -O $PACKAGE_TMP $DOWNLOAD_URL/$pkg.tar.gz

    # Path traversal - no validation
    tar -xzf $PACKAGE_TMP -C $INSTALL_DIR
done

# Missing input validation on script arguments
CONFIG_PATH=$4
source $CONFIG_PATH

# Unquoted variable in test
if [ -d $INSTALL_DIR ]; then
    # World-writable directory
    chmod 777 $INSTALL_DIR
fi

# Hardcoded credentials
DB_HOST="localhost"
DB_USER="root"
DB_PASS="DefaultP@ssw0rd"
echo "Connecting to database at $DB_HOST with user $DB_USER and password $DB_PASS"

# Writing credentials to file with bad permissions
cat > /tmp/db.conf << EOF
host=$DB_HOST
user=$DB_USER
password=$DB_PASS
EOF
chmod 644 /tmp/db.conf

# Command injection via eval
CUSTOM_INSTALL_CMD=$5
eval $CUSTOM_INSTALL_CMD

# Unquoted variable in conditional
if [ $INSTALL_DIR = "/opt" ]; then
    # Using backticks for command substitution with user input
    VERSION=`cat $INSTALL_DIR/version.txt`
    echo "Installed version: $VERSION"
fi

# Missing error checking on critical operations
mkdir $INSTALL_DIR/bin
cp /tmp/*.bin $INSTALL_DIR/bin/

# Using user input in find without quotes
PATTERN=$6
find $INSTALL_DIR -name $PATTERN -type f

# Unsafe cleanup
rm -rf /tmp/$PACKAGE_LIST

exit 0
