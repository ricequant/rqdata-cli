# 贡献指南

感谢你对 RQData CLI 的关注。

## 问题与功能请求

- Bug 报告请使用 GitHub Issues
- 新功能建议也请先在 GitHub Issues 讨论
- 提交问题时请包含复现步骤、预期行为、实际行为、运行环境和关键日志

## 本地开发

```bash
git clone https://github.com/ricequant/rqdata-cli.git
cd rqdata-cli
```

构建：

```bash
VERSION=dev ./build.sh
```

验证：

```bash
./rqdata --help
./rqdata --version
```

如果需要验证 npm 分发链路：

```bash
npm run build:npm
npm run link:platform
node bin/rqdata.js --version
```

如果需要验证 Python wheel 分发链路：

```bash
python3 -m pip install build
python3 scripts/python/build-wheels.py --target linux-arm64
```

`--target` 用于指定平台（`linux-x64`、`linux-arm64`、`darwin-x64`、`darwin-arm64`、`win32-x64`），省略时构建宿主平台。wheel 标签由 `setup.py` 显式指定，可在任意平台交叉构建；完整的参数说明、验证与发布流程见 [BUILD_GO.md](../BUILD_GO.md#python-wheel-构建)。

## 提交流程

1. Fork 仓库并创建分支
2. 完成修改并更新相关文档
3. 确认构建通过
4. 提交清晰的 commit message
5. 发起 Pull Request

推荐的提交前缀：

- `feat:` 新功能
- `fix:` Bug 修复
- `docs:` 文档更新
- `refactor:` 重构
- `test:` 测试相关
- `chore:` 构建或工具调整

## 代码约定

- 使用 Go 1.21+
- 保持命令行行为向后兼容
- 新增功能时同步更新 README 或相关文档
- 尽量保持输出格式稳定，避免破坏脚本或 Agent 集成

## 发布

npm 发布采用“主包 + 平台包”模式：

- `@ricequant2026/rqdata-cli`
- `@ricequant2026/rqdata-cli-linux-x64`
- `@ricequant2026/rqdata-cli-linux-arm64`
- `@ricequant2026/rqdata-cli-darwin-x64`
- `@ricequant2026/rqdata-cli-darwin-arm64`
- `@ricequant2026/rqdata-cli-win32-x64`

本地打包：

```bash
npm run pack:all
```

发布：

```bash
npm run publish:all
```

预发布测试：

```bash
npm version prerelease --preid rc
npm run publish:next
```

Python wheel 发布采用“单一项目名 + 多平台 wheel”模式：

- `rqdata-cli`

本地打包：

```bash
python3 scripts/python/build-wheels.py --target all --clean
```

`--clean` 会先删除整个 `wheelhouse/` 目录，因此 `--target all` 时应始终配合使用；只构建单个平台时省略该参数，避免删掉其他平台的产物。

发布：

```bash
python3 -m twine upload wheelhouse/*.whl
```

发布前确认 `wheelhouse/` 中 5 个平台的 wheel 版本号一致且均已构建，版本号取自 `package.json`。

## 获取帮助

- Issues: https://github.com/ricequant/rqdata-cli/issues
- Repository: https://github.com/ricequant/rqdata-cli
