FROM golang:alpine AS builder

WORKDIR /app
RUN apk update && apk add --no-cache ca-certificates && update-ca-certificates
COPY ["go.mod", "go.sum", "./"]
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo ./cmd/workflow

FROM scratch

WORKDIR /app
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/workflow ./

EXPOSE 8080

CMD [ "./workflow" ]