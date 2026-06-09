# Echo Practice (study_go_web)

[简体中文](#简体中文) | [English](#english)

---

<a name="简体中文"></a>
## 简体中文

> **免责声明：** 本项目仅用于 **学习与练习**。旨在探索 Go 语言、Echo Web 框架以及整洁架构（Clean Architecture）原则。

### 🚀 项目概览
基于 Go 和 Echo 框架构建的简单 RESTful API，展示了基础的 CRUD 操作、依赖注入以及使用 GORM 进行数据库交互。

#### 功能特性
- 用户管理（增、删、查）
- 整洁架构分层（Handler -> Service -> Repository）
- 使用 Viper 进行配置管理
- 使用 `go-playground/validator` 进行请求校验
- 基于 GORM (MySQL) 的数据库 ORM
- 使用 Air 支持热重载

### 🛠 技术栈
- **语言：** Go
- **Web 框架：** [Echo](https://echo.labstack.com/)
- **ORM：** [GORM](https://gorm.io/)
- **数据库：** MySQL
- **配置管理：** [Viper](https://github.com/spf13/viper)
- **校验库：** [Validator](https://github.com/go-playground/validator)

### 🏗 项目结构
```text
.
├── cmd/server/          # 应用程序入口
├── configs/             # 配置文件 (YAML)
├── internal/
│   ├── config/          # 配置加载逻辑
│   ├── handler/         # HTTP 处理函数与路由注册
│   ├── model/           # 数据库模型
│   ├── repository/      # 数据访问层
│   └── service/         # 业务逻辑层
└── pkg/                 # 公共工具库
```

### 🚦 快速上手
#### 前置要求
- Go 1.25+
- MySQL 实例
- [Air](https://github.com/air-verse/air) (可选，用于开发环境热重载)

#### 安装步骤
1. 克隆仓库：
   ```bash
   git clone git@github.com:kingoflongevity/study_go_web.git
   cd study_go_web
   ```
2. 安装依赖：
   ```bash
   go mod download
   ```
3. 在 `configs/config.yaml` 中配置数据库连接。

#### 运行项目
**使用 Air (推荐):** `air`  
**标准运行:** `go run cmd/server/main.go`

---

<a name="english"></a>
## English

> **Disclaimer:** This project is for **learning and practice purposes only**. It is used to explore the Go language, the Echo web framework, and clean architecture principles.

### 🚀 Project Overview
A simple RESTful API built with Go and the Echo framework, demonstrating basic CRUD operations, dependency injection, and database interaction using GORM.

#### Features
- User Management (Create, Read, Delete)
- Clean Architecture (Handler -> Service -> Repository)
- Configuration management using Viper
- Request validation using `go-playground/validator`
- Database ORM with GORM (MySQL)
- Hot-reloading with Air

### 🛠 Tech Stack
- **Language:** Go
- **Web Framework:** [Echo](https://echo.labstack.com/)
- **ORM:** [GORM](https://gorm.io/)
- **Database:** MySQL
- **Config Management:** [Viper](https://github.com/spf13/viper)
- **Validation:** [Validator](https://github.com/go-playground/validator)

### 🏗 Project Structure
```text
.
├── cmd/server/          # Application entry point
├── configs/             # Configuration files (YAML)
├── internal/
│   ├── config/          # Config loading logic
│   ├── handler/         # HTTP handlers & routing
│   ├── model/           # Database models
│   ├── repository/      # Data access layer
│   └── service/         # Business logic layer
└── pkg/                 # Shared utilities
```

### 🚦 Getting Started
#### Prerequisites
- Go 1.25+
- MySQL instance
- [Air](https://github.com/air-verse/air) (optional, for live reload)

#### Installation
1. Clone the repository:
   ```bash
   git clone git@github.com:kingoflongevity/study_go_web.git
   cd study_go_web
   ```
2. Install dependencies:
   ```bash
   go mod download
   ```
3. Configure your database in `configs/config.yaml`.

#### Running the Application
**With Air (Recommended):** `air`  
**Standard Go run:** `go run cmd/server/main.go`

---

## 🛣 API Endpoints / 接口清单
- `GET /api/v1/users/:id` - Get user by ID / 根据 ID 获取用户
- `POST /api/v1/users` - Create user / 创建用户
- `DELETE /api/v1/users/:id` - Delete user / 删除用户
