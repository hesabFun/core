FROM alpine:3.6

RUN mkdir -p /app
WORKDIR /app

COPY ./core /app
COPY ./sql-migrate /app
COPY ./migrations /appmigrations
COPY ./seeds /app/seeds
COPY ./dbconfig.yml /app
COPY ./run.sh /app

#RUN chmod +x ./run

EXPOSE 8080

ENTRYPOINT ["sh", "/app/run.sh"]

# Specify the default user for the Docker image to run as.
#USER 1001