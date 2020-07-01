# activities

---
项目结构:

- common
  - errs: 错误定义
  - 公共常量

- gateway
  - handler: 处理
  - router: 路由

- library
  - clients
    - nats: 消费消息
    - 远程服务调用
  - config: 配置
  - logger: 日志
  - storage: 数据库存储
  - utils

- models: ORM模型
- scripts: 脚本

- service
  - famework: 活动抽象类&活动实现
  - 服务实现: 整个活动服务的封装

- logs
  - 日志分割

---
