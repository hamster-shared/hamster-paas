FROM golang:1.20.3-alpine3.16 as builder

WORKDIR /root/server
COPY . .

# swagger
ENV GOPROXY https://goproxy.cn,direct
#RUN go install github.com/swaggo/swag/cmd/swag@latest
#RUN swag init

RUN go build .


FROM debian:11

WORKDIR /app

COPY --from=builder /root/server/hamster-paas .

VOLUME /var/log/nginx
RUN apt-get update && apt-get install -y ca-certificates

EXPOSE 9898
ENTRYPOINT ["./hamster-paas"]
