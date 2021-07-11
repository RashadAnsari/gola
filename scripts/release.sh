#!/usr/bin/env bash

APP=gola

APP_VERSION=$1

BUILD_INFO_PKG="github.com/RashadAnsari/gola/pkg/version"
LDFLAGS="-w -s -X '${BUILD_INFO_PKG}.AppVersion=${APP_VERSION}' -X '${BUILD_INFO_PKG}.Date=$(date)' -X '${BUILD_INFO_PKG}.BuildVersion=$(git rev-parse HEAD | cut -c 1-8)' -X '${BUILD_INFO_PKG}.VCSRef=$(git rev-parse --abbrev-ref HEAD)'"

RELEASE_FOLDER=releases

function release {
    GO_OS=$1
    GO_ARCH=$2

    GOOS=${GO_OS} GOARCH=${GO_ARCH} go build -ldflags "${LDFLAGS}" ./cmd/${APP}

    BINARY_NAME=${APP}-${APP_VERSION}-${GO_OS}-${GO_ARCH}
    mv ${APP} ${BINARY_NAME}

    mv ${BINARY_NAME} ${RELEASE_FOLDER}
}

mkdir -p ${RELEASE_FOLDER}

release darwin amd64
release linux amd64
