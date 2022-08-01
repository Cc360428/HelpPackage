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
├── HelpPackage
├── LICENSE
├── README.md
├── aliyun
│   ├── client.go
│   ├── model.go
│   └── sms.go
├── async.go
├── async_test.go
├── buildtime.go
├── combinations
│   ├── combinations.go
│   └── combinations_test.go
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
├── go.mod
├── go.sum
├── help_package.go
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
├── kafkac
│   ├── docker-compose.yml
│   ├── example.go
│   ├── kafka.go
│   └── kafka_test.go
├── logger_zap
│   ├── zap.go
│   └── zap_test.go
├── logrus
│   ├── dingrus.go
│   ├── logrus.go
│   ├── logrus_test.go
│   ├── readme.md
│   └── utils
├── lru
│   ├── lru.go
│   ├── lru_test.go
│   └── readme.md
├── map.go
├── map_test.go
├── mongodbc
│   ├── docker-compose.yml
│   ├── example.go
│   ├── mongodb.go
│   └── mongodb_test.go
├── monitor
│   ├── conf
│   ├── monitor.go
│   └── web_server.go
├── monitor_message
│   ├── dd.go
│   ├── dd_test.go
│   ├── telegram.go
│   └── telegram_test.go
├── monitor_message.go
├── mysql
│   ├── affair.go
│   └── affair_test.go
├── netc
│   ├── GPS.go
│   ├── GPS_test.go
│   ├── ip.go
│   └── ip_test.go
├── protoc
│   ├── server.bm.go
│   ├── server.pb.go
│   └── server.proto
├── rabbitmqc
│   ├── docker-compose.yml
│   ├── example.go
│   ├── rabbitmq.go
│   └── rabbitmq_test.go
├── randc
│   ├── README.md
│   ├── random.go
│   ├── random_test.go
│   ├── randomness.go
│   ├── randomness_test.go
│   ├── snowflake.go
│   ├── snowflake_test.go
│   ├── weight.go
│   └── weight_test.go
├── randomness.go
├── randomness_test.go
├── ratelimit
│   ├── README.md
│   ├── ratelimit.go
│   ├── ratelimit_test.go
│   └── reader.go
├── redisc
│   ├── docker-compose.yml
│   ├── example.go
│   ├── redis.go
│   └── redis_test.go
├── regexp.go
├── regexp_test.go
├── string.go
├── string_test.go
├── struct.go
├── struct_test.go
├── swagger
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── swagger.go
├── task
│   ├── README.md
│   ├── task.go
│   └── task_test.go
├── timec
│   ├── format.go
│   ├── time.go
│   └── time_test.go
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
└── web
    ├── beego
    ├── gin
    └── iris

34 directories, 124 files
```
### web
- beego: beego web utils
- gin: gin web utils
- iris: iris web utils
- utils
### Other 
- log
- aliyun
- conf
- lru
- ORM
### Code Check
```shell
go fmt $(go list ./... | grep -v /vendor/)
go vet $(go list ./... | grep -v /vendor/)
go test $(go list ./... | grep -v /vendor/)
```