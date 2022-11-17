# build stage
FROM golang:alpine AS build
WORKDIR /go/src/app
COPY . .
# RUN go mod init eclaim-api
RUN apk add git
RUN go mod tidy
# RUN apk add --no-cache git
RUN GOOS=linux go build -o ./bin/api ./main.go
# final stage
FROM alpine:latest
# RUN apk add --no-cache git
RUN apk update && apk add --no-cache tzdata
ENV TZ Asia/Jakarta
WORKDIR /usr/app
COPY --from=build /go/src/app/bin /go/bin
# COPY --from=build /go/src/app/ ./
EXPOSE 3030
ENTRYPOINT /go/bin/api