FROM golang:1.14-alpine AS build
RUN mkdir /app
WORKDIR /app
COPY forest-mock-data-receiver /app/
RUN ls -ltrh forest-mock-data-receiver
COPY forest-mock-data-receiver /usr/bin/
CMD ["/usr/bin/forest-mock-data-receiver"]
