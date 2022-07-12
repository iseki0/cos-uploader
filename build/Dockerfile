FROM golang AS builder

WORKDIR /usr/src/app
COPY . .
RUN go build -trimpath -ldflags '-extldflags "-static" -s -w' -o cos-uploader .

FROM ubuntu
COPY --from=builder cos-uploader /bin/
RUN apt update && \
    apt install add zip curl -y --no-install-recommend && \
    apt clean && \
    rm -rf /var/lib/apt/lists/*
