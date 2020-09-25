FROM golang:1.15.2-alpine3.12
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go mod download
RUN go build -o out/bin/deliverymuch-challenge cmd/api/main.go
CMD ["/app/main"]