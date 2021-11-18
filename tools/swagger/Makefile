#/bin/bash
URL=https://github.com/go-swagger/go-swagger/releases/download/v0.28.0/swagger_linux_amd64

install: tool

pre:
ifndef JQ
	$(error "jq not exist, please run 'apt install jq'")
endif

tool:
	curl -sL -o /usr/local/bin/swagger ${URL}
	chmod +x /usr/local/bin/swagger