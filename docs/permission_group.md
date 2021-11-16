## 权限分组-列表
```
---------------- 请求 ----------------
{
    "api": "goback.permission.group.list",
    "data": {
        "page":      1,  // [int] 分页
        "page_size": 1,  // [int] 每页显示几条数据 
    }
}

---------------- 响应 ----------------
{
    "code": 10000,
    "msg": "",
    "data": {
        "datas": [ // [array[object]] 列表数据
            {
                "id":   1,  // [int] 分组ID 
                "name": "", // [string] 名称 
                "ord":  1,  // [int] 排序
            },
            ........
        ],
        "count":     1, // [int] 数据总数
        "page":      1, // [int] 当前页
        "page_size": 1, // [int] 每页显示条数 
    }
}
```

## 权限分组-查询
```
---------------- 请求 ----------------
{
    "api": "goback.permission.group.info",
    "data": {
        "id": 1, // [int] 分组ID
    }
}
---------------- 响应 ----------------
{
    "code": 10000,
    "msg": "",
    "data": {
        "id":   1,  // [int] 分组ID 
        "name": "", // [string] 名称
        "ord":  1,  // [int] 排序 
    }
}
```


## 权限分组-添加
```
---------------- 请求 ----------------
{
    "api": "goback.permission.group.add",
    "data": { 
        "name": "", // [string] 名称
        "ord":  1,  // [int] 排序 
    }
}
---------------- 响应 ----------------
{
    "code": 10000,
    "msg": "",
    "data": {
        "id":   1,  // [int] 分组ID  
    }
}
```


## 权限分组-编辑
```
---------------- 请求 ----------------
{
    "api": "goback.permission.group.edit",
    "data": {
        "id":   1,  // [int] 分组ID  
        "name": "", // [string] 名称
        "ord":  1,  // [int] 排序 
    }
}
---------------- 响应 ----------------
{
    "code": 10000,
    "msg": "",
    "data": {
        "id": 1,  // [int] 分组ID  
    }
}
```


## 权限分组-删除
```
---------------- 请求 ----------------
{
    "api": "goback.permission.group.delete",
    "data": {
        "id":   1, // [int] 分组ID 
    }
}
---------------- 响应 ----------------
{
    "code": 10000,
    "msg": "", 
}
```
