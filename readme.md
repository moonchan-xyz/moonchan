# moonchan 的后端

- [swagger](swagger.md)
  - 虽然感觉是不是swaggo和swagger分开来
- [gin](gin.md)
  - 怎么接受数据


## 概述

大概会分为

- mastodon api
- activity pub worker
- futaba like api

```txt
┌─────────────┐           ┌────────────────┐               ┌───────────────────────┐
│             │           │                │               │                       │
│             │ ◄───────  │                │               │                       │
│             │           │                │  ◄──────────  │                       │
│  mastodon   │ ───────►  │   core_worker  │               │ activity pub api      │
│             │           │                │  ───────────► │                       │
│    api      │           │                │               │                       │
│             │           │                │               │                       │
└─────────────┘           │                │               │                       │
                          └─────┬──────────┘               │                       │
                                │    ▲                     └───────────────────────┘
                                │    │
                                │    │
                                │    │
                                ▼    │
                         ┌───────────┴────────────────────┐
                         │                                │
                         │                                │
                         │    data base                   │
                         │                                │
                         └────────────────────────────────┘

```

虽然为什么要写这个呢。夙愿吧。

会和每一个重写的 fedi 软件一起，淹没在汪洋大海里面

## 怎么测试

好像跑不起来
localhost 姑且是 3000
fedi 方法还没做。
还是 local 的 api。
麻了
到时候要仔细看 local 的方法作得对不对才行

## misc

需要配置的 key

```txt
DefaultHeader
DefaultAvatar
```

在`mastodon/profile/profile.go`
目的是为了配合 default 操作

# log

## 10/13
api当中accounts做了个架子
methods好多。要死了
只做出不影响前端工作的就好了吧。
还有activitypub的工作要做了
之后要不把swagger省略掉，要debug或者最后的时候补上
反正文档都是mastodon

## 10/12
file能接住了
hash也能接住了
swagger上传都处理好了
后面应该多写 swag，相关文档在wiki swagger.md
gin也可以写一下吧

## 10/11
http body是reader所以不能读两遍，gin调用了参照body的内容会读掉body
## 10/09

极限摸鱼，写一下 swag
接下来先把 swag 补完
考虑 param 的获取
然后考虑代码的获取
然后考虑逻辑

## 09/29

~~稍微想把 s2s 的打通~~
不用想了
先加 swag，然后再写 api
server to server （之后）
mastdon （现在）
moonchan （不怎么在乎）

把 POST，GET 什么的参数全都提取出来就是了。别的没啥
直接对 db 的影响可以先写上。
对其他 server 的影响先用 core 的 dummy 撑着
activitypub 分成 api 和 action？

## 09/23

@20 点 37 分
注意 source 和 account 的 field 一起更新？
现在是用 json 化写的
gorm 的关系实在用不来。

@23 点 20 分
想起来加了个 index 在 relationship 里面

怎么 git 啊

`mastodon/accounts.go`救救命啊

notification 你一定要加已读吗
要来不及做正经事情了。

@23 点 59 分
留给之后的 note

`mastodon`下面是 gin 的界面
子文件夹对应了 mastodon api 的方法们
请严格按照 mastodon 的 api 做
目前还在写
`entities`请完全对应 mastodon 的 enity，多余的关系请在 db 当中加表
（留下效率啊之类的坑以后说吧）

希望下次还能找到什么在哪里

## 09/20

@09 点 17 分
不知道写到哪里了
上次应该是测试 gorm

@11 点 59 分

```md
[custom_emoji.go](/back/core/entities/custom_emoji.go)
```

这么写链接

@12 点 04 分
查看 entities，然后下面似乎有关于他的方法

@12 点 36 分
写了 lookup，这个之后要补远程过程
待测试
等 create 有了一起来。
不写 db 里的 testcase 了，好像很麻烦。（估计会被自己打）

@12 点 53 分
绷不住了，form parameter 要之后自己测试的啊
完了越活越回去了，谁让你不补基础的
para[]还看不懂。
找个地方注册一个吧。log 一下

没验证

```go
name := c.PostForm("name")
message := c.PostForm("message")
```

几个 key 也没试过

简化流程直接通过把。。这里是留 todo

设置 profile 的时候
api 他妈没用啊！难绷。

写了 get，没测试
改成了 json 存进去，没问题吗

ReadRelationFollowedBy 这一系列可能要重点查。

我是不是要写个测试方案还是咋地

@17 点 46 分
写了 following relationship
怎么测试呃。
估计要写出 remote..

真的好麻烦
fedi.POST 出去要
签名
还要 local 做 objecct
方法估计要做到 core 里面。
乱包会出事情么

写了 webfinger 的 server 和 fetch
没测，但是是复制过来的

记得全加上 json。不然会变成原名 -- 这是在说啥。

[这里](/back/mastodon/accounts/update.go)
加上 patch 到 db 的逻辑
[前头](/back/mastodon/accounts.go)
记得做好先前处理。null 怎么弄啊。

```go
func form2data(c) (xx)

```

db 怎么处理 update 行为的。
要试试不然会不会盖住。
好困，睡了

## 09/15

_17 点 48 分_

> _18 点 04 分_
> 不知道要怎么安排目录呃

> _18 点 37 分_

加载好慢

记得 json 还是 gorm 的时候
对于 user defined type
有个 serializer 的大病
所以这个雷先趟了在说

绷不住了做一次测试好慢，还是回家做吧。

gorm

```log
2023/09/15 18:43:44 c:/workplace/fedi/moonchan/back/core/db/db.go:28
[error] failed to parse value &entities.Account{Id:"", Username:"", Acct:"", Url:"", DisplayName:"", Note:"", Avatar:"", AvatarStatic:"", Header:"", HeaderStatic:"", Locked:false, Fields:[]entities.Field(nil), Emojis:[]entities.CustomEmoji(nil), Bot:false, Group:false, Discoverable:(*bool)(nil), Noindex:(*bool)(nil), Moved:(*entities.Account)(nil), Suspended:false, Limited:false, CreatedAt:"", LastStatusAt:(*entities.DATEstring)(nil), StatusesCount:0, FollowersCount:0, FollowingCount:0}, got error invalid field found for struct github.com/moonchan-xyz/moonchan/back/mastodon/entities.Account's field Fields: define a valid foreign key for relations or implement the Valuer/Scanner interface
```
