FROM alpine:3.6

RUN mkdir -p /app
WORKDIR /app

COPY ./core /app
COPY ./sql-migrate /app
COPY ./migrations /app/migrations
COPY ./seeds /app/seeds
COPY ./dbconfig.yml /app
COPY ./run.sh /app

EXPOSE 80

CMD ["sh", "/app/run.sh"]