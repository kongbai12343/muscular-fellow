# routes

负责集中注册 Gin 路由。

建议按功能拆分：

- auth routes
- user routes
- exercise routes
- workout routes
- stats routes

routes 只做路径和 controller 的绑定，不写业务逻辑。
