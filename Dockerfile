FROM alpine:3.6

RUN mkdir -p /app
WORKDIR /app

COPY ./core .
COPY ./sql-migrate .
COPY ./migrations .
COPY ./seeds .
COPY ./dbconfig.yml .
#COPY ./run .

#RUN chmod +x ./run

EXPOSE 8080

ENTRYPOINT ["./sql-migrate up", "./sql-migrate up -env=seed", "./core"]

# Specify the default user for the Docker image to run as.
#USER 1001