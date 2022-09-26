ARG BUILDER_IMAGE=golang:1.19.1
ARG BASE_IMAGE=quay.io/argoproj/argocd:v2.4.12

FROM --platform=$BUILDPLATFORM $BUILDER_IMAGE AS build
ENV CGO_ENABLED=0
WORKDIR /go/src/

RUN mkdir kpt-render && mkdir kpt-repo

COPY kpt-repo/go.mod kpt-repo/go.sum kpt-repo/
COPY kpt-render/go.mod kpt-render/go.sum kpt-render/
RUN cd kpt-repo && go mod download && cd ..
RUN cd kpt-render && go mod download && cd ..

COPY . .
ARG TARGETOS TARGETARCH
RUN cd kpt-repo/ && GOOS=$TARGETOS GOARCH=$TARGETARCH go build -o /usr/local/bin/kpt-repo
RUN cd kpt-render/ && GOOS=$TARGETOS GOARCH=$TARGETARCH go build -o /usr/local/bin/kpt-render

#############################################

FROM $BASE_IMAGE
COPY --from=build /usr/local/bin/kpt-repo /usr/local/bin/kpt-repo
COPY --from=build /usr/local/bin/kpt-render /usr/local/bin/kpt-render
