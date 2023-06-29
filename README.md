# table-migration

-------
> 基于go语言实现，reason：因为go语言特性
## 项目介绍
### 背景
>  在本人任职的公司中，由于数据量过大导致部分查询慢相关的问题，严重影响了业务逻辑，基于此情况该项目诞生。
### 结构说明（根据[golang-standards](https://github.com/golang-standards/project-layout)规范进行标准划分定义）
> 该模块用于介绍项目结构


| --- configs：系统配置文件存放目录<br/>
  --- build：打包完成的二进制文件会输出到这里，提供window和linux的版本<br/>
  --- deployments: 容器化部署提供在这里<br/>
  --- pkg:源程序代码<br/>
  |---config<br/>
&nbsp;  |---dao ：访问层交互                                              <br/>
&nbsp;  |---model：orm对象                                            <br/>
&nbsp;  |---service：业务处理                                          <br/>
&nbsp;  main.go:程序入口<br/>
  --- scripts:相关脚本<br/>
  ---testing：测试代码<br/>
  ---vendor：项目依赖<br/>

### 功能
> 此项目提供的能力描述
* 通过读取migration.ini文件,加载当前项目中的所有待移动的表配置，当前文件属于热加载配置文件
* 并发处理数据迁移操作，各个配置间互不印象
* 若存在移表失败，则进行数据回写入原表，并且记录当前迁移表的数据以及迁移的最后主键id
## 框架选型
###  日志
  > 为了和市面商进行统一，考虑还是选择使用最广的[logrus](https://github.com/sirupsen/logrus)
### 数据库
  > [gorm](https://github.com/go-gorm/gorm) 对于此项目来说比较方便，目前考虑orm，后续可替代
### 配置文件
  > 个人觉得[ini](https://github.com/go-ini/ini)文件比较清晰