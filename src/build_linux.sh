
# 交叉编译到linux

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
