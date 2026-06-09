# Echo Practice — 公司资产管理后端 (study_go_web)

[简体中文](#简体中文) | [English](#english)

---

<a name="简体中文"></a>
## 简体中文

> **免责声明：** 本项目仅用于 **学习与练习**。以 Go Web 开发入门为目标，从基础 CRUD 逐步扩展到完整的业务系统。

### 🚀 项目概览

这是一个基于 **Go + Echo + GORM** 构建的公司资产管理后端学习项目。从最简单的 User CRUD 起步，以"公司资产管理"为业务场景，逐步迭代学习 Web 开发的完整技术栈。

#### 包含的功能模块

| 模块 | 状态 | 说明 |
|------|------|------|
| 用户管理 | ✅ 完成 | 用户的增删改查 |
| 部门管理 | ✅ 完成 | 树形结构，支持多级部门 |
| 员工管理 | 🚧 待开发 | 所属部门关联 |
| 资产分类 | 🚧 待开发 | 树形分类 |
| 资产管理 | 🚧 待开发 | 资产编号、状态流转、分配管理 |
| 维修记录 | 🚧 待开发 | 资产生命周期跟踪 |

#### 架构设计

严格遵循 **Handler → Service → Repository** 三层架构：

```
请求 → handler（参数校验、响应） → service（业务逻辑） → repository（数据访问） → GORM → MySQL
```

- **关注点分离**：每层只做自己的事
- **依赖注入**：通过构造函数注入，便于测试
- **Context 贯穿**：`context.Context` 逐层传递，支持超时取消

#### 技术栈

- **语言：** Go 1.25+
- **Web 框架：** [Echo v4](https://echo.labstack.com/)
- **API 规范：** [Huma v2](https://huma.rocks/)（自动生成 OpenAPI 文档）
- **ORM：** [GORM](https://gorm.io/)
- **数据库：** MySQL
- **配置管理：** [Viper](https://github.com/spf13/viper)
- **热重载：** [Air](https://github.com/air-verse/air)

### 🏗 项目结构

```
.
├── cmd/server/              # 应用程序入口
├── configs/
│   ├── config.yaml          # 默认配置
│   ├── config.dev.yaml      # 开发环境
│   └── config.prod.yaml     # 生产环境
├── internal/
│   ├── config/              # 配置加载（支持多环境切换）
│   ├── handler/             # HTTP 处理器 + 路由注册
│   ├── model/               # 数据模型（6 张表）
│   ├── repository/          # 数据访问层（接口 + GORM 实现）
│   └── service/             # 业务逻辑层
├── .air.toml                # Air 热重载配置
└── go.mod
```

### 🚦 快速上手

#### 前置要求
- Go 1.25+
- MySQL 实例
- [Air](https://github.com/air-verse/air)（可选，推荐）

#### 运行

```bash
# 1. 克隆
git clone git@github.com:kingoflongevity/study_go_web.git
cd study_go_web

# 2. 配置数据库连接
#    编辑 configs/config.dev.yaml 中的 database.source

# 3. 启动（热重载）
air

# 4. 查看 API 文档
#    浏览器打开 http://localhost:8080/docs
```

### 📡 API 接口

#### 用户管理
- `GET    /api/v1/users/{id}` — 获取用户
- `POST   /api/v1/users`     — 创建用户
- `PUT    /api/v1/users/{id}` — 更新用户
- `DELETE /api/v1/users/{id}` — 删除用户

#### 部门管理
- `GET    /api/v1/departments`     — 部门树形列表
- `POST   /api/v1/departments`     — 创建部门
- `GET    /api/v1/departments/{id}` — 获取部门详情
- `PUT    /api/v1/departments/{id}` — 更新部门
- `DELETE /api/v1/departments/{id}` — 删除部门

---

<a name="english"></a>
## English

### 🚀 Project Overview

A Go web backend learning project built with **Echo + GORM + Huma**, simulating a company asset management system. Started from basic User CRUD and progressively expanding into a full-featured business system.

#### Architecture

**Handler → Service → Repository** three-layer architecture with:
- Separation of concerns
- Constructor-based dependency injection
- `context.Context` propagation through all layers

#### Tech Stack
- **Language:** Go 1.25+
- **Web Framework:** [Echo v4](https://echo.labstack.com/)
- **API Spec:** [Huma v2](https://huma.rocks/) (auto-generated OpenAPI docs)
- **ORM:** [GORM](https://gorm.io/)
- **Database:** MySQL
- **Config:** [Viper](https://github.com/spf13/viper)
- **Hot Reload:** [Air](https://github.com/air-verse/air)

#### Quick Start

```bash
git clone git@github.com:kingoflongevity/study_go_web.git
cd study_go_web
# Edit configs/config.dev.yaml database.source
air
# Open http://localhost:8080/docs
```
