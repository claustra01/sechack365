dev-cert:
	sudo apt install mkcert -y
	mkcert -install
	mkcert -cert-file ./nginx/default.crt -key-file ./nginx/default.key localhost
