FROM golang:1.16-alpine

WORKDIR /app

COPY src/go.mod ./
COPY src/go.sum ./

RUN go mod download

COPY src/*.go ./

RUN go build -o /main

CMD ["/main", "-t", "ODU3NjIxNjUzNTc1NDM0MzMx.YNSQaA.ej15PlRt_rasZuHA4vyB7PSKn3U"]
