FROM docker.io/safewaters/docker-lock:latest@sha256:432d90ddc2891f4845241adc63e5eef2dd1486fa14ea7882433cbd3f8ed64622 AS docker-lock
FROM docker.io/golangci/golangci-lint:latest-alpine@sha256:a392d4e44049a444a927878792dae9534075ec57880e0657647ce818bd8278c2 AS golangci-lint
FROM docker.io/mikefarah/yq:4@sha256:d4e1d060e248bdc3ac9e0fccfb4209e24ff42ed3d39beab33d3c3a19480b14fd AS yq
FROM docker.io/hadolint/hadolint:latest@sha256:fff226bdf9ebcc08db47fb90ee144dd770120b35c2b1cbbb46e932a650cfe232 AS hadolint
FROM --platform=linux/amd64 docker.io/goodwithtech/dockle:v0.4.10@sha256:26524f9f4d1750be7316870f2e74066305b1747a16b1584e522a136b2d507af1 AS dockle
FROM docker.io/moby/buildkit:rootless@sha256:0dc42f831f18275de5e3d667933f2618e8a0e435c538411347a12be8266ab11f AS buildkit
FROM gcr.io/go-containerregistry/crane:latest@sha256:9c275cc44c2811414a9686673f94f7ab220d65593e796860fd31e1f3d8cc0807 AS crane
FROM docker.io/bufbuild/buf:latest@sha256:7898c72aff6c60f1016a5b854e2f84f7ddb5b97e424eb8439b0a535c27c09b22 AS buf
FROM docker.io/library/golang:1.19.4-bullseye@sha256:9d235ae41bff699b44255bf912e77acac5046b06e61b018561239a35d9bd2ebb AS base

SHELL ["/bin/bash", "-o", "pipefail", "-c"]

RUN \
    go env > /usr/local/bin/goenv.sh && \
    true

# https://packages.debian.org/search?keywords=busybox&searchon=names&suite=buster&section=all
RUN \
    apt-get update && \
    apt-get install -y \
        --no-install-recommends \
        make \
        socat \
        curl \
        ca-certificates \
        busybox \
        rsync \
        protobuf-compiler \
        git \
        jq \
        && \
    rm -vr /var/lib/apt/lists/* && \
    true

COPY go.mod go.sum tools.go /var/tmp/build/

RUN \
    export GOBIN=/usr/local/bin && \
    cd /var/tmp/build/ && \
    awk -F '"' '/_.*toolchain/{print $2}' tools.go | xargs -t -I{} go install {} && \
    rm -rf "$(go env GOPATH)" && \
    rm -rf /var/tmp/build/* && \
    protoc-gen-doc --version && \
    true


VOLUME /srv/workspace
WORKDIR /srv/workspace

COPY --from=golangci-lint /usr/bin/golangci-lint /usr/bin/golangci-lint
COPY --from=docker-lock /prod/docker-lock /usr/local/bin/docker-lock
COPY --from=yq /usr/bin/yq /usr/local/bin/yq
COPY --from=hadolint /bin/hadolint /usr/local/bin/hadolint
COPY --from=dockle /usr/bin/dockle /usr/local/bin/dockle
COPY --from=crane /ko-app/crane /usr/local/bin/
COPY --from=buf /usr/local/bin/buf /usr/local/bin/
COPY --from=buildkit /usr/bin/buildkit* /usr/bin/buildctl* /usr/bin/rootlesskit /usr/local/bin/

CMD ["make", "validate", "build"]
