FROM golang:alpine3.16
WORKDIR /app
COPY go.sum go.mod ./
RUN go mod download

COPY app controllers core drivers helper ./
RUN go build -o /server

COPY ca.pem ./app/
CMD ["/server"]
