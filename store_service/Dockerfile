FROM golang:bookworm AS download
WORKDIR /app
COPY store_service/go.mod store_service/go.sum ./
RUN go mod download
FROM golang:bookworm AS build
WORKDIR /app
COPY . .
COPY --from=download /go/pkg /go/pkg
RUN go build -o store_service store_service/cmd/main.go
FROM debian:bookworm
WORKDIR /
COPY store_service/store-service.yaml /store-service.yaml
COPY --from=build /app/store_service /store_service
EXPOSE 5001
CMD ["/store_service"]
