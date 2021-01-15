FROM golang:alpine as buider
USER root
WORKDIR /app

COPY . .

RUN mkdir build
RUN go mod download; GOOS=linux GOARCH=amd64 go build main.go; mv main ./build

FROM alpine:latest

USER root

COPY --from=buider /app/build /app/build

WORKDIR /app/build

ENV LISTEN_PORT=8080

EXPOSE ${LISTEN_PORT}

ENTRYPOINT [ "./main" ]