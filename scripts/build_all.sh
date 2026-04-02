#!/bin/bash
set -e

VERSION="${VERSION:-1.0.0}"
OUTPUT_DIR="dist"
LDFLAGS="-s -w -X github.com/ricequant/rqdata-cli/cmd.Version=${VERSION}"
GOCACHE_DIR="${GOCACHE_DIR:-$(pwd)/.cache/go-build}"

echo "Building rqdata CLI v${VERSION} for all platforms..."

mkdir -p ${OUTPUT_DIR}
mkdir -p "${GOCACHE_DIR}"

# Linux
echo "Building for Linux..."
GOCACHE="${GOCACHE_DIR}" CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -trimpath -ldflags="${LDFLAGS}" -o ${OUTPUT_DIR}/rqdata-linux-amd64 main.go

# macOS Intel
echo "Building for macOS (Intel)..."
GOCACHE="${GOCACHE_DIR}" CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -trimpath -ldflags="${LDFLAGS}" -o ${OUTPUT_DIR}/rqdata-macos-amd64 main.go

# macOS Apple Silicon
echo "Building for macOS (Apple Silicon)..."
GOCACHE="${GOCACHE_DIR}" CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -trimpath -ldflags="${LDFLAGS}" -o ${OUTPUT_DIR}/rqdata-macos-arm64 main.go

# Windows
echo "Building for Windows..."
GOCACHE="${GOCACHE_DIR}" CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -trimpath -ldflags="${LDFLAGS}" -o ${OUTPUT_DIR}/rqdata-windows-amd64.exe main.go

echo ""
echo "Build complete! Binaries in ${OUTPUT_DIR}/"
ls -lh ${OUTPUT_DIR}/
