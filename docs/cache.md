## 缓存-列表
```
---------------- 请求 ----------------
{
    "api": "goback.cache.list",
    "data": {
        "page":      1,  // [int] 分页
        "page_size": 1,  // [int] 每页显示几条数据
        "sign":      "", // [string] 查询缓存KEY
    }
}

---------------- 响应 ----------------
{
    "code": 10000,
    "msg": "",
    "data": {
        "datas": [ // [array[object]] 列表数据
            {
                "id":        1,  // [int] 数据ID 
                "name":      "", // [string] 名称
                "sign":      "", // [string] 标识
                "data":      "", // [string] 数据 
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


## 缓存-查询
```
---------------- 请求 ----------------
{
    "api": "goback.cache.info",
    "data": {
        "id": 1, // [int] 数据ID
    }
}
---------------- 响应 ----------------
{
    "code": 10000,
    "msg": "",
    "data": {
        "id":        1,  // [int] 数据ID 
        "name":      "", // [string] 名称
        "sign":      "", // [string] 标识
        "data":      "", // [string] 数据 
        "updated_at": 1, // [int] 更新时间
    }
}
```


## 缓存-生成
```
---------------- 请求 ----------------
{
    "api": "goback.cache.generate", 
}
---------------- 响应 ----------------
{
    "code": 10000,
    "msg": "", 
}
```