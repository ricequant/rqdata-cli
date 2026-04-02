# Go 版本构建指南

## 前置要求
- Go 1.21+

## 安装 Go
```bash
# Linux
wget https://go.dev/dl/go1.21.0.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.21.0.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin

# macOS
brew install go

# Windows
# 下载并安装 https://go.dev/dl/
```

## 构建

### 单平台构建
```bash
VERSION=1.0.0 ./build.sh
```

### 跨平台构建（静态链接）
```bash
# Linux
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -trimpath -ldflags="-s -w -X github.com/ricequant/rqdata-cli/cmd.Version=1.0.0" -o rqdata-linux main.go

# macOS
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -trimpath -ldflags="-s -w -X github.com/ricequant/rqdata-cli/cmd.Version=1.0.0" -o rqdata-macos main.go

# macOS Apple Silicon
CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -trimpath -ldflags="-s -w -X github.com/ricequant/rqdata-cli/cmd.Version=1.0.0" -o rqdata-macos-arm64 main.go

# Windows
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -trimpath -ldflags="-s -w -X github.com/ricequant/rqdata-cli/cmd.Version=1.0.0" -o rqdata.exe main.go
```

## npm 发布构建

如果通过 npm 分发，发布前执行：

```bash
npm run pack:all
```

该命令会构建 Linux/macOS/Windows 预编译二进制，并生成 4 个平台包和 1 个主包的 npm tarball。

正式发布使用：

```bash
npm run publish:all
```

## 验证单文件
```bash
# Linux/macOS
ldd rqdata  # 应显示 "not a dynamic executable"

# macOS
otool -L rqdata  # 应只显示系统库

# Windows
dumpbin /dependents rqdata.exe  # 应只显示系统 DLL
```

## 优势
- 真正的静态编译，无需任何外部依赖
- 单个可执行文件，跨平台分发
- 编译速度快，二进制体积小（~5-10MB）
- 原生支持并发和网络操作
