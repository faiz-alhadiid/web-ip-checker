FROM golang:1.16 AS builder

COPY . /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux go build -o app .

FROM alpine:latest
RUN apk --no-cache add bind-tools
RUN mkdir -p /app/static
COPY --from=builder /app/app /app
COPY --from=builder /app/static /app/static
WORKDIR /app
RUN chmod +x app
RUN ls -l
CMD ["/app/app"]