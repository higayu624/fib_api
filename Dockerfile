FROM golang:1.18-alpine

# 1
ENV GOPATH /go
ENV GO111MODULE on
RUN GOOS=linux GOARCH=arm64
ENV TZ Asia/Tokyo

WORKDIR /go/src/fib_api
COPY src/ .

EXPOSE 5000

# RUN go mod init fib_api
RUN apk upgrade --update && apk --no-cache add git
RUN go get -u github.com/cosmtrek/air@latest && go build -o /go/bin/air github.com/cosmtrek/air

# air -c [tomlファイル名] // 設定ファイルを指定してair実行(WORKDIRに.air.tomlを配置しておくこと)
CMD ["air", "-c", ".air.toml"]