FROM golang:bookworm AS download

RUN mkdir /app

ADD . /app

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

FROM download AS build

COPY .env /

RUN go build -o store_service cmd/main.go

RUN ["ls","/"]

FROM debian:bookworm

COPY --from=build /app/store_service /store_service

EXPOSE 5001

CMD ["/store_service"]