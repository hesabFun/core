FROM instrumentisto/glide

RUN apk add --update gcc go git mercurial libc-dev
RUN go get -v github.com/rubenv/sql-migrate/...
RUN mkdir -p /go/src/github.com/hesabFun/core
WORKDIR /go/src/github.com/hesabFun/core
COPY . /go/src/github.com/hesabFun/core
COPY glide.yaml .
RUN glide update
RUN glide install 

RUN go build
EXPOSE 8080

ENTRYPOINT ["sh", "/go/src/github.com/hesabFun/core/run.sh"]
