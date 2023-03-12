FROM golang:1.14.6-alpine3.12 as builder
COPY go.mod go.sum /go/src/github.com/rizperdana/prototype-golang-rest/
WORKDIR /go/src/github.com/rizperdana/prototype-golang-rest
RUN go mod download
COPY . /go/src/github.com/rizperdana/prototype-golang-rest
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/prototype-golang-rest github.com/rizperdana/prototype-golang-rest

FROM alpine
RUN apk add --no-cache ca-certificates && update-ca-certificates
COPY --from=builder /go/src/github.com/rizperdana/prototype-golang-rest /usr/bin/prototype-golang-rest
RUN chmod +x /usr/bin/prototype-golang-rest
EXPOSE 8080 8080
ENTRYPOINT [ "/usr/bin/prototype-golang-rest" ]
