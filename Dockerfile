FROM golang:1.9

USER nobody

RUN mkdir -p /go/src/github.com/hesabFun/core
WORKDIR /go/src/github.com/hesabFun/core

COPY . /go/src/github.com/hesabFun/core
RUN curl https://glide.sh/get | sh && \
        glide install && \
        go get github.com/rubenv/sql-migrate/... && \
#        sql-migrate up && \
#        sql-migrate up -env=seed && \
        cp -r .env.example .env && \
        go build

CMD ["go", "run"]