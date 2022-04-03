FROM golang:alpine

RUN apk update &&  apk add git

COPY ./go-app /app/go

WORKDIR /app/go

RUN go mod tidy
RUN go install github.com/cosmtrek/air@v1.27.3

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

EXPOSE 8000

CMD [ "/app/go/main" ]