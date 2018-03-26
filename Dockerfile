FROM golang:1.10.0-alpine
LABEL maintainer="ymdarake <https://github.com/ymdarake>"

RUN apk update && apk add curl && apk add git
RUN curl -fsSL -o /usr/local/bin/dep https://github.com/golang/dep/releases/download/v0.4.1/dep-linux-amd64 && chmod +x /usr/local/bin/dep

ADD . $GOPATH/src/chat-app
WORKDIR $GOPATH/src/chat-app
RUN dep ensure
WORKDIR $GOPATH/src/chat-app
RUN go build -o chat

ENV API_CLIENT_ID_GOOGLE="FIXME"
ENV API_CLIENT_SECRET_GOOGLE="FIXME"
# use default port specified in main.go
EXPOSE 8080

CMD ./chat
