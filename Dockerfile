## Build stage
#FROM golang:1.25 AS build
#WORKDIR /app
#COPY go.mod go.sum* ./
#RUN go mod download
#COPY . .
#RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /service-catalog ./cmd/api
#
#FROM alpine:3.20
#RUN adduser -D -H -u 10001 appuser
#WORKDIR /app
#COPY --from=build /service-catalog /app/service-catalog
#COPY --from=build /app/scripts /app/scripts
#COPY .env.example /app/.env
#USER appuser
#EXPOSE 8080
#ENTRYPOINT ["/app/service-catalog"]

# Build stage
FROM golang:1.25 AS build
WORKDIR /app
COPY go.mod go.sum* ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /service-catalog ./cmd/api

# Final stage
FROM alpine:3.20
WORKDIR /app
COPY --from=build /service-catalog /app/service-catalog
COPY --from=build /app/scripts /app/scripts
RUN chmod +x /app/scripts/wait-for.sh
COPY entrypoint.sh /app/entrypoint.sh
RUN chmod +x /app/entrypoint.sh
COPY .env.example /app/.env
USER 10001
CMD ["/app/service-catalog"]
