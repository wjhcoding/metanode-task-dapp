# 📝 Metanode Task DApp

## 🚀 项目概览
metanode-task-dapp 是一个基于 Go 的轻量级区块链实验项目，主要功能：
- 连接 Ethereum Sepolia 测试网络
- 查询区块信息
- 发送以太坊交易
- 提供 RESTful HTTP API 供前端或其他服务调用

---

## 📂 项目结构
````
metanode-task-dapp/01-task/
├── api/v1                  # HTTP 控制器层
│   ├── blockchain_api.go   # 区块链接口
│   └── user_api.go         # 用户相关接口
├── cmd
│   └── main.go             # 启动 HTTP 服务
├── config
│   └── toml_config.go      # 配置文件解析
├── config.toml             # 配置文件
├── internal
│   ├── dao                 # 数据库访问层
│   ├── middleware          # JWT 鉴权等中间件
│   ├── model               # 数据模型
│   ├── router              # 路由注册
│   └── service
│       └── blockchain      # 区块链服务模块
│           ├── client.go   # RPC 客户端初始化
│           ├── block.go    # 查询区块
│           └── tx.go       # 发送交易
└── pkg
    ├── common/response     # HTTP 响应封装
    └── global/log          # 日志模块
````

---

## ⚙️ 环境依赖

- **Go** >= 1.19  
- **MySQL** >= 5.7  
- **Gin**  
- **GORM**  
- **Viper**  
- **JWT-go**  
- **Zap**

安装依赖：
```bash
go mod init github.com/wjhcoding/metanode-task-dapp/01-task
go mod tidy
````

---

## 🧩 配置文件（config.toml）

````toml
AppName = "metanode-task-dapp/01-task"

[Log]
Path = "./logs"
Level = "info"

[StaticPath]
FilePath = "./uploads"

[Blockchain]
RPC_URL = "https://sepolia.infura.io/v3/你的API_KEY"
PrivateKey = "你的私钥（仅测试用）"
GasLimit = 21000
GasTipCapGwei = 2
GasFeeCapGwei = 100
````

---

## 🏃‍♂️ 启动项目

### 1️⃣ 运行 MySQL 并导入表结构

````bash
go mod init github.com/wjhcoding/metanode-task-dapp/01-task
go mod tidy
````

### 2️⃣ 启动项目

````bash
go run cmd/main.go
````

服务器默认启动在：

````
http://localhost:8888
````

---

## 🔗 API 接口示例

### 🧍 用户注册

`POST /api/v1/user/register`

````json
{
  "username": "wjh",
  "password": "123456",
  "email": "wjh@example.com"
}
````

### 🔑 用户登录

`POST /api/v1/user/login`

````json
{
  "username": "wjh",
  "password": "123456"
}
````

返回：

````json
{
  "code": 200,
  "msg": "success",
  "data": {
    "token": "<JWT_TOKEN>"
  }
}
````

### 📝 发送交易接口

`POST /api/v1/blockchain/tx`

````json
{
  "to": "0xRecipientAddress",
  "amount": 0.001
}
````

Header：

````
Authorization: Bearer <JWT_TOKEN>
````

---

## 🪵 日志示例

项目运行后会在 `logs` 目录生成日志文件：

````
logs/
 ├── app.log
````

---

## 👨‍💻 作者信息

* 作者：**wjhcoding**
* 项目地址：[GitHub](https://github.com/wjhcoding/metanode-task-dapp/01-task)
* 邮箱：`wjhcoding@example.com`