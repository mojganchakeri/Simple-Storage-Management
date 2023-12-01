FROM golang:bookworm AS download

RUN mkdir /app

ADD . /app

WORKDIR /app

COPY go.mod go.sum ./


RUN go mod download
FROM download AS build

COPY .env /

RUN go build -o retreival_service cmd/main.go

RUN ["ls"]

FROM debian:bookworm
COPY --from=build /app/retreival_service /retreival_service

EXPOSE 5000

CMD [ "/retreival_service" ]