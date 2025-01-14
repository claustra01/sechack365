FROM debian:bookworm-slim
ENV TZ=Asia/Tokyo

RUN apt update && \
    apt install -y postgresql-client cron bash curl

RUN curl -O https://dl.min.io/client/mc/release/linux-amd64/mc && \
    chmod +x mc && \
    mv mc /usr/local/bin/

COPY backup.sh /backup.sh
RUN chmod +x /backup.sh

RUN mkdir -p /etc/crontabs && \
    echo "* * * * * /backup.sh" > /etc/crontabs/root

CMD ["cron", "-f"]
