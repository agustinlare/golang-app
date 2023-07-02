FROM golang:1.20-bullseye

WORKDIR $GOPATH/src/go-dummy/app/

COPY . .

RUN go mod tidy
RUN go build

EXPOSE 8082

CMD ["./go-dummy"]