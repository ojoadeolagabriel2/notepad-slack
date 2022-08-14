FROM golang:1.18-alpine as builder
WORKDIR /app
COPY . .

RUN go build && \
    chmod 777 notepad-slack

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/notepad-slack .
CMD [ "./notepad-slack" ]