FROM golang:1.7.3-alpine3.11 as builder

ENV GOGO111MODULE=on \
    GOPROXY="https://goproxy.cn,direct" \

COPY go.mod .
COPY go.sum .

RUN go mod download

WORKDIR /build

COPY . .

RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -a -o webhook .


FROM alpine:3.10 AS final

COPY --from=builder /build/webhook  /app/

EXPOSE 8989
CMD ["/app/webhook"]

