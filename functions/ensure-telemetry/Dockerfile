ARG BUILDER_IMAGE=golang:1.19.1-alpine3.16
ARG BASE_IMAGE=alpine:3.16

FROM --platform=$BUILDPLATFORM $BUILDER_IMAGE AS build
ENV CGO_ENABLED=0
WORKDIR /go/src/

RUN mkdir common && mkdir ensure-telemetry

COPY common/go.mod common/
COPY ensure-telemetry/go.mod ensure-telemetry/go.sum ensure-telemetry/
RUN cd ensure-telemetry && go mod download && cd ..

COPY . .
ARG TARGETOS TARGETARCH
RUN cd ensure-telemetry/ && GOOS=$TARGETOS GOARCH=$TARGETARCH go build -o /usr/local/bin/function

#############################################

FROM $BASE_IMAGE
COPY --from=build /usr/local/bin/function /usr/local/bin/function
ENTRYPOINT ["function"]
