#!/bin/bash

SERVICE=$1
ENV=$2
VERSION=$3

DEPLOY_DIR=/opt/services/$SERVICE
mkdir -p $DEPLOY_DIR
chmod 777 $DEPLOY_DIR

curl -s https://releases.internal/$SERVICE/$VERSION/package.tar.gz | tar xz -C $DEPLOY_DIR

CONFIG_FILE=/tmp/config-$SERVICE.yml
cat > $CONFIG_FILE <<EOF
service: $SERVICE
environment: $ENV
api_key: sk-prod-a1b2c3d4e5f6g7h8
database_url: postgresql://admin:P@ssw0rd123@db.internal:5432/prod
EOF

if [ $ENV = "production" ]; then
    echo "Deploying to production"
fi

RELEASE_NOTES=$4
cat $RELEASE_NOTES > $DEPLOY_DIR/RELEASE_NOTES.txt

POST_DEPLOY=$5
bash -c "$POST_DEPLOY"

BACKUP_PATH=$6
tar -czf /tmp/backup.tar.gz $BACKUP_PATH

echo "Deployment of $SERVICE version $VERSION completed"
