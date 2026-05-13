# Build the manager binary
FROM registry.access.redhat.com/ubi9/go-toolset:1.25 AS builder
ENV PATH="$PATH:/opt/app-root/src/go/bin"

WORKDIR /go/src/github.com/hrathina/odh-trainer-operator
COPY go.mod  go.mod
COPY go.sum  go.sum
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download

COPY cmd/trainer-module/ cmd/trainer-module/
COPY pkg/              pkg/
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 GOOS=linux GOFLAGS=-mod=readonly go build -a -o manager ./cmd/trainer-module

# Collect trainer manifests (placeholder for now, will be implemented in manifest collection ticket)
COPY hack/  hack/
# TODO: Uncomment when manifest collection script is ready
# RUN bash hack/get_trainer_manifests.sh

# Runtime
FROM registry.access.redhat.com/ubi9/ubi-minimal:latest
RUN microdnf install -y --disablerepo=* --enablerepo=ubi-9-baseos-rpms shadow-utils && \
    microdnf clean all && \
    useradd trainer -m -u 1000 && \
    microdnf remove -y shadow-utils

COPY --from=builder /go/src/github.com/hrathina/odh-trainer-operator/manager /manager

# TODO: Copy manifests when manifest collection is implemented
# COPY --from=builder /go/src/github.com/hrathina/odh-trainer-operator/opt/manifests/ /opt/manifests/

# For now, create empty manifests directory to satisfy main.go validation
RUN mkdir -p /opt/manifests

USER 1000:1000
ENTRYPOINT ["/manager"]
