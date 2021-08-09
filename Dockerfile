FROM golang:1.16-alpine
MAINTAINER Sebastián Remaggi Flores

RUN apk add --no-cache gcc musl-dev

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download



COPY *.go ./

RUN go build -o /kafka-ms-poc

EXPOSE 8080

CMD [ "/kafka-ms-poc" ]

