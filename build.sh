#!/bin/bash
set -e

VERSION="${VERSION:-dev}"
LDFLAGS="-s -w -X github.com/ricequant/rqdata-cli/cmd.Version=${VERSION}"
GOCACHE_DIR="${GOCACHE_DIR:-$(pwd)/.cache/go-build}"

echo "Building rqdata CLI (Go version)..."

mkdir -p "${GOCACHE_DIR}"

# 构建当前平台
GOCACHE="${GOCACHE_DIR}" go build -trimpath -ldflags="${LDFLAGS}" -o rqdata main.go

echo "Build complete: ./rqdata"
echo "Size: $(du -h rqdata | cut -f1)"
echo "Version: ${VERSION}"

# 测试
./rqdata --help
