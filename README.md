# pscan

![Language](https://img.shields.io/badge/language-golang-blue.svg)

> [!IMPORTANT]
> 仅供学习交流使用，请勿用于非法用途

# Introduction

安全编程大作业

轻量级网络扫描器，后端采用golang 语言开发，web框架采用了gin微服务框架，使用websocket进行消息实时推送，支持以下功能：
-	主机存活扫描：通过ICMP 探测主机是否在线
-	端口开放扫描：检测目标主机上的端口开放情况
-	服务识别扫描：识别开放端口上的服务类型

# project struct

```
pscan/
├── main.go                 # 程序入口，初始化配置并启动服务
├── go.mod                  # Go 模块配置文件
├── plugins/                # 插件模块
│   ├── log.go              # 记录程序运行日志
│   ├── parseIPs.go         # 解析并验证输入的 IP 地址范围
│   ├── parsePorts.go       # 解析端口范围，支持单端口或区间格式
│   ├── ping.go             # 实现主机存活探测功能
│   ├── portscan.go         # 执行端口扫描任务
│   └── svcDetect.go        # 检测开放端口上的服务类型
├── web/                    # Web 模块
│   ├── controller/         # 控制器
│   │   ├── IndexController.go # 提供首页控制逻辑
│   │   └── ScanController.go  # 处理扫描请求并返回结果
│   ├── public/             # 前端页面
│   │   └── index.html      # 支持用户输入和结果显示
│   ├── router/             # 路由模块
│   │   └── routers.go      # 定义路由规则，关联控制器
│   └── service/            # 服务模块
│       └── service.go      # HTTP服务
```

# environment

> develop env

Go版本： 1.23.1

浏览器：支持 HTML5 的现代浏览器（推荐 Chrome）

> deploy env

操作系统：Windows、Linux或MacOS跨平台

端口：默认监听8989端口（可在web/services/service.go文件中修改）

# Usage

1.	编译：go build main.go命令可编译为适合运行环境的可执行文件
2.	运行：运行编译得到的可执行文件即可，linux和macos需要赋予可执行权限（chmod +x main）
3.	WebUI交互：项目启动后访问 http://localhost:8989 即可进入webui，然后按照页面提示进行操作
4.	扫描：填入必要的参数后，选择合适的扫描模式，即可开始扫描，实时扫描结果会打印到前端页面，完成扫描后会输出 `[+] Done` 字样
