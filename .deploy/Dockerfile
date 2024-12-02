FROM golang:1.23-alpine as builder

ENV APP_UID=1000
ENV APP_GID=1100
ENV APP_USER=apps
ENV APP_HOME=/apps
ENV SRC_HOME=$APP_HOME/src

ENV APP_SERVICE date-me
ENV APP_PORT 10501

# set time zone, as default this base file will
# use Asia/Jakarta as timezone
RUN apk add --no-cache tzdata && \
        cp /usr/share/zoneinfo/Asia/Jakarta /etc/localtime && \
        echo "Asia/Jakarta" > /etc/timezone

# install most used linux command
RUN apk add --no-cache nano zip net-tools wget curl iputils bash

# create non-root user
RUN addgroup -g $APP_GID $APP_USER && \
        adduser -u $APP_UID -G $APP_USER -h $APP_HOME -s /bin/bash -D $APP_USER

# create home directory
# it should place at : /apps/src
RUN mkdir -p ${SRC_HOME}
RUN chown --recursive ${APP_USER}:${APP_USER} ${APP_HOME}

# copy current directory into image file
USER ${APP_USER}
WORKDIR ${SRC_HOME}

COPY --chown=${APP_USER} . .

# running the unit test before building the images
WORKDIR ${SRC_HOME}/cmd

RUN go mod tidy && go mod verify
RUN go build -o ${APP_HOME} ${SRC_HOME}/cmd/main.go

# rebuild image to compress image size as much as possible
# FROM gcr.io/distroless/base-debian10
FROM alpine:3.13

ENV APP_USER apps
ENV APP_HOME /apps
ENV SRC_HOME ${APP_HOME}/src

# set time zone, as default this base file will
# use Asia/Jakarta as timezone
RUN apk add --no-cache tzdata
ENV TZ=Asia/Jakarta

COPY --from=builder ${APP_HOME}/main ${APP_HOME}/main

WORKDIR ${APP_HOME}

EXPOSE ${APP_PORT}
CMD ["./main"]
