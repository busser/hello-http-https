# HTTP/HTTPS Hello World!

## Install the binary

```bash
go get github.com/busser/hello-http-https
```

## Generate SSL key and certificate

> You don't need to do this if you already have a key and certificate. The
> certificate, if signed, should contain the entire chain up to the root CA.

```bash
openssl genrsa \
  -out server.key \
  2048
openssl req \
  -new \
  -x509 \
  -sha256 \
  -key server.key \
  -out server.crt \
  -days 365 \
  -subj "/C=FR/ST=Paris/L=Paris/O=Global Security/OU=IT Department/CN=example.com"
```

## Start the server

```bash
hello-http-https
```
