## 请求

+ 统一 **http post** 请求
+ **Header** 中增加 `Go-Api` 字段，接口路由
+ **Header** 中增加 `Go-Time` 字段，请求时间
+ **Header** 中增加 `Go-Token` 字段，登陆凭证，没有为空字符串
+ **Header** 中增加 `Go-Sign` 字段，数据签名，签名方式见下
+ 请求参数
```
{
   "api": "string", //接口路径
   "data": "object"  //请求数据
}
```
+ 例子：
```
Content-Type: application/json
Go-Sess: aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa
Go-Sign: xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
Go-Api: login
Go-Time: 1456787654

{"path":"login","data":{"uername":"lucy","password":"123"}}
```

+ 签名方式
```
key：签名的一个参数，用户登陆成功后，服务器会传回该参数。没有登陆时为空字符串
签名使用md5,方法为：md5(api + body + time + sess + key)
-----------
let api   = "/api"
let time  = 123123123
let sess  = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
let body  = JSON.stringify({"path":"/api"})
let key   = "kkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkk"
let sign  = md5(api + body + time + sess + key)
```

+ 返回数据
```
格式：JSON
code [int]    错误码， 10000 默认成功，其他失败，详细见 code.md
msg  [string] 错误信息
data [object] 详细数据
{
    "code": 10000,
    "msg":  "操作成功",
    "data": {"name":"lucy"}
}
```