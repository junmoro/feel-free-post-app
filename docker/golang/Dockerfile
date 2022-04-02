FROM golang:alpine

COPY ./go-app /app/go

WORKDIR /app/go

RUN go build -o main . 

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

EXPOSE 8000

CMD ["/app/go/main"]