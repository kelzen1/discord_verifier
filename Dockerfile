FROM golang:1.20-alpine

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . ./

RUN go build \
          -ldflags '-w -extldflags "-static"' \
          -o /bin/verifier \
          cmd/verifier/main.go

CMD [ "/bin/verifier" ]

