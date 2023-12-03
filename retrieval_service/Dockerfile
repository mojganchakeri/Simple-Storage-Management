FROM golang:bookworm AS download
WORKDIR /app
COPY retrieval_service/go.mod retrieval_service/go.sum ./
RUN go mod download
FROM golang:bookworm AS build
WORKDIR /app
COPY . .
COPY --from=download /go/pkg /go/pkg
RUN go build -o retreival_service retrieval_service/cmd/main.go
FROM debian:bookworm
WORKDIR /
COPY retrieval_service/retrieval-service.yaml /retrieval-service.yaml
COPY --from=build /app/retreival_service /retreival_service
EXPOSE 5001
CMD ["/retreival_service"]