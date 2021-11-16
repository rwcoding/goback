## 个人信息-查询
```
---------------- 请求 ----------------
{
    "api": "goback.profile.info", 
}

---------------- 响应 ----------------
{
    "code": 10000,
    "msg": "",
    "data": { 
        "username":     "", // [string] 用户名
        "name":         "", // [string] 名字
        "phone":        "", // [string] 手机
        "roles":        "", // [string] 角色id集合，逗号分隔
        "status":       1,  // [int] 状态
        "is_super":     1,  // [int] 是否超级管理员 1 是  0 否
        "created_at":   1,  // [int] 创建时间戳
        "role_names":   {"1":"管理","2":"开发"}, // [string] 角色id对应名字
        "status_names": {"1":"正常","2":"锁定"}, // [string] 状态对应
    }
}
```

## 个人信息-编辑
```
---------------- 请求 ----------------
{
    "api": "goback.profile.edit",
    "data": { 
        "name":  "", // [string] 名称
        "phone": "", // [string] 电话 
    }
}
---------------- 响应 ----------------
{
    "code": 10000,
    "msg": "", 
}
```


## 个人信息-密码修改
```
---------------- 请求 ----------------
{
    "api": "goback.profile.password",
    "data": { 
        "old_password": "", // [string] 原密码
        "new_password": "", // [string] 新密码
    }
}
---------------- 响应 ----------------
{
    "code": 10000,
    "msg": "", 
}
```