# Makefile variables
DOMAIN=local.demodomain.com

# CSR variables
CSR_COUNTRY=US
CSR_STATE=Virginia
CSR_CITY=Norfolk
CSR_ORG=themccallister
CSR_ORG_UNIT=Engineering
CSR_COMMON_NAME=${DOMAIN}
CSR_EMAIL=themccallister@gmail.com

build:
	go build .

cert:
	openssl x509 -req -sha256 -days 365 -in certs/${DOMAIN}.csr \
	-signkey certs/${DOMAIN}.key -out certs/${DOMAIN}.crt

csr:
	openssl genrsa -des3 -passout pass:x -out certs/${DOMAIN}.pass.key 2048 && \
	openssl rsa -passin pass:x -in certs/${DOMAIN}.pass.key -out certs/${DOMAIN}.key && \
	rm certs/${DOMAIN}.pass.key && \
	openssl req -new -key certs/${DOMAIN}.key -out certs/${DOMAIN}.csr

run: build
	./go-https-demo
