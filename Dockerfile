FROM golang:1.14.1-alpine AS builder

ADD . /go/src/hesab.fun/core

RUN apk add --no-cache --virtual .build-deps git gcc g++ libc-dev make \
    && apk add --no-cache ca-certificates bash \
    && cd /go/src/hesab.fun/core && make all \
    && apk del .build-deps

FROM alpine:3.6

ARG APP_NAME
ARG APP_PREFIX

COPY --from=builder /go/src/hesab.fun/core/bin/${APP_PREFIX}server /bin/server
COPY --from=builder /go/src/hesab.fun/core/bin/${APP_PREFIX}migration /bin/migration
ADD scripts/server.sh /bin/server.sh
ADD scripts/migration.sh /bin/migration.sh
ADD scripts/$APP_NAME/Procfile /bin/Procfile
ADD scripts/$APP_NAME/CHECKS /bin/CHECKS
ADD scripts/$APP_NAME/app.json /bin/app.json
RUN chmod a+x /bin/server.sh /bin/migration.sh

EXPOSE 80

WORKDIR /bin
