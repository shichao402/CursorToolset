# CursorToolset 文档目录

本目录包含 CursorToolset 的所有文档，按用途分为三类。

## 目录结构

```
docs/
├── README.md                 # 本文件
├── public/                   # 公开文档（随二进制分发）
│   ├── package-dev-guide.md  # 包开发指南
│   ├── release-workflow-template.yml
│   └── examples/             # 配置示例
├── internal/                 # 内部开发文档
│   ├── DEVELOPMENT.md        # 开发指南
│   ├── TESTING.md            # 测试指南
│   ├── BUILD.md              # 构建安装指南
│   ├── RELEASE.md            # 发布指南
│   ├── ARCHITECTURE.md       # 核心设计文档
│   └── AI_COMMUNITY_PROTOCOL.md  # AI 社区协议设计
└── temp/                     # 临时文档（不纳入版本控制）
    └── .gitkeep
```

## 文档分类说明

### 1. 公开文档 (`public/`)

通过 `cursortoolset init` 命令可以获取的文档，面向**包开发者**。

| 文件 | 说明 |
|------|------|
| `package-dev-guide.md` | 包开发完整指南 |
| `release-workflow-template.yml` | GitHub Actions 发布模板 |
| `examples/` | 配置示例（package.json 等） |

### 2. 内部文档 (`internal/`)

面向 **CursorToolset 项目开发者**的文档。

| 文件 | 说明 |
|------|------|
| `DEVELOPMENT.md` | 日常开发流程、代码规范 |
| `TESTING.md` | 测试指南、测试脚本使用 |
| `BUILD.md` | 构建、安装、环境配置 |
| `RELEASE.md` | 版本发布流程 |
| `ARCHITECTURE.md` | 系统架构、核心设计 |
| `AI_COMMUNITY_PROTOCOL.md` | AI 社区协议设计 |

### 3. 临时文档 (`temp/`)

临时性文档，如需求讨论、调试记录等。此目录已加入 `.gitignore`。

## 快速链接

- **我要开发一个包** → [public/package-dev-guide.md](public/package-dev-guide.md)
- **我要参与开发 CursorToolset** → [internal/DEVELOPMENT.md](internal/DEVELOPMENT.md)
- **我要运行测试** → [internal/TESTING.md](internal/TESTING.md)
- **我要发布新版本** → [internal/RELEASE.md](internal/RELEASE.md)
