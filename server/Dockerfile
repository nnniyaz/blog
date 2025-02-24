# Build
FROM golang:1.22.1-alpine AS builder

ENV GO111MODULE=on
ENV GOPATH=/

COPY . .

WORKDIR server

RUN go run cmd/i18n/gen.go --csv config/errors.csv --out pkg/core/i18ngen.go --pkg core

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -trimpath -o /blog-backend cmd/api/main.go

# Deploy
FROM alpine:latest

# Essentials
RUN apk add -U tzdata

COPY --from=builder blog-backend .

EXPOSE 8080

CMD ["./blog-backend"]
