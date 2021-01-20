#From: https://github.com/chris-crone/containerized-go-dev/blob/main/Makefile
#NOTE: looks like docker build uses buildkit to build images.
# you can disable this by setting DOCKER_BUILDKIT=0
all: image

PLATFORM=local
TARGETOS=linux
TARGETARCH=amd64
IMAGE_NAME=inventory-web-app

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
	--build-arg TARGETOS=${TARGETOS} \
	--build-arg TARGETARCH=${TARGETARCH} \
	--tag inventory-web-app:latest

.PHONY: tagPush
tagPush: image
	@docker tag ${IMAGE_NAME} rmccabe3701/${IMAGE_NAME}:latest && \
		docker push rmccabe3701/${IMAGE_NAME}
