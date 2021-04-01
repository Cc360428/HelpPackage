# HelpPackage
[![Language](https://img.shields.io/badge/Language-Go-blue.svg)](https://golang.org/)
[![v1.14](https://img.shields.io/github/go-mod/go-version/Cc360428/HelpPackage)](https://github.com/Cc360428/HelpPackage)
[![stars](https://img.shields.io/github/stars/Cc360428/HelpPackage?style=social)](https://github.com/Cc360428/HelpPackage/stargazers)
[![pull](https://img.shields.io/github/issues-pr/Cc360428/HelpPackage)](https://github.com/Cc360428/HelpPackage)
[![Download](https://img.shields.io/github/downloads/Cc360428/HelpPackage/total)](https://github.com/Cc360428/HelpPackage)
[![GoDoc](https://godoc.org/github.com/go-redis/redis?status.svg)](https://pkg.go.dev/github.com/Cc360428/HelpPackage?tab=doc)
[![Build Status](https://travis-ci.com/Cc360428/HelpPackage.svg?branch=master)](https://travis-ci.com/Cc360428/HelpPackage)
[![codecov](https://codecov.io/gh/Cc360428/HelpPackage/branch/master/graph/badge.svg)](https://codecov.io/gh/Cc360428/HelpPackage)   
golang Development Kit
- 添加使用
```shell
go get github.com/Cc360428/HelpPackage
```
## 目录说明
```readme
.
├── LICENSE
├── README.md
├── beego
│   ├── request_helper.go
│   └── result.go
├── gin
│   ├── json.go
│   ├── limiter.go
│   ├── limiter_test.go
│   ├── request.go
│   ├── response.go
│   └── util.go
├── go.mod
├── go.sum
├── help_package.go
├── iris
│   └── response.go
└── utils
    ├── GPS.go
    ├── GPS_test.go
    ├── aliyun
    │   ├── client.go
    │   ├── model.go
    │   └── sms.go
    ├── async.go
    ├── async_test.go
    ├── conf
    │   ├── README.md
    │   ├── ini
    │   ├── iniv1.go
    │   ├── iniv2.go
    │   ├── test.go
    │   ├── toml
    │   ├── toml.go
    │   ├── yaml
    │   └── yaml.go
    ├── dd.go
    ├── dd_test.go
    ├── elasticsearch
    │   ├── base.go
    │   └── readme.md
    ├── email.go
    ├── eureka_client
    │   ├── api.go
    │   ├── client.go
    │   ├── config.go
    │   ├── util.go
    │   └── util_test.go
    ├── hs_token.go
    ├── hs_token_test.go
    ├── image
    │   ├── golang.png
    │   ├── new_qr_code.png
    │   ├── qr_code.go
    │   ├── qr_code_test.go
    │   ├── qr_in.go
    │   ├── qr_in.png
    │   └── qr_in_test.go
    ├── ip.go
    ├── ip_test.go
    ├── logger
    │   └── main.go
    ├── logrus
    │   ├── dingrus.go
    │   ├── logruslogtest
    │   ├── main.go
    │   └── readme.md
    ├── logs
    │   ├── README.md
    │   ├── accesslog.go
    │   ├── alils
    │   ├── conn.go
    │   ├── console.go
    │   ├── es
    │   ├── file.go
    │   ├── jianliao.go
    │   ├── log.go
    │   ├── logger.go
    │   ├── multifile.go
    │   ├── slack.go
    │   └── smtp.go
    ├── map.go
    ├── map_test.go
    ├── mongodb
    │   ├── mongodb.go
    │   └── mongodb_test.go
    ├── monitor
    │   ├── conf
    │   ├── monitor.go
    │   └── web_server.go
    ├── mysql
    │   ├── affair.go
    │   └── affair_test.go
    ├── protoc
    │   ├── server.bm.go
    │   ├── server.pb.go
    │   └── server.proto
    ├── randomness.go
    ├── randomness_test.go
    ├── ratelimit
    │   ├── README.md
    │   ├── ratelimit.go
    │   ├── ratelimit_test.go
    │   └── reader.go
    ├── redis
    │   └── redis.go
    ├── redisUtil.go
    ├── regexp.go
    ├── regexp_test.go
    ├── snowflake.go
    ├── snowflake_test.go
    ├── string.go
    ├── string_test.go
    ├── struct.go
    ├── struct_test.go
    ├── swagger
    │   ├── docs.go
    │   ├── swagger.json
    │   └── swagger.yaml
    ├── task
    │   ├── README.md
    │   ├── task.go
    │   └── task_test.go
    ├── time.go
    ├── time_test.go
    ├── token
    │   ├── newToken.go
    │   └── readme.md
    ├── util.go
    ├── util_test.go
    ├── uuid
    │   ├── README.md
    │   ├── codec.go
    │   ├── codec_test.go
    │   ├── generator.go
    │   ├── generator_test.go
    │   ├── sql.go
    │   ├── sql_test.go
    │   ├── uuid.go
    │   └── uuid_test.go
    └── zap
        └── zap.go

30 directories, 112 files
```
### beego
>beego web utils
### gin
>gin web utils
### iris
>iris web utils
### utils
- 阿里云（短信发送）
- eureka（client）
- image（二维码）
- logs（日志）
- conf（配置文件读取）
- monitor（监听文件）
- mysql
- protoc
- ratelimit
- redis
- swagger
- task
- uuid
- email、Excel、swagger、ip、token、GPS、map、redis、regexp、snowflake、string、time、util
## commit 
```shell
go fmt $(go list ./... | grep -v /vendor/)
go vet $(go list ./... | grep -v /vendor/)
go test $(go list ./... | grep -v /vendor/)
```