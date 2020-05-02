FROM golang as builder
WORKDIR /src/

COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN CGO_ENABLED=0 go build -a -ldflags '-s' -installsuffix cgo -o bin/puppet-master-cli .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /src/bin/puppet-master-cli .
RUN chmod +x puppet-master-cli
ENTRYPOINT ["./puppet-master-cli"]
