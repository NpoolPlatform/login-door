#/bin/bash
URL=https://github.com/go-swagger/go-swagger/releases/download/v0.28.0/swagger_linux_amd64

install: pre tool

pre:
ifndef ($(wildcard /user/local/bin/swagger)
    curl -sL -o /usr/local/bin/swagger ${URL}
endif

tool:
	chmod +x /usr/local/bin/swagger