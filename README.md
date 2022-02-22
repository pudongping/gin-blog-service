# gin-blog-service

> 学习 gin 框架写的一个简易博客 http api 服务

## 启动项目

- 下载相关依赖扩展

```shell
go mod tidy
```

- 打开项目目录下的 `sql/init.sql` 文件，执行里面的初始化 sql 语句，创建数据表

- 执行以下命令，并修改相关配置信息

```shell
cp configs/config.yaml.example configs/config.yaml

# 编辑相关配置信息
vim configs/config.yaml
```

**关于读取配置**

> [go-bindata](https://github.com/go-bindata/go-bindata) 库可以将数据文件转换为 Go 代码。
> 因此读取配置信息也可以通过 [go-bindata/go-bindata](https://github.com/go-bindata/go-bindata/) 包提供的方式来读取。使用方式如下：

1. 安装 `go-bindata/go-bindata` 包

```shell
go get -u github.com/go-bindata/go-bindata/...
```

2. 将配置文件生成 go 代码

```shell
# 执行这条命令后，会将 `configs/config.yaml` 文件打包，并通过 `-o` 参数选择指定的路径输出到 `configs/config.go` 文件中
# 再通过 `-pkg` 选项指定生成的包名为 `configs`
go-bindata -o configs/config.go -pkg-configs configs/config.yaml
```

3. 读取文件中的配置信息

```go
data, err := configs.Asset("configs/config.yaml")

if err == nil {
    fmt.Println(string(data))
}
```

- 执行以下命令运行项目

```shell
# 普通启动项目
go run main.go

# 以自定义模式启动项目，可以通过命令行参数设置项目启动的端口（port）、启动模式（mode）、配置文件读取文件目录（config）
go run main.go -port=8001 -mode=debug -config=configs/
```

- 请求接口示例

```shell
# 请求授权登录接口
curl --location --request POST 'http://127.0.0.1:8000/auth' \
--form 'app_key="test"' \
--form 'app_secret="123456"'

# 请求获取标签列表接口，需要携带 `/auth` 接口返回的 token
# token 的携带，支持 get、post 请求方式，或者将 token 放置 header 头中
curl --location --request GET 'http://127.0.0.1:8000/api/v1/tags' \
--header 'Token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhcHBfa2V5IjoiNTM0YjQ0YTE5YmYxOGQyMGI3MWVjYzRlYjc3YzU3MmYiLCJhcHBfc2VjcmV0IjoiZTYwYmFmNjEzMmE3NDFmNjAyNmRlYjM0OTg4ZWRmMjkiLCJleHAiOjE2NDU0MTc2MDIsImlzcyI6Imdpbi1ibG9nLXNlcnZpY2UifQ.m0pEQgUNVTVlw8cudtbGXEy0PibxgfgUTnM6TGpCfsY'
```

- 若需要链路追踪，则需要（非必要实现）

```shell
# 使用 docker 安装 Jaeger 分布式链路追踪系统
docker run -d --name jaeger \
  -e COLLECTOR_ZIPKIN_HOST_PORT=:9411 \
  -p 5775:5775/udp \
  -p 6831:6831/udp \
  -p 6832:6832/udp \
  -p 5778:5778 \
  -p 16686:16686 \
  -p 14268:14268 \
  -p 14250:14250 \
  -p 9411:9411 \
  jaegertracing/all-in-one:1.28

# 如果需要更改链路追踪相关配置，详见 `main.go@setupTracer()` 方法

```

## 编译及打包

- 将编译信息写入二进制文件中

```shell
# 编译项目二进制文件时，可以通过 `ldflags` 工具，将一些编译相关的信息写入二进制文件中，方法日后查看二进制文件相关的信息
go build -ldflags \
"-X main.buildTime=`date +%Y-%m-%d,%H:%M:%S` -X main.buildVersion=1.0.0 -X main.gitCommitID=`git rev-parse HEAD`"
```

- 查看编译后的二进制文件和版本信息

```shell
./gin-blog-service -version

# output
# build_time: 2022-02-19,20:39:04
# build_version: 1.0.0
# git_commit_id: 23a5400593f3a4da768c10743e20bc035936f577
```

## 其他

- 查看项目接口文档

```shell
curl http://127.0.0.1:8000/swagger/index.html
```

- 查看 jaeger ui 控制面板

```shell
curl http://127.0.0.1:16686/search?service=gin-blog-service
```

- 查看请求日志

```shell
tail -f storage/logs/app.log
```