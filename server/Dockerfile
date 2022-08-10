FROM golang:1.18

WORKDIR /go/src/moral-ambiguity

COPY . .

RUN go build -o bin/server ./main.go

CMD ["./bin/server"]