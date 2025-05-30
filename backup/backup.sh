#!/bin/bash

set -e

export PGPASSWORD=$DB_PASSWORD

MINIO_ALIAS="myminio"
MINIO_BUCKET="backup"
/usr/local/bin/mc alias set $MINIO_ALIAS http://$MINIO_HOST:$MINIO_PORT $MINIO_ROOT_USER $MINIO_ROOT_PASSWORD

BACKUP_DIR=/tmp/backups
mkdir -p $BACKUP_DIR

DATE=$(date +\%Y-\%m-\%d_\%H-\%M-\%S)
BACKUP_NAME=backup_$DATE.dump

echo "Backup started at $DATE"
pg_dump -h $DB_HOST -p $DB_PORT -U $DB_USER -F c $DB_NAME > $BACKUP_DIR/$BACKUP_NAME
/usr/local/bin/mc cp $BACKUP_DIR/$BACKUP_NAME $MINIO_ALIAS/$MINIO_BUCKET/$BACKUP_NAME
echo "Backup finished: $MINIO_BUCKET/$BACKUP_NAME"
