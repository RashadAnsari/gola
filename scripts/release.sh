#!/usr/bin/env bash

APP=$1
APP_VERSION=$2
LDFLAGS=$3

RELEASE_FOLDER=releases

function release {
    GO_OS=$1
    GO_ARCH=$2

    GOOS=${GO_OS} GOARCH=${GO_ARCH} go build -ldflags "${LDFLAGS}"  ./cmd/${APP}

    BINARY_NAME=${APP}-${APP_VERSION}-${GO_OS}-${GO_ARCH}
    mv ${APP} ${BINARY_NAME}

    mv ${BINARY_NAME} ${RELEASE_FOLDER}
}

mkdir -p ${RELEASE_FOLDER}

release darwin amd64
release linux amd64
