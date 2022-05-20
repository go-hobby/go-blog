# 简介
```
 根据 blibli 的kratos  开发的单个grpc服务
 
 脚手架核心基于 [kratos](https://github.com/go-kratos/kratos) 框架和 [wire](https://github.com/google/wire) 依赖注入框架，请先自行根据官方文档学习
  
  
```
# 目录结构

```
|-- cmd # 编译入口
|   `-- app
|-- docs # 文档目录
|-- etc # 配置文件目录
|-- internal
|   `-- app
|       |-- command # 命令行功能模块
|       |   |-- handler
|       |   `-- script # 临时脚本
|       |-- component # 功能组件，如：db, redis  gorm 等
|       |-- config # 配置模型
|       |-- cron # 定时任务功能模块
|       |   `-- job
|       |-- model # 数据库模型
|       |-- pkg # 功能类库
|       |-- repository # 数据处理层
|       |-- service # 业务逻辑层
|       |-- test // 测试
|       `-- transport
|           |-- grpc
|           |   |-- api # proto 文件目录
|           |   |-- handler # 控制层
|           |   `-- middleware # 中间件
|           `-- http
|               |-- api # swagger 文档
|               |-- handler # 控制层
|               |-- middleware # 中间件
|               `-- router # 路由
|-- logs # 日志目录
|-- pkg # 功能类库
`-- proto # 第三方 proto 文件目录
```



# 如何运行

1. `安装kratos`
2. `修改etc 配置文件` 方式
3. `使用命令  kratos  run ` 方式


