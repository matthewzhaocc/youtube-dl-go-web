FROM alpine:latest AS builder
RUN mkdir /app
WORKDIR /app
COPY . .
RUN setup.sh

FROM busybox:latest
COPY --from=0 /app/youtube-dl-go-web .
RUN mkdir downloads/
CMD ./youtube-dl-go-web