FROM golang:1.20-alpine

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . ./

RUN go build -o server

COPY main.go ./
RUN go build -o /verifier
CMD [ "/verifier" ]
