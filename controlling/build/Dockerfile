# syntax=docker/dockerfile:1

ARG GO_VERSION=1.22
ARG ALPINE_VERSION=3.19

#-------------------------------------------------------------------------------
# STAGE: BASE
#-------------------------------------------------------------------------------
FROM golang:${GO_VERSION}-alpine AS base
WORKDIR /app/src
RUN apk add --no-cache tzdata curl
RUN --mount=type=cache,target=${GOMODCACHE} \
    --mount=type=bind,source=go.mod,target=go.mod \
    --mount=type=bind,source=go.sum,target=go.sum \
    --mount=type=bind,from=common_source,source=go.mod,target=/app/common/go.mod \
    --mount=type=bind,from=common_source,source=go.sum,target=/app/common/go.sum \
    go mod download

#-------------------------------------------------------------------------------
# STAGE: TEST
#-------------------------------------------------------------------------------
FROM base AS test
RUN --mount=type=cache,target=${GOMODCACHE} \
    --mount=type=bind,target=. \
    go test -v -coverprofile=/tmp/coverage.txt ./.. > /tmp/result.txt; \
    [[ $? -eq 0 ]] || { cat /tmp/result.txt; exit 1; }

FROM scratch AS test-export
COPY --from=test /tmp/coverage.txt /
COPY --from=test /tmp/result.txt /

#-------------------------------------------------------------------------------
# STAGE: BUILD
#-------------------------------------------------------------------------------
FROM base AS build
ARG VERSION="v0.0.0+unknown"
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
RUN --mount=type=cache,target=${GOMODCACHE} \
    --mount=type=cache,target=${GOCACHE} \
    --mount=type=bind,target=. \
    --mount=type=bind,from=common_source,target=/app/common \
    --mount=type=bind,from=common_source,target=/app/common \
    go build -ldflags="-X './cmd/api/main.Version=${VERSION}'" -o /bin/app ./cmd/api

#-------------------------------------------------------------------------------
# STAGE: DEVELOPMENT
#-------------------------------------------------------------------------------
FROM base AS development
ARG AIR_VERSION=1.52.0
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b ${GOPATH}/bin v${AIR_VERSION}
CMD ["sh", "-c", "air -c .air.toml"]

#-------------------------------------------------------------------------------
# STAGE: PRODUCTION
#-------------------------------------------------------------------------------
FROM alpine:${ALPINE_VERSION} AS production
COPY --from=build /bin/app /bin/app
ENTRYPOINT ["/bin/app"]
