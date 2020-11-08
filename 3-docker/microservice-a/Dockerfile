FROM golang:1.15 as builder

ENV GO111MODULE=on

WORKDIR /go/src/microservice-a

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o app .

FROM scratch

WORKDIR /go/bin/microservice-a

COPY --from=builder /go/src/microservice-a .

ENTRYPOINT ["./app"]

EXPOSE 9090