## 权限-列表
```
---------------- 请求 ----------------
{
    "api": "goback.permission.list",
    "data": {
        "page":      1,  // [int] 分页
        "page_size": 1,  // [int] 每页显示几条数据 
        "permission":"", // [string] 权限标识
    }
}

---------------- 响应 ----------------
{
    "code": 10000,
    "msg": "",
    "data": {
        "datas": [ // [array[object]] 列表数据
            {
                "id":         1,  // [int] 权限ID 
                "name":       "", // [string] 名称
                "permission": "", // [string] 标识
                "gid":        1,  // [string] 分组id  
                "type":       1,  // [string] 类型 
            },
            ........
        ],
        "count":     1, // [int] 数据总数
        "page":      1, // [int] 当前页
        "page_size": 1, // [int] 每页显示条数 
        "group_names": {"1":"日志","2":"配置"}, // [string] 分组id对应名字
        "type_names":  {"1":"接口","2":"自定义"}, // [string] 类型
    }
}
```

## 权限-查询
```
---------------- 请求 ----------------
{
    "api": "goback.permission.info",
    "data": {
        "id": 1, // [int] 权限ID
    }
}
---------------- 响应 ----------------
{
    "code": 10000,
    "msg": "",
    "data": {
        "id":         1,  // [int] 权限ID 
        "name":       "", // [string] 名称
        "permission": "", // [string] 标识
        "gid":        1,  // [string] 分组id  
        "type":       1,  // [string] 类型 
    }
}
```


## 权限-添加
```
---------------- 请求 ----------------
{
    "api": "goback.permission.add",
    "data": {  
        "name":       "", // [string] 名称
        "permission": "", // [string] 标识
        "gid":        1,  // [string] 分组id  
        "type":       1,  // [string] 类型 
    }
}
---------------- 响应 ----------------
{
    "code": 10000,
    "msg": "",
    "data": {
        "id":   1,  // [int] 权限ID  
    }
}
```


## 权限-编辑
```
---------------- 请求 ----------------
{
    "api": "goback.permission.edit",
    "data": {
        "id":         1,  // [int] 权限ID 
        "name":       "", // [string] 名称
        "permission": "", // [string] 标识
        "gid":        1,  // [string] 分组id  
        "type":       1,  // [string] 类型 
    }
}
---------------- 响应 ----------------
{
    "code": 10000,
    "msg": "",
    "data": {
        "id":   1,  // [int] 权限ID  
    }
}
```


## 权限-删除
```
---------------- 请求 ----------------
{
    "api": "goback.permission.delete",
    "data": {
        "id":   1, // [int] 权限ID 
    }
}
---------------- 响应 ----------------
{
    "code": 10000,
    "msg": "", 
}
```
