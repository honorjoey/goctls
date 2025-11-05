# Goctls

## Goctls is a powerful tools for simple admin to gen codes, based on goctl | Goctls 是一个为 simple admin 设计的代码生成工具, 基于 go zero 官方工具 goctl 开发而成
Based on https://github.com/honorjoey/goctls
修改适合单体服务的代码生成
goctls api ent 增加sub_api_dir参数，表示api子目录，生成的api会放在对应子目录下，即原来的service_name 和 子路径分成两个参数，原service_name 只表示服务名, 其他部分不变


## Install | 安装方法

```shell
go install github.com/honorjoey/goctls@latest
```

## Features | 特性

- go-swagger : it is different with origin which uses @doc comments
- multi-language support
- optimize error message
- fully support validator and easy to use
- code auto generation for API, RPC and web
- error handling which support multi-languages
- several plugins such as RocketMQ, GORM
- group rpc logic
- fully support Ent code generating
---
- go-swagger : 基于go-swagger而不是官方的@doc注解
- 多国语言支持
- 优化错误信息处理,支持多语言错误
- 简单易用的校验器
- 支持代码生成，生成API,RPC 和 web 端的CRUD代码
- 支持多种额外插件如GORM, RocketMQ
- 对Simple Admin 的针对性优化
- rpc logic group分组
- 全面支持 Ent 代码生成 

## 本项目遵循 GPL 3.0 协议， 使用请注明出处. 原许可证位于 originallicense 文件夹下，为 MIT 协议.
## This project uses the GPL 3.0 license, please indicate the source when using it. The original license is located in the originallicense folder under the MIT license.

> [go zero原版工具 | the original tools](https://github.com/zeromicro/go-zero/tree/master/tools/goctl)

## Simple Admin Core [https://github.com/suyuan32/simple-admin-core](https://github.com/suyuan32/simple-admin-core)