FROM golang:1.18

# metadata for better organization
LABEL app="go-moral-ambiguity"
LABEL environment="development"
LABEL maintainer="Josh <defidunce@proton.me>"

WORKDIR /go/src/moral-ambiguity

COPY . .

RUN go build -o bin/server ./cmd/main.go

CMD ["./bin/server"]