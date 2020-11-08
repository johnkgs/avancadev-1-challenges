FROM golang:1.15 as builder

WORKDIR /go/src/microservice-c

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o app .

FROM scratch

WORKDIR /go/bin/microservice-c

COPY --from=builder /go/src/microservice-c .

ENTRYPOINT ["./app"]

EXPOSE 9092