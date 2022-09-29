FROM ubuntu:focal
COPY cos-uploader-linux-amd64 /bin/
RUN apt update && \
    apt install zip curl ca-certificates -y --no-install-recommends && \
    apt clean && \
    rm -rf /var/lib/apt/lists/*
