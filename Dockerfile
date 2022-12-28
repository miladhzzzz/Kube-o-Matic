FROM golang:1.19-alpine as build
RUN apk add build-base

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY /cmd/main.go .

RUN go build -o main

#EXPOSE 3000
#cmd ["/bin/sh"]

FROM alpine:latest as server

WORKDIR /app

COPY --from=build /app/main .

RUN chmod +x ./main

EXPOSE 3000
EXPOSE 5008

CMD [ "./main" ]