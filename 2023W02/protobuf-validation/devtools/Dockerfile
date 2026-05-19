FROM docker.io/safewaters/docker-lock:latest@sha256:432d90ddc2891f4845241adc63e5eef2dd1486fa14ea7882433cbd3f8ed64622 AS docker-lock
FROM docker.io/golangci/golangci-lint:latest-alpine@sha256:91b27804074a0bacea298707f016911e60cf0cdbc6c7bf5ccacb5f0606d18d60 AS golangci-lint
FROM docker.io/mikefarah/yq:4@sha256:0cb4a78491b6e62ee8a9bf4fbeacbd15b5013d19bc420591b05383a696315e60 AS yq
FROM docker.io/hadolint/hadolint:latest@sha256:27086352fd5e1907ea2b934eb1023f217c5ae087992eb59fde121dce9c9ff21e AS hadolint
FROM --platform=linux/amd64 docker.io/goodwithtech/dockle:v0.4.15@sha256:eade932f793742de0aa8755406c7677cd7696f8675b6180926f7eeffa7abe6b9 AS dockle
FROM docker.io/moby/buildkit:rootless@sha256:d76eb1caecac5733ef7553c1e90a1b21f1bb218cd1142d3553de0747b4a14ba9 AS buildkit
FROM gcr.io/go-containerregistry/crane:latest@sha256:0ae17ecb34315aa7cbff28f6eddee3b7adae0b2f90101260d990804db1eb0084 AS crane
FROM docker.io/bufbuild/buf:latest@sha256:daa3bf554150247bc0d47f73fb8b355faa193dad5428aa1efcf74480e210ec40 AS buf
FROM docker.io/library/golang:1.24.6-bullseye@sha256:2cdc80dc25edcb96ada1654f73092f2928045d037581fa4aa7c40d18af7dd85a AS base

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
