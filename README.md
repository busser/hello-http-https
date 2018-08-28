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

### Server configuration

The server is configurable through CLI flags. For details on usage, run the
following command:

```bash
hello-http-https -help
```

You should see something like this:

```
Usage of ./hello-http-https:
  -help
    	Prints this usage message
  -http-port int
    	Port used to listen for HTTP requests (default 80)
  -https-port int
    	Port used to listen for HTTPS requests (default 443)
  -ssl-certificate string
    	SSL certificate file for HTTPS server (default "server.crt")
  -ssl-key string
    	SSL key file for HTTPS server (default "server.key")
  -uri-path string
    	URI path to respond to (default "/hello-world")
```

The server's default behavior is to listen for HTTP requests on port 80 and for
HTTPS requests on port 443. It will respond to requests to `/hello-world` with a
short message, and a 404 error code for other paths.

If you wish for the server to answer to all paths, use the `-uri-path` flag with
the `/` value:

```bash
hello-http-https -uri-path /
```
