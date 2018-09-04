FROM golang:1.10.3-alpine3.8 as builder

RUN apk update \
    && apk add ca-certificates \
    && rm -rf /var/cache/apk/* \
    && update-ca-certificates \
    && apk add git

WORKDIR /go/src/github.com/jcorioland/keyvault-aad-pod-identity
COPY . .

RUN go get -d ./...
RUN go build -o /keyvault-sample main.go

FROM alpine:3.8
RUN apk update \
    && apk add ca-certificates \
    && rm -rf /var/cache/apk/* \
    && update-ca-certificates

EXPOSE 8080
ENTRYPOINT [ "/keyvault-sample" ]
COPY --from=builder /keyvault-sample /