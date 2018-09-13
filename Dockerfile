FROM scalify/glide:0.13.0 as builder
WORKDIR /go/src/github.com/Scalify/puppet-master-cli/

COPY glide.yaml glide.lock ./
RUN glide install --strip-vendor

COPY . ./
RUN CGO_ENABLED=0 go build -a -ldflags '-s' -installsuffix cgo -o bin/puppet-master-cli .


FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/Scalify/puppet-master-cli/bin/puppet-master-cli .
RUN chmod +x cli
ENTRYPOINT ["./puppet-master-cli"]
