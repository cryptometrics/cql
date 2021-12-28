FROM golang:1.17-alpine

RUN go install golang.org/x/tools/cmd/goimports@latest
# FROM cytopia/goimports
