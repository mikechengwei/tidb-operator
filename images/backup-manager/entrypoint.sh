#!/bin/sh
set -e

echo "Create rclone.conf file."
cat <<EOF > /tmp/rclone.conf
[s3]
type = s3
env_auth = false
provider =  ${S3_PROVIDER}
access_key_id = ${AWS_ACCESS_KEY_ID}
secret_access_key = ${AWS_SECRET_ACCESS_KEY:-$AWS_SECRET_KEY}
region = ${AWS_REGION}
acl = ${AWS_ACL}
endpoint = ${S3_ENDPOINT}
storage_class = ${AWS_STORAGE_CLASS}
[gcs]
type = google cloud storage
project_number = ${GCS_PROJECT_ID}
service_account_file = /tmp/google-credentials.json
object_acl = ${GCS_OBJECT_ACL}
bucket_acl = ${GCS_BUCKET_ACL}
location =  ${GCS_LOCATION}
storage_class = ${GCS_STORAGE_CLASS:-"COLDLINE"}
[azure]
type = azureblob
account = ${AZUREBLOB_ACCOUNT}
key = ${AZUREBLOB_KEY}
EOF

if [[ -n "${GCS_SERVICE_ACCOUNT_JSON_KEY:-}" ]]; then
    echo "Create google-credentials.json file."
    cat <<EOF > /tmp/google-credentials.json
    ${GCS_SERVICE_ACCOUNT_JSON_KEY}
EOF
else
    touch /tmp/google-credentials.json
fi

BACKUP_BIN=/tidb-backup-manager

# exec command
case "$1" in
    backup)
        shift 1
        echo "$BACKUP_BIN backup $@"
        exec $BACKUP_BIN backup "$@"
        ;;
    restore)
        shift 1
        echo "$BACKUP_BIN restore $@"
        exec $BACKUP_BIN restore "$@"
        ;;
    clean)
        shift 1
        echo "$BACKUP_BIN clean $@"
        exec $BACKUP_BIN clean "$@"
        ;;
    *)
        echo "Usage: $0 {backup|restore|clean}"
        echo "Now runs your command."
        echo "$@"

        exec "$@"
esac
