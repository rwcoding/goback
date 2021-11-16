## 会话-列表
```
---------------- 请求 ----------------
{
    "api": "goback.session.list",
    "data": {
        "page":       1,  // [int] 分页
        "page_size":  1,  // [int] 每页显示几条数据 
        "adminer_id": 1,  // [int] 管理员ID
    }
}

---------------- 响应 ----------------
{
    "code": 10000,
    "msg": "",
    "data": {
        "datas": [ // [array[object]] 列表数据
            {
                "id":         1, // [int] 配置ID 
                "adminer_id": 1, // [int] 管理员
                "session_id": 1, // [string] session id
                "expire":     1, // [int] 过期时间   
                "created_at": 1, // [int] 创建时间
                "updated_at": 1, // [int] 更新时间
            },
            ........
        ],
        "count":     1, // [int] 数据总数
        "page":      1, // [int] 当前页
        "page_size": 1, // [int] 每页显示条数  
    }
}
```


## 会话-删除
```
---------------- 请求 ----------------
{
    "api": "goback.session.delete",
    "data": {
        "id":   1, // [int] 会话ID 
    }
}
---------------- 响应 ----------------
{
    "code": 10000,
    "msg": "", 
}
```

## 会话-清理
```
---------------- 请求 ----------------
{
    "api": "goback.session.clean",
    "data": {
        "type": 1, // [int] 类型 0 所有session, 1 用户会话  2 验证码 
    }
}
---------------- 响应 ----------------
{
    "code": 10000,
    "msg": "", 
}
```
