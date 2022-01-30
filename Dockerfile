FROM golang:1.17.3-alpine3.13 as builder

WORKDIR /build

ENV GOPROXY="https://goproxy.cn,direct"

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -a -o webhook .


FROM alpine:3.13 AS final

COPY --from=builder /build/webhook /webhook

EXPOSE 8989
CMD ["/webhook"]

