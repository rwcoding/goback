# Gobui
Golang业务后端开发的适用性探究

# 使用
#### 模块 
`go get -u github.com/rwcoding/gobui`

#### 入口 
`main.go`
```go
package main

import (
	"github.com/rwcoding/gobui"
	"github.com/rwcoding/gobui/config"
	_ "github.com/rwcoding/gobui/module"
	"log"
)

func main() {
	app := gobui.DefaultApp()
	err := app.Run(config.Addr())
	if err != nil {
		log.Fatal(err)
	}
}
```

#### 配置 
`config.json`
```json
{
  "env": "dev",
  "addr": ":8080",
  "log": "",

  "db" : {
    "host": "127.0.0.1",
    "port": 3306,
    "username": "root",
    "password": "root",
    "dbname": "gobui",
    "charset": "utf8mb4",

    "pool_max_open": 100,
    "pool_max_idle": 50,
    "pool_max_life": 600
  }
}
```

#### 运行
```shell 
go run main.go

查看：http://127.0.0.1:8080
```

# 配置文件优先级
+ 运行时指定绝对目录：`--conf /etc/config.json`
+ 可执行文件当前目录的 `config.json` 文件
+ 运行终端所在目录的 `config.json` 文件

# 发布版本测试
+ 导入数据库（MySQL） `gobui.sql`
+ 配置 `config.json`
+ 运行 `gobui.linux --conf /etc/config.json`, 不指定配置文件，默认读取可执行文件当前目录的 `config.json` 文件
+ windows 系统使用 `gobui.windows.exe`
+ 测试打包文件仅仅 `amd64` 环境