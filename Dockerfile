FROM golang:1.16-alpine

ENV GO111MODULE=on
ENV APP_HOME /app

RUN mkdir -p "$APP_HOME"

WORKDIR "$APP_HOME"

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /url-shortener

EXPOSE 8080

CMD [ "/url-shortener" ]