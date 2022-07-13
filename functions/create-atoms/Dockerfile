ARG BUILDER_IMAGE=golang:1.17-alpine3.15
ARG BASE_IMAGE=alpine:3.15


FROM --platform=$BUILDPLATFORM $BUILDER_IMAGE AS build
ENV CGO_ENABLED=0
WORKDIR /go/src/

COPY go.mod go.sum ./
RUN go mod download

COPY . .
ARG TARGETOS TARGETARCH
RUN GOOS=$TARGETOS GOARCH=$TARGETARCH go build -o /usr/local/bin/function ./

#############################################

FROM $BASE_IMAGE
COPY --from=build /usr/local/bin/function /usr/local/bin/function
COPY elements.yaml /etc/elements.yaml
ENTRYPOINT ["function"]
