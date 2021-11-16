## 后台管理员列表
```
---------------- 请求 ----------------
{
    "api": "goback.adminer.list",
    "data": {
        "page":      1,  // [int] 分页
        "page_size": 1,  // [int] 每页显示几条数据
        "username":  "", // [string] 查询用户名 
    }
}

---------------- 响应 ----------------
{
    "code": 10000,
    "msg": "",
    "data": {
        "datas": [ // [array[object]] 列表数据
            {
                "id":       1,  // [int] 用户名
                "username": "", // [string] 用户名
                "name":     "", // [string] 姓名
                "phone":    "", // [string] 电话
                "roles":    "", // [string] 角色ID集合
                "status":   1,  // [int] 状态
                "is_super": 1,  // [int] 超级管理员
            },
            ........
        ],
        "count":     1, // [int] 数据总数
        "page":      1, // [int] 当前页
        "page_size": 1, // [int] 每页显示条数
        "role_names": {"1":"管理员","2":"开发"}, //角色 id对应名字
        "status_names": {"1":"正常","2":"锁定"}, //状态 值对应文本描述
    }
}
```

## 后台管理员信息
```
---------------- 请求 ----------------
{
    "api": "goback.adminer.info",
    "data": {
        "id": 1, // [int] 用户ID
    }
}
---------------- 响应 ----------------
{
    "code": 10000,
    "msg": "",
    "data": {
        "id": 1,        // [int] 用户ID
        "username": "", // [string] 用户名 
        "password": "", // [string] 密码
        "name": "",     // [string] 姓名
        "phone": "",    // [string] 电话
        "status": 1,    // [int] 状态
        "is_super": 0,  // [int] 是否超级管理员
        "role_names": {"1":"管理员","2":"开发"}, //角色 id对应名字
        "status_names": {"1":"正常","2":"锁定"}, //状态 值对应文本描述
    }
}
```

## 后台管理员添加
```
---------------- 请求 ----------------
{
    "api": "goback.adminer.add",
    "data": {
        "username": "", // [string] 用户名 
        "password": "", // [string] 密码
        "name": "",     // [string] 姓名
        "phone": "",    // [string] 电话
        "status": 1,    // [int] 状态
        "is_super": 0,  // [int] 是否超级管理员
    }
}

---------------- 响应 ----------------
{
    "code": 10000,
    "msg": "",
    "data": {
        "id": 1, // [int] 用户ID
    }
}
```

## 后台管理员编辑
```
---------------- 请求 ----------------
{
    "api": "goback.adminer.edit",
    "data": {
        "id":       1,  // [int] 用户ID  
        "password": "", // [string] 密码
        "name":     "", // [string] 姓名
        "phone":    "", // [string] 电话
        "status":   1,  // [string] 状态
        "is_super": 0,  // [string] 是否超级管理员
    }
}

---------------- 响应 ----------------
{
    "code": 10000,
    "msg": "",
}
```

## 后台管理员删除
```
---------------- 请求 ----------------
{
    "api": "goback.adminer.delete",
    "data": {
        "id": 1, // [int] 管理员ID
    }
}

---------------- 响应 ----------------
{
    "code": 10000,
    "msg": "",
}
```