# Echo Practice (study_go_web)

[简体中文](#简体中文) | [English](#english)

---

<a name="简体中文"></a>
## 简体中文

> **免责声明：** 本项目仅用于 **学习与练习**。旨在以 **整洁架构 (Clean Architecture)** 为核心思想，深入探索 Go 语言、Echo Web 框架及后端开发规范。

### 🚀 项目概览
这是一个基于 Go 和 Echo 框架构建的 RESTful API 示例项目。本项目通过实践 **整洁架构 (Clean Architecture)** 模式，将业务逻辑、数据访问与接口层进行严格分离，是学习大型 Go 项目架构设计的理想参考。

#### 核心设计理念
- **整洁架构驱动**：严格遵循 `Handler -> Service -> Repository` 分层模式。
- **关注点分离**：每一层职责明确，降低耦合度。
- **依赖注入**：通过手动注入依赖，提升代码的可测试性和灵活性。
- **面向接口编程**：利用 Go 的接口特性，确保层与层之间的解耦。

#### 功能特性
- 用户管理（增、删、查）
- 结构化配置管理 (Viper)
- 请求参数校验 (Validator)
- 数据库 ORM (GORM + MySQL)
- 开发环境热重载 (Air)

### 🛠 技术栈
- **语言：** Go
- **Web 框架：** [Echo](https://echo.labstack.com/)
- **ORM：** [GORM](https://gorm.io/)
- **数据库：** MySQL
- **配置管理：** [Viper](https://github.com/spf13/viper)
- **校验库：** [Validator](https://github.com/go-playground/validator)

### 🏗 项目结构
本项目遵循 [Standard Go Project Layout](https://github.com/golang-standards/project-layout) 规范：
```text
.
├── cmd/server/          # 应用程序入口 (组装各层依赖并启动)
├── configs/             # 配置文件 (YAML)
├── internal/
│   ├── config/          # 配置加载逻辑
│   ├── handler/         # 接口层 (Controller)：处理 HTTP 请求与响应
│   ├── model/           # 模型层 (Domain Model)：定义数据结构与 GORM 模型
│   ├── repository/      # 持久层 (Data Access)：封装数据库操作
│   └── service/         # 业务层 (Business Logic)：核心业务规则处理
└── pkg/                 # 公共工具库 (非业务相关的通用代码)
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

> **Disclaimer:** This project is for **learning and practice purposes only**. It focuses on mastering **Clean Architecture** principles while exploring the Go language and the Echo web framework.

### 🚀 Project Overview
A RESTful API built with Go and Echo, designed as a practical implementation of **Clean Architecture**. By strictly decoupling business logic, data access, and the interface layer, this project serves as a reference for building maintainable and scalable Go applications.

#### Core Architectural Concepts
- **Clean Architecture Driven**: Strictly follows the `Handler -> Service -> Repository` layering.
- **Separation of Concerns**: Each layer has a single responsibility, reducing tight coupling.
- **Dependency Injection**: Dependencies are manually injected to improve testability and flexibility.
- **Interface-Oriented**: Leverages Go interfaces to ensure decoupling between layers.

#### Features
- User Management (Create, Read, Delete)
- Structured Configuration Management (Viper)
- Request Validation (Validator)
- Database ORM (GORM + MySQL)
- Hot-reloading for development (Air)

### 🛠 Tech Stack
- **Language:** Go
- **Web Framework:** [Echo](https://echo.labstack.com/)
- **ORM:** [GORM](https://gorm.io/)
- **Database:** MySQL
- **Config Management:** [Viper](https://github.com/spf13/viper)
- **Validation:** [Validator](https://github.com/go-playground/validator)

### 🏗 Project Structure
This project follows the [Standard Go Project Layout](https://github.com/golang-standards/project-layout):
```text
.
├── cmd/server/          # Entry point (Dependency assembly & Server start)
├── configs/             # Configuration files (YAML)
├── internal/
│   ├── config/          # Config loading logic
│   ├── handler/         # Interface Layer (Controller): Handles HTTP req/res
│   ├── model/           # Model Layer (Domain Model): Data structures & GORM models
│   ├── repository/      # Persistence Layer (Data Access): Encapsulates DB operations
│   └── service/         # Service Layer (Business Logic): Core business rules
└── pkg/                 # Shared Utilities (Generic, non-business code)
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
