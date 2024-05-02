FROM golang:1.21.5-alpine as builder
RUN mkdir /app
WORKDIR /app
COPY . .
RUN go build -o transcoder .

# stage 2 migrations
FROM migrate/migrate as migrate

FROM alpine:latest as host

RUN apk update
RUN apk upgrade
RUN apk add --no-cache ffmpeg

RUN mkdir /scripts

COPY --from=builder /app/transcoder /
COPY --from=migrate /migrate /usr/local/bin/migrate
COPY ./migrations /migrations

COPY ./docker/entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

WORKDIR /

#ENV AWS_ACCESS_KEY_ID=test
#ENV AWS_SECRET_ACCESS_KEY=test
ENV AWS_REGION=us-east-1
ENV STORAGE_PATH=/tmp
ENV INPUT_S3_BUCKET=bucket-name
ENV OBJECT_NAME=test.mp4

ENV DBUSER="test"
ENV DBPASS="password"
ENV DBHOST="localhost"
ENV DBPORT="5432"
ENV DBNAME="test"
ENV SSLMODE="disable"

ENTRYPOINT ["/entrypoint.sh"]
CMD ["/transcoder"]