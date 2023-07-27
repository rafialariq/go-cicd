FROM golang:1.17-alpine

WORKDIR /app

COPY go.mod ./
COPY *.go ./
COPY pkg ./pkg

RUN go build -o /go-cicd

EXPOSE 3000

CMD ["/go-cicd"]
