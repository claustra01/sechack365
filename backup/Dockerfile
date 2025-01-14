FROM debian:bookworm-slim
ENV TZ=Asia/Tokyo

RUN apt update && \
    apt install -y postgresql-client cron bash curl

RUN curl -O https://dl.min.io/client/mc/release/linux-amd64/mc && \
    chmod +x mc && \
    mv mc /usr/local/bin/

COPY backup.sh /backup.sh
RUN chmod +x /backup.sh

ARG DB_HOST
ARG DB_PORT
ARG DB_USER
ARG DB_PASSWORD
ARG DB_NAME
ARG MINIO_HOST
ARG MINIO_PORT
ARG MINIO_ROOT_USER
ARG MINIO_ROOT_PASSWORD

RUN touch /var/log/cron.log && \
    echo "DB_HOST=${DB_HOST}" > /etc/cron.d/backup && \
    echo "DB_PORT=${DB_PORT}" >> /etc/cron.d/backup && \
    echo "DB_USER=${DB_USER}" >> /etc/cron.d/backup && \
    echo "DB_PASSWORD=${DB_PASSWORD}" >> /etc/cron.d/backup && \
    echo "DB_NAME=${DB_NAME}" >> /etc/cron.d/backup && \
    echo "MINIO_HOST=${MINIO_HOST}" >> /etc/cron.d/backup && \
    echo "MINIO_PORT=${MINIO_PORT}" >> /etc/cron.d/backup && \
    echo "MINIO_ROOT_USER=${MINIO_ROOT_USER}" >> /etc/cron.d/backup && \
    echo "MINIO_ROOT_PASSWORD=${MINIO_ROOT_PASSWORD}" >> /etc/cron.d/backup && \
    echo "TZ=${TZ}" >> /etc/cron.d/backup && \
    echo "0 4 * * * bash /backup.sh >> /var/log/cron.log 2>&1" >> /etc/cron.d/backup && \
    crontab /etc/cron.d/backup

CMD ["cron", "-f"]
