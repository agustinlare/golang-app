FROM golang:1.20-bullseye

WORKDIR $GOPATH/src/go-dummy/app/

COPY . .

RUN go mod tidy
RUN go build

CMD ["./go-dummy"]