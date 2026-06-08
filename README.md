# HelpPackage

[![Language](https://img.shields.io/badge/Language-Go-blue.svg)](https://golang.org/)
[![v1.20](https://img.shields.io/github/go-mod/go-version/Cc360428/HelpPackage)](https://github.com/Cc360428/HelpPackage)
[![stars](https://img.shields.io/github/stars/Cc360428/HelpPackage?style=social)](https://github.com/Cc360428/HelpPackage/stargazers)
[![pull](https://img.shields.io/github/issues-pr/Cc360428/HelpPackage)](https://github.com/Cc360428/HelpPackage)
[![Download](https://img.shields.io/github/downloads/Cc360428/HelpPackage/total)](https://github.com/Cc360428/HelpPackage)
[![GoDoc](https://godoc.org/github.com/go-redis/redis?status.svg)](https://pkg.go.dev/github.com/Cc360428/HelpPackage?tab=doc)
[![Build Status](https://travis-ci.com/Cc360428/HelpPackage.svg?branch=master)](https://travis-ci.com/Cc360428/HelpPackage)
[![codecov](https://codecov.io/gh/Cc360428/HelpPackage/branch/master/graph/badge.svg)](https://codecov.io/gh/Cc360428/HelpPackage)

Golang 开发工具包 - HelpPackage 是一个包含多个常用功能模块的 Go 语言开发库，提供数据库、缓存、配置、消息队列、日志等常用功能的封装，帮助开发者快速构建应用程序。

## 安装

```shell
go get github.com/Cc360428/HelpPackage
```

## 可用包列表

### 云服务相关

- **aliyun**: 阿里云服务客户端，包含短信发送功能
- **apollo**: Apollo 配置中心客户端
- **object-storage**: 对象存储服务
  - minio: MinIO 对象存储客户端
  - oss: 阿里云 OSS 对象存储客户端

### 数据库与缓存

- **elasticsearch**: Elasticsearch 搜索引擎客户端
- **mongodbc**: MongoDB 数据库客户端
- **mysqlc**: MySQL 数据库客户端
- **redisc**: Redis 缓存客户端
- **sqlite3c**: SQLite3 数据库客户端

### 消息队列

- **kafkac**: Kafka 消息队列客户端
- **rabbitmqc**: RabbitMQ 消息队列客户端

### 配置管理

- **conf**: 配置文件解析，支持 INI、TOML、YAML 格式

### 日志与监控

- **logger_zap**: 基于 Uber Zap 的日志库
- **logrus**: 基于 Logrus 的日志库，支持钉钉通知
- **monitor_message**: 监控消息通知（钉钉、Telegram、邮件）
- **prometheus**: Prometheus 监控指标

### Web 框架工具

- **web**: Web 框架工具
  - gin: Gin 框架工具类
  - iris: Iris 框架工具类

### 工具类

- **combinations**: 组合算法工具
- **image**: 图片处理（二维码生成等）
- **lru**: LRU 缓存实现
- **netc**: 网络工具（IP 定位、GPS）
- **other**: 其他工具函数（异步、Map、字符串、结构体等）
- **randc**: 随机数生成（雪花算法、权重随机等）
- **ratelimit**: 限流器
- **task**: 任务调度
- **timec**: 时间处理工具
- **token**: Token 生成
- **uuid**: UUID 生成

## 基本使用示例

### MySQL 客户端使用示例

```go
import "github.com/Cc360428/HelpPackage/mysqlc"

// 初始化
mysqlc.Init(&mysqlc.Config{
    Host:     "localhost",
    Port:     3306,
    User:     "root",
    Password: "password",
    Database: "test",
})

// 查询
var result []YourStruct
mysqlc.GetDB().Where("id = ?", 1).Find(&result)
```

### Redis 客户端使用示例

```go
import "github.com/Cc360428/HelpPackage/redisc"

// 初始化
redisc.Init(&redisc.Config{
    Addr:     "localhost:6379",
    Password: "",
    DB:       0,
})

// 设置值
redisc.GetClient().Set(ctx, "key", "value", 0)

// 获取值
value := redisc.GetClient().Get(ctx, "key").Val()
```

## 代码检查

```shell
go fmt $(go list ./... | grep -v /vendor/)
go vet $(go list ./... | grep -v /vendor/)
go test $(go list ./... | grep -v /vendor/)
```

## License

LICENSE
