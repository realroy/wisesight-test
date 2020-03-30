FROM golang:1.14-alpine3.11 AS build

RUN apk update && apk add gcc git 

WORKDIR /usr/app

COPY . .

RUN go build main.go

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:3.11

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=build /usr/app .

CMD ["./main"]