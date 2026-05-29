# controllers

MVC 的 Controller 层。

负责：

- 接收 HTTP 请求。
- 读取路径参数、查询参数和请求体。
- 调用 service。
- 返回统一 JSON 响应。

不要在 controller 里写复杂业务逻辑和数据库查询。
