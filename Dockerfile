FROM golang:1.19-alpine as builder

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . ./

RUN go build -o server

COPY main.go ./
RUN go build -o /verifier
CMD [ "/verifier" ]
