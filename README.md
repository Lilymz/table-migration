# table-migration

-------
> 基于go语言实现，reason：因为go语言特性
## 项目介绍
### 背景
>  在本人任职的公司中，由于数据量过大导致部分查询慢相关的问题，严重影响了业务逻辑，基于此情况该项目诞生。
### 结构说明（根据[golang-standards](https://github.com/golang-standards/project-layout)规范进行标准划分定义）
> 该模块用于介绍项目结构
### 功能
## 框架选型
###  日志
  > 为了和市面商进行统一，考虑还是选择使用最广的[logrus](https://github.com/sirupsen/logrus)
### 数据库
  > [gorm](https://github.com/go-gorm/gorm) 对于此项目来说比较方便，目前考虑orm，后续可替代
### 配置文件
  > 个人觉得[ini](https://github.com/go-ini/ini)文件比较清晰