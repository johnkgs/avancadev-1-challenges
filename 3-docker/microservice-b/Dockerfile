FROM golang:1.15 as builder

ENV GO111MODULE=on

WORKDIR /go/src/microservice-b

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o app .

FROM scratch

WORKDIR /go/bin/microservice-b

COPY --from=builder /go/src/microservice-b .

ENTRYPOINT ["./app"]