# goback
Golang业务后端开发的适用性探究

# 使用
#### 模块 
`go get -u github.com/rwcoding/goback`

#### 测试
```shell 
go run examples/main.go

查看：http://127.0.0.1:9090
```

# 发布版本测试
+ 导入数据库（MySQL） `goback.sql`
+ 配置 `config.json`
+ 运行 `goback.linux --conf /etc/config.json`, 不指定配置文件，默认读取可执行文件当前目录的 `config.json` 文件
+ windows 系统使用 `goback.windows.exe`
+ 测试打包文件仅仅 `amd64` 环境