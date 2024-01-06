FROM golang:1.21.1

WORKDIR /usr/local/src

COPY ./ ./

RUN go mod tidy
RUN go build -o ./my_app_start ./main.go

CMD ["./my_app_start"]
