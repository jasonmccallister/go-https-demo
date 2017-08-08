cert:
	openssl x509 -req -sha256 -days 365 -in certs/local.demodomain.com.csr -signkey certs/local.demodomain.com.key -out certs/local.demodomain.com.crt

csr:
	openssl genrsa -des3 -passout pass:x -out certs/local.demodomain.com.pass.key 2048 && \
	openssl rsa -passin pass:x -in certs/local.demodomain.com.pass.key -out certs/local.demodomain.com.key && \
	rm certs/local.demodomain.com.pass.key && \
	openssl req -new -key certs/local.demodomain.com.key -out certs/local.demodomain.com.csr