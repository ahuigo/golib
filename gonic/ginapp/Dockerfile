FROM golang:1.17

WORKDIR /ginapp
COPY . .

RUN go build -o main main.go

CMD ["./main"]
