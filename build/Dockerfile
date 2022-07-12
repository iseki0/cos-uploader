FROM golang AS builder
RUN pwd && ls
RUN go build -ldflags '-extldflags "-static" -s -w' -o cos-uploader .
FROM alpine
COPY --from=builder cos-uploader /bin/
RUN apk add zip curl
