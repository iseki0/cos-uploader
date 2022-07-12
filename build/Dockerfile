FROM golang AS builder

WORKDIR /usr/src/app
COPY . .
RUN go build -trimpath -ldflags '-extldflags "-static" -s -w' -o cos-uploader .

FROM ubuntu
COPY --from=builder /usr/src/app/cos-uploader /bin/
RUN apt update && \
    apt install zip curl -y --no-install-recommends && \
    apt clean && \
    rm -rf /var/lib/apt/lists/*
