FROM golang:bookworm AS download
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
FROM golang:bookworm AS build
WORKDIR /app
COPY . ./
RUN ["ls"]
COPY --from=download /go/pkg /go/pkg
RUN go build -o retrieval_service cmd/main.go
FROM debian:bookworm
WORKDIR /
COPY retrieval-service.yaml /retrieval-service.yaml
COPY --from=build /app/retrieval_service /retrieval_service
EXPOSE 5000
CMD ["./retrieval_service"]