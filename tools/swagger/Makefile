#/bin/bash
URL=curl -s https://api.github.com/repos/go-swagger/go-swagger/releases/latest | \
  jq -r '.assets[] | select(.name | contains("'"$(uname | tr '[:upper:]' '[:lower:]')"'_amd64")) | .browser_download_url'

install: tool

pre:
ifndef JQ
	$(error "jq not exist, please run 'apt install jq'")
endif

tool:
	curl -o /usr/local/bin/swagger -L'#' ${URL}
	chmod +x /usr/local/bin/swagger