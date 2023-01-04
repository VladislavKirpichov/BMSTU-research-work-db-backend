FROM golang:alpine as builder
WORKDIR /app
COPY . .
RUN go build cmd/main.go

FROM alpine
COPY --from=builder /app/main executable
COPY --from=builder /app/.env .env
COPY --from=builder /app/configs /configs/
EXPOSE 8080
CMD [ "./executable" ]
