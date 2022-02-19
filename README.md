# gin-blog-service
write blog api service for learning gin framework

## 启动项目

```shell

# 以调试模式启动
go run main.go -port=8001 -mode=debug -config=configs/

```

## 编译信息写入二进制文件中

### 写入

```shell
go build -ldflags \
"-X main.buildTime=`date +%Y-%m-%d,%H:%M:%S` -X main.buildVersion=1.0.0 -X main.gitCommitID=`git rev-parse HEAD`"
```

### 查看编译后的二进制文件和版本信息

```shell
./gin-blog-service -version

# output
# build_time: 2022-02-19,20:39:04
# build_version: 1.0.0
# git_commit_id: 23a5400593f3a4da768c10743e20bc035936f577
```