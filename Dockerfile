FROM golang:1.14.1-alpine AS builder

ADD . /go/src/hesab.fun/core

RUN apk add --no-cache --virtual .build-deps git gcc g++ libc-dev make \
    && apk add --no-cache ca-certificates bash \
    && cd /go/src/hesab.fun/core && make all \
    && apk del .build-deps

FROM alpine:3.6

COPY --from=builder /go/src/hesab.fun/core/bin/qserver /bin/server
COPY --from=builder /go/src/hesab.fun/core/bin/qmigration /bin/migration
ADD scripts/entrypoint.sh /bin/entrypoint.sh
RUN chmod a+x /bin/entrypoint.sh

EXPOSE 80

WORKDIR /bin

CMD ["/bin/entrypoint.sh"]
