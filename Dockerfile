# Server build
FROM golang:1.19-alpine as app-builder

RUN apk add --no-cache \
    alpine-sdk

# Force the go compiler to use modules
ENV GO111MODULE=on

ARG APP_VERSION
ENV APP_VERSION=${APP_VERSION}

ADD . /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux go build -a -o dotenv-updater .

# Final image
FROM alpine:3.16

RUN apk add --no-cache \
    ca-certificates \
    alsa-utils \
    tzdata

# copy app files
COPY --from=app-builder /app/dotenv-updater /app/

ENTRYPOINT ["/app/dotenv-updater"]

EXPOSE 5000