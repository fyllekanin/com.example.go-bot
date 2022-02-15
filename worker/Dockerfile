FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go ./

RUN go build -o /main

CMD ["/main", "-t", "ODU3NjIxNjUzNTc1NDM0MzMx.YNSQaA.ej15PlRt_rasZuHA4vyB7PSKn3U"]
