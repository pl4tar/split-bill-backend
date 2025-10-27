FROM golang:1.25-alpine

RUN mkdir /app

ADD . /app/

WORKDIR /app/cmd/app

RUN go build -o main .

CMD ["./main"]