FROM golang:1.21.5-alpine as builder
RUN mkdir /app
WORKDIR /app
COPY . .
RUN go build -o transcoder .

FROM alpine:latest as host

RUN apk update
RUN apk upgrade
RUN apk add --no-cache ffmpeg

RUN mkdir /scripts

COPY --from=builder /app/transcoder /
COPY resources/app/convert_video_cpu.sh /convert_video_cpu.sh

WORKDIR /

RUN chmod +x convert_video_cpu.sh

ENV AWS_ACCESS_KEY_ID=test
ENV AWS_SECRET_ACCESS_KEY=test
ENV AWS_REGION=us-east-1
ENV STORAGE_PATH=/tmp
ENV INPUT_S3_BUCKET=bucket-name
ENV OBJECT_NAME=test.mp4
ENV CONVERT_FFMPEG_SCRIPT_PATH=/convert_video_cpu.sh

CMD ["/transcoder"]