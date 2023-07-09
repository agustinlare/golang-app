FROM golang:1.20-bullseye

WORKDIR $GOPATH/src/go-dummy/app/

ARG EXPOSED_PORT
ENV EXPOSED_PORT=${EXPOSED_PORT}

COPY . .

RUN go mod tidy
RUN go build

EXPOSE $EXPOSED_PORT

CMD ["./golang-app"]