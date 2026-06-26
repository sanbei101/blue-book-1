# 小蓝书Ai开发指南

## 1. 项目结构

```
├── cmd
│   └── main.go // 程序入口
├── go.mod
├── go.sum
├── internal
│   ├── api // 直接写handler+service
│   ├── db  // 数据库访问层,sqlc生成的代码放在这里
│   │   ├── copyfrom.go
│   │   ├── db.go
│   │   ├── models.go
│   │   ├── querier.go
│   │   ├── query.sql.go
│   │   └── store.go
│   └── pkg
│       └── render
│           └── render.go // 统一的响应渲染工具
├── sqlc
│   ├── query.sql
│   └── schema.sql
└── sqlc.yaml // sqlc配置文件
```

## 技术栈
- Go 1.26
- Chi 路由
- PostgreSQL 数据库
- pgx/v5 数据库驱动
- phuslu/log 日志库
- sqlc 数据库访问代码生成工具

## 开发规范
- 尽可能不要忽略err,调用数据库操作禁止忽略err!
- 代码需要尽可能简洁,多复用写好的函数,不要重复造轮子
- 每次跑完需要调用`golangci-lint run --fix`检查代码规范,确保没有报错

## Swagger 注释规范
- `@accept` 和 `@produce` 只在 `routes.go` 全局定义,其他文件禁止重复写
- `@Success` 必须使用泛型包裹: `render.Response[T]`,不能直接写裸类型
  - 单个对象: `@Success 200 {object} render.Response[listPostsResponse]`
  - 数组: `@Success 200 {object} render.Response[[]listPostsResponse]`
  - 无数据: `@Success 204 {object} render.ResponseWithoutData`
- 注释格式使用 tab 对齐,保持一致性
- 每次修改完成代码需要使用以下命令进行`swagger`生成:
```
swag init -g routes.go -d internal/api,internal/pkg/render --ot yaml
swag fmt -g routes.go -d internal/api
```
