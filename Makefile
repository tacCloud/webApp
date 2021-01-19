#From: https://github.com/chris-crone/containerized-go-dev/blob/main/Makefile
all: bin/webApp

PLATFORM=local

.PHONY: bin/webApp
bin/webApp:
	@docker build . \
	--output bin/ \
	--target bin \
	--tag inventory-web-app:latest \
	--platform ${PLATFORM}

.PHONY: image
image:
	@docker build . \
	--tag inventory-web-app:latest \
	--target bin-darwin \
	--platform ${PLATFORM}