# database

负责 PostgreSQL 连接、GORM 初始化、数据库迁移和事务入口。

第一版可以把表结构初始化和 AutoMigrate 相关逻辑放在这里，不单独创建 `migrations/` 目录。

训练记录创建、编辑、删除这类需要同时操作多张表的场景，后续应通过事务处理。
