# 后端项目骨架

本目录用于放置健身训练记录 App 的 Go 后端代码。

当前只搭建项目框架，不包含业务代码。正式开发时再初始化 Go module、安装依赖并补充入口文件。

## 技术栈

- Go
- Gin
- GORM
- PostgreSQL
- zap
- JWT
- bcrypt
- godotenv

## 目录说明

- `main.go`：服务启动入口，后续直接放在 `backend/` 根目录。
- `routes/`：集中注册路由，按功能把 URL 映射到 controller。
- `controllers/`：MVC 的 Controller 层，负责接收请求、读取参数、返回响应。
- `models/`：MVC 的 Model 层，负责 GORM 数据模型和表结构映射。
- `services/`：业务逻辑层，例如注册登录、创建训练、统计计算。
- `repositories/`：数据库访问层，封装 GORM 查询。
- `dto/`：请求和响应 DTO，避免 controller 直接暴露 model。
- `config/`：读取环境变量和应用配置。
- `database/`：数据库连接、迁移、事务入口。
- `logger/`：zap 日志初始化。
- `middleware/`：鉴权、请求日志、错误恢复、request id。
- `response/`：统一 JSON 响应结构，API 项目不单独做传统 View。
- `validator/`：请求参数校验。

## 入口职责

后续 `main.go` 负责：

- 加载配置。
- 初始化日志。
- 连接数据库。
- 在 `database/` 中执行表结构初始化或迁移。
- 注册路由。
- 启动 HTTP 服务。

## 本地环境

复制 `.env.example` 为 `.env`，再根据本机数据库配置修改。

```bash
cp .env.example .env
```

第一阶段建议只用 Docker 启动 PostgreSQL，Go 服务仍在本机直接运行。
