# TLS Proxy

A reverse proxy that handles TLS connections and forwards requests to another
URL.

## Install

```
go install github.com/francescomari/tls-proxy
```

## Generate certificates

[mkcert](https://github.com/FiloSottile/mkcert) can generate valid certificates
signed by a local certificate authority (CA).

Install the local CA:

```
mkcert -install
```

Create a certificate and a private key:

```
mkcert -cert-file crt.pem -key-file key.pem localhost
```

## Usage

Start a proxy on port `8080` that redirects every requesto to port `8081` on
`localhost`. The proxy runs with TLS, using the specified certificate and private
key:

```
tls-proxy -source :8080 -target http://localhost:8081 -cert crt.pem -key key.pem
```
