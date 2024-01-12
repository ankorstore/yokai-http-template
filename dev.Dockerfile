FROM golang:1.20-alpine

WORKDIR /app

CMD ["air", "-c", ".air.toml", "--", "server"]
