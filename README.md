# Meta

## 为什么会有这个项目

[Meta](https://github.com/one-meta/meta)，到底为何物？其实就是我想用 Go 来写项目，不仅仅是重构魔方，后续新的项目应该也会基于这个框架。也有很多优秀的框架了，但是跟个人习惯还是有点差异，于是根据自己的需求把需要的功能加到 Web 框架里，就有了 Meta。

## 功能特性

1. 遵循 RESTful API 设计规范。

2. 基于[Go Fiber](https://gofiber.io/)Web 框架，[中间件](https://docs.gofiber.io/api/middleware)
   丰富：Cache、CORS、CSRF、Limiter、Logger、RequestID 等。

3. 基于[Casbin](https://casbin.org/zh/) 的 RBAC 访问控制模型。

4. [gofiber 日志中间件](https://docs.gofiber.io/api/middleware/logger)
   记录 web 日志，[uber zap](https://github.com/uber-go/zap)记录程序日志，支持日志轮转。

5. 基于[JWT](https://github.com/golang-jwt/jwt)认证，用户密码使用[bcrypt](http://golang.org/x/crypto/bcrypt)加密。

6. [swagger api](https://gofiber.iogithub.com/swaggo/swag)自动生成 swagger 文档。

7. 基于[Entgo](https://entgo.io/zh)的数据库存储，[支持多种数据库后端](https://entgo.io/zh/docs/dialects)
   ：MySQL、MariaDB、PostgreSQL、CockroachDB (preview)、SQLite、Gremlin、TiDB (preview)；支持表自动创建及删除。

8. 基于[google Wire](https://github.com/google/wire)的依赖注入。

9. [gg](https://github.com/Xuanwo/gg) go 代码生成，自动生成 Api、Api 自动化测试案例（自动生成对象数据）、Controller、Service、Request 扩展参数；自动生成增删改查、批量创建、批量删除、任意字段搜索代码。

10. 扩展 Entgo 模板：增加各个实体的 Query、QuerySearch(任意字段搜索)、Create 的 SetEntity、UpdateOne 的 SetEntity 功能。

11. 基于 Entgo filter 的多租户实现。

12. 基于 httptest 和 fiber.Test 的单元测试，可对通用的方法实现自动化 Api 测试。

13. 基于[air](https://github.com/cosmtrek/air)go 程序热重启，可在开发中热重启。

14. 基于[gocron](github.com/go-co-op/gocron)的定时任务。

**涉及组件/框架**

[Go Fiber](https://gofiber.io/)：Web 框架、[Entgo](https://entgo.io/zh)：ORM、[viper](https://github.com/spf13/viper)：配置文件、[Casbin](https://casbin.org/zh/) ：RABC、[JWT](https://github.com/golang-jwt/jwt)、[uber zap 日志记录](https://github.com/uber-go/zap)、[lumberjack 日志轮转](https://github.com/natefinch/lumberjack)、[google Wire](https://github.com/google/wire)：依赖注入、[gg](https://github.com/Xuanwo/gg)：go 代码生成、[swagger api](https://github.com/swaggo/swag)、[air](https://github.com/cosmtrek/air)：go 程序热重启、[gocron](github.com/go-co-op/gocron)定时任务。

## 用法步骤【建议参考 [meta wiki](https://github.com/one-meta/meta/wiki/)】

1. ent 创建模块，ent 同目录执行（app 目录执行）
   `go run entgo.io/ent/cmd/ent new <moduleName>`

2. 根据实际修改`app/ent/schema`，字段、关系、Mixin 等，[参考](https://entgo.io/zh/docs/getting-started/)

3. 生成代码，ent 同目录执行（app 目录执行）
   `go generate ./ent`，没有报错即正常运行

   如果提示没有 entgo，则需要安装`go install entgo.io/ent/cmd/ent@latest`

   如果有报错，则可能需要删除 ent 目录下 extend\_\*.go 文件

   如果时间参数报错，则需要复制`data/query_param.go`到`app/ent/extend`，再`go generate ./ent`

4. 之后在项目根目录`air`运行即可

## 更多信息

1. 需要对 entgo 有一定了解
2. 权限管理需要了解 casbin
3. 依赖注入，可以参考已有的代码
