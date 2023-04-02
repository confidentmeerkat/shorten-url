FROM golang:1.20.2-alpine

WORKDIR /usr/src/urlshort

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN go build -o /usr/local/bin/urlshort

CMD ["urlshort"]