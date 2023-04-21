FROM golang:1.20-alpine AS builder

LABEL maintainer="贰拾壹 (https://github.com/er10yi)"

ENV CGO_ENABLED 0
ENV GOPROXY https://goproxy.cn,direct

#RUN apk update --no-cache && apk add --no-cache tzdata
RUN apk add --no-cache tzdata

WORKDIR /build

COPY go.mod .
RUN go mod download

COPY . .

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -ldflags="-s -w" -o /build/meta .

FROM scratch

COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
ENV TZ Asia/Shanghai

WORKDIR /app

COPY --from=builder /build/meta /meta
COPY --from=builder /build/resource/model.conf /app/resource/model.conf
COPY --from=builder /build/resource/config.toml /app/resource/config.toml


ENTRYPOINT ["/meta"]
EXPOSE 9001