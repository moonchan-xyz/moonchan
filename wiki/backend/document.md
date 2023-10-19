# 我也不懂

~~这里记录已经实现的部分~~
业务的具体流程


[swagger API](./swagger.md)
[mastodon API 说明](./mastodon.md)
[gin 使用](./gin.md)
[database 说明](./database.md)
[gorm 使用](./gorm.md)
[develop](./develop.md)
[document](./document.md) 这个是干嘛的，~~log写这？~~
[readme.md](../../readme.md) log写这里

**基本大修了，记得改**
**2023年9月29日19点30分**

# ====

mastodon api的function直接调用core提供的方法
core提供的方法是基础的不冲突的？
到时候再吃屎吧。flag。
好想0721

层级大概是
mastodon
mastodon/* 具体逻辑
core/* 碎屑的业务单元


## core
姑且是核心吧

## db
db别的地方也碰不到的样子就放这里吧

## ~~core/entities~~ mastodon/entities
照着mastodon api写的

## masdtodon/*

### mastodon
mastodon api业务逻辑，啊？

这里gin相关

sub package里面gin无关。

### mastodon/accounts
里面的是accounts相关
structural在core/entities里面