# HelpPackage
[![GoDoc](https://godoc.org/github.com/go-redis/redis?status.svg)](https://pkg.go.dev/github.com/Cc360428/HelpPackage?tab=doc)
[![Build Status](https://travis-ci.com/Cc360428/HelpPackage.svg?branch=master)](https://travis-ci.com/Cc360428/HelpPackage)      
[Reference](https://pkg.go.dev/github.com/Cc360428/HelpPackage?tab=doc)  

golang Development Kit
- 添加使用
```go
go get github.com/Cc360428/HelpPackage
```
## 目录说明
```shell
.
├── LICENSE
├── README.md
├── beego_utils
│   ├── request_helper.go
│   └── result.go
├── gin_utils
│   ├── json.go
│   ├── limiter.go
│   ├── limiter_test.go
│   ├── request.go
│   ├── response.go
│   └── util.go
├── go.mod
├── go.sum
├── help_package.go
└── utils
    ├── aliyun
    ├── async.go
    ├── async_test.go
    ├── asynchronous.go
    ├── email.go
    ├── eureka_client
    ├── excel.go
    ├── hs_token.go
    ├── hs_token_test.go
    ├── image
    ├── ip.go
    ├── ip_test.go
    ├── json.go
    ├── logs
    ├── map.go
    ├── map_test.go
    ├── mongodb
    ├── mysql_help
    ├── protoc
    ├── ratelimit
    ├── redis
    ├── redisUtil.go
    ├── regexp.go
    ├── regexp_test.go
    ├── snowflake.go
    ├── snowflake_test.go
    ├── string.go
    ├── swagger
    ├── time.go
    ├── time_test.go
    ├── token
    ├── util.go
    ├── util_test.go
    └── uuid

15 directories, 35 files
```
### beego_utils
>beego web 框架工具集成
### gin_utils
>gin web 框架集成工具
### utils
>golang 开发时候的工具包（aliyun、logs、mogoodb、mysql、proto、redis、token、utils）