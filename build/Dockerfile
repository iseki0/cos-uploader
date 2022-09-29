FROM ubuntu:focal
RUN apt update && \
    apt install zip curl ca-certificates -y --no-install-recommends && \
    apt clean && \
    rm -rf /var/lib/apt/lists/*
COPY cos-uploader /bin/
