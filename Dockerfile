FROM golang:alpine3.14 as builder
RUN apk add curl
RUN sh -c "$(curl --location https://taskfile.dev/install.sh)" -- -d -b /bin
RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN task build -v

FROM gcr.io/distroless/static
WORKDIR /root/
COPY --from=builder  /build/cmd/tripletstat/tripletstat .
COPY --from=builder  /build/cmd/tripletserver/tripletserver .
COPY --from=builder  /build/cmd/tripletclient/tripletclient .

ENTRYPOINT ["/root/tripletstat"]