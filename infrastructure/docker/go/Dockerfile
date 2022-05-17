# 2020/10/14最新versionを取得
FROM golang:1.15.2-alpine
# アップデートとgitのインストール！！
RUN apk update && apk add git
RUN apk update && apk add curl
RUN go get -u gorm.io/gorm && \ 
    go get -u gorm.io/driver/mysql 
RUN go get github.com/gin-gonic/gin

WORKDIR /work/backend
