## 日志-列表
```
---------------- 请求 ----------------
{
    "api": "goback.log.list",
    "data": {
        "page":      1,  // [int] 分页
        "page_size": 1,  // [int] 每页显示几条数据 
        "type":      1,  // [int] 类型 
        "start":     1,  // [int] 开始时间戳
        "end":       1,  // [int] 结束时间戳
    }
}

---------------- 响应 ----------------
{
    "code": 10000,
    "msg": "",
    "data": {
        "datas": [ // [array[object]] 列表数据
            {
                "id":         1,  // [int] 配置ID 
                "adminer_id": 1,  // [int] 操作人
                "type":       1,  // [int] 类型
                "msg":        "", // [string] 信息  
                "target":     1,  // [int] 目标关联ID
                "created_at": 1,  // [int] 创建时间
            },
            ........
        ],
        "count":     1, // [int] 数据总数
        "page":      1, // [int] 当前页
        "page_size": 1, // [int] 每页显示条数 
        "type_names: {"1":"登录","2":"操作"}, // [string] 日志类型
    }
}
```

## 日志-查询
```
---------------- 请求 ----------------
{
    "api": "goback.log.info",
    "data": {
        "id": 1, // [int] 配置ID
    }
}
---------------- 响应 ----------------
{
    "code": 10000,
    "msg": "",
    "data": {
        "id":         1,  // [int] 配置ID 
        "adminer_id": 1,  // [int] 操作人
        "type":       1,  // [int] 类型
        "msg":        "", // [string] 信息  
        "target":     1,  // [int] 目标关联ID
        "created_at": 1,  // [int] 创建时间
    }
}
```
 