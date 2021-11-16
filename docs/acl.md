## 角色权限查询
```
---------------- 请求 ----------------
{
    "api": "goback.acl.role.query",
    "data": {
        "role_id": 1,  // [int] 角色ID 
    }
}

---------------- 响应 ----------------
{
    "code": 10000,
    "msg": "",
    "data": {
        "role_id": 1, // [int] 角色ID
        "role_name": "", // [string] 角色名称
        "groups": [ // [array[object]] 分组
            {
                "id":       1,  // [int] 分组id
                "name":     "", // [string] 分组名称
            },
            ........
        ],
        "permissions": [ // [array[object]] 权限列表
            {
                "id":         1,  // [int] 权限id
                "name":       "", // [string] 权限名称
                "permission": "", // [string] 权限标识
            },
            ........
        ],
    }
}
```

## 角色权限设置
```
---------------- 请求 ----------------
{
    "api": "goback.acl.role.set",
    "data": {
        "role_id": 1,  // [int] 角色ID 
        "permissions": "", // [string] 权限Id集合
    }
}

---------------- 响应 ----------------
{
    "code": 10000,
    "msg": "", 
}
```


## 批量操作-权限查询
```
---------------- 请求 ----------------
{
    "api": "goback.acl.batch.query",
}

---------------- 响应 ----------------
{
    "code": 10000,
    "msg": "",
    "data": {
        "roles": [ // [array[object]] 角色列表
            {
                "id":    1,  // [int] 角色id
                "name":  "", // [string] 角色名称
            },
            ........
        ],
        "groups": [ // [array[object]] 分组
            {
                "id":       1,  // [int] 分组id
                "name":     "", // [string] 分组名称
            },
            ........
        ],
        "permissions": [ // [array[object]] 权限列表
            {
                "id":         1,  // [int] 权限id
                "name":       "", // [string] 权限名称
                "permission": "", // [string] 权限标识
            },
            ........
        ],
    }
}
```

## 批量操作-权限设置
```
---------------- 请求 ----------------
{
    "api": "goback.acl.batch.set",
    "data": {
        "type": 1, // [int] 类型 1 新增  2 删除
        "roles": "",  // [string] 角色ID集合，逗号分隔
        "permissions": "", // [string] 权限Id集合，逗号分隔
    }
}

---------------- 响应 ----------------
{
    "code": 10000,
    "msg": "", 
}
```