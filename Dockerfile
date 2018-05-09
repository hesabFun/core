FROM alpine:3.6

RUN mkdir -p /migrations
WORKDIR /migrations

COPY ./core /bin
COPY ./sql-migrate /bin
COPY ./migrations .
COPY ./seeds .
COPY ./dbconfig.yml .
COPY ./run .

EXPOSE 8080

CMD ["./run"]
