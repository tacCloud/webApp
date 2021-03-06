#FROM --platform=${BUILDPLATFORM} golang:1.15.2-alpine AS base
FROM golang:1.15.2-alpine AS base
WORKDIR /src
ENV CGO_ENABLED=0
COPY . .

FROM base AS build
ARG TARGETOS
ARG TARGETARCH
RUN GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o /out/webApp ./cmd
RUN echo "I'm building for $TARGETPLATFORM"
#FROM scratch as bin-linux
FROM scratch


COPY --from=build /out/webApp /webApp
COPY cmd/template.html /template.html
#FROM bin-linux as bin-darwin

#FROM bin-${TARGETOS} AS bin

ENTRYPOINT [ "/webApp" ]
