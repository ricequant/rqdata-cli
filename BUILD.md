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
# Linux x64
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -trimpath -ldflags="-s -w -X github.com/ricequant/rqdata-cli/cmd.Version=1.0.0" -o rqdata-linux-amd64 main.go

# Linux arm64
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -trimpath -ldflags="-s -w -X github.com/ricequant/rqdata-cli/cmd.Version=1.0.0" -o rqdata-linux-arm64 main.go

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

该命令会构建 Linux x64、Linux arm64、macOS 和 Windows 预编译二进制，并生成 5 个平台包和 1 个主包的 npm tarball。

正式发布使用：

```bash
npm run publish:all
```

## Python wheel 构建

Python wheel 采用“单一项目名 + 多平台 wheel”模式，每个 wheel 内置对应平台的 Go 二进制。wheel 平台标签由 `setup.py` 显式指定、与宿主环境无关，因此**可以在任意平台交叉构建**（不需要 Linux 机器，也不依赖 CGO）。

### 前置要求

- Go 1.21+：用于编译 wheel 内置的二进制
- Python 3.8+
- 构建前端，二选一：`python3 -m pip install build`（推荐）或 `python3 -m pip install wheel`

### 快速开始

```bash
# 构建宿主平台 wheel，产物写入 wheelhouse/
python3 scripts/python/build-wheels.py

# 构建指定平台 wheel，例如 Linux arm64
python3 scripts/python/build-wheels.py --target linux-arm64

# 构建全部 5 个平台 wheel
python3 scripts/python/build-wheels.py --target all
```

### 平台对照表

| `--target` | Go 目标 | 内置二进制 | wheel 标签 |
| --- | --- | --- | --- |
| `linux-x64` | `GOOS=linux GOARCH=amd64` | `rqdata` | `manylinux2014_x86_64` |
| `linux-arm64` | `GOOS=linux GOARCH=arm64` | `rqdata` | `manylinux2014_aarch64` |
| `darwin-x64` | `GOOS=darwin GOARCH=amd64` | `rqdata` | `macosx_10_15_x86_64` |
| `darwin-arm64` | `GOOS=darwin GOARCH=arm64` | `rqdata` | `macosx_11_0_arm64` |
| `win32-x64` | `GOOS=windows GOARCH=amd64` | `rqdata.exe` | `win_amd64` |

### 参数说明

| 参数 | 说明 |
| --- | --- |
| `--target <key\|all>` | 目标平台，默认自动识别宿主平台，`all` 表示全部平台 |
| `--out-dir <dir>` | 输出目录，默认 `wheelhouse/` |
| `--clean` | 构建前删除 `--out-dir` 整个目录（含目录中已有 wheel）；只构建单平台时慎用，会连带删掉其他平台的产物 |
| `--isolated` | 使用隔离构建环境，需要联网安装构建依赖；默认关闭 |

### 构建流程

1. `scripts/python/build-wheels.py` 设置 `RQDATA_PY_TARGET`，调用 `python3 -m build --wheel --no-isolation`；环境里没有 `build` 时回退到 `python3 setup.py bdist_wheel`。
2. `setup.py` 的 `build_py` 检查 `dist/<二进制名>`：不存在时按目标平台执行 `CGO_ENABLED=0 GOOS=... GOARCH=... go build -trimpath -ldflags "-s -w -X github.com/ricequant/rqdata-cli/cmd.Version=<version>"`，然后把二进制复制到包内 `rqdata_cli/bin/` 并设置 `0o755`。`<version>` 取自 `package.json`，wheel 版本号同样取自该文件。
3. `setup.py` 的 `bdist_wheel.get_tag()` 强制返回 `py3-none-<平台标签>`，并设置 `Root-Is-Purelib: false`，wheel 标签不依赖宿主 Python/系统。

不经过封装脚本的等价命令：

```bash
RQDATA_PY_TARGET=linux-arm64 python3 -m build --wheel --no-isolation --outdir wheelhouse
```

### 二进制缓存

`dist/` 中已存在目标二进制时直接复用，不会重新编译 Go 代码。修改 Go 源码后需要先删除对应文件再构建：

```bash
rm -f dist/rqdata-linux-arm64
python3 scripts/python/build-wheels.py --target linux-arm64
```

### 验证

检查每个 wheel 的平台标签与内置二进制：

```bash
python3 - <<'PY'
import glob, zipfile
for path in sorted(glob.glob("wheelhouse/*.whl")):
    with zipfile.ZipFile(path) as z:
        meta = next(n for n in z.namelist() if n.endswith(".dist-info/WHEEL"))
        tag = next(l for l in z.read(meta).decode().splitlines() if l.startswith("Tag: "))
        bins = [i for i in z.infolist() if "/bin/" in i.filename]
    print(path)
    print(f"  {tag}")
    for b in bins:
        print(f"  {b.filename} ({b.file_size} bytes)")
PY
```

每个 wheel 应只包含一个内置二进制（`rqdata` 或 Windows 的 `rqdata.exe`），标签应与[平台对照表](#平台对照表)一致。若出现多余条目，说明该 wheel 由旧构建流程产出，需重新构建后再发布。

在目标平台（或对应架构的容器）中安装验证：

```bash
pip install wheelhouse/rqdata_cli-1.0.5-py3-none-manylinux2014_aarch64.whl
rqdata --version
rqdata --help
```

### 发布

```bash
python3 -m twine upload wheelhouse/*.whl
```

发布前检查：

- `package.json` 的 `version` 已更新（wheel 版本号与内置二进制版本都取自该文件）
- `wheelhouse/` 中 5 个平台的 wheel 版本号一致，且都已完成构建；缺任一平台时，该平台的 `pip install rqdata-cli` 会直接失败（仓库不提供 sdist 兜底）

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
