# Go HTTPS Demo

[![Build Status](https://travis-ci.org/themccallister/go-https-demo.svg?branch=master)](https://travis-ci.org/themccallister/go-https-demo)

This is a quick demo project on generating a local certificate and spinning up a HTTPS server in Go.

## Generating a Certificate

1. Run `make csr` to generate the Certificate Signing Request (CSR)
2. Run `make cert` and fill in the prompted information to generate the certificate for the server

## Building and Running the Application

1. `make run` will build the project and run the application

> Note: You might be denied permissions from running as the default/fallback port is 443 which will require sudo permissions.

### Environment Variables

The application checks for the following environment variables, if not found, it falls back to a default.

1. `APP_PORT` the port for the web server to listen on
2. `APP_CERT_FILE` the path to the certificate file for HTTPS
3. `APP_KEY_FILE` the path to the key file for HTTPS
