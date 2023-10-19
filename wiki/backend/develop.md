# 开发

在laptop上一共有两个project
moonchan和moonchan-wiki
分别是代码和文档
目前暂时在写后端的代码，所以code根目录设置成了moonchan/back，不然会golang开始发疯
自宅电脑就不懂了
很多没跑起来所以暂时不知道
比如db的位置之类的

这里记录设想的部分

## 目录结构

- controller
  - mastodon
  - activitypub
- db
- core
- masotodn
  - accounts
  - notifications
  - profile
  - search
  - entities
- activitypub
  - activity
- webfinger
- utils
```txt
`            ┌───────────┐
`            │           │
`            │  main.go  │
`            │           │
`            └───┬───────┘
`                │
`┌───────────┬───┴────────────────────┬───────────────────────────┬───────────────────────┬───────────┐
`│           │                        │                           │                       │           │
`│    ┌──────▼─────────────┐   ┌──────▼────────────┐    ┌─────────▼────────────┐   ┌──────▼────────┐  │
`│    │controller/webfinger│   │controller/mastodon│    │controller/activitypub│   │controller/chan│  │
`│    └────────┬───────────┘   └─────────┬─────────┘    └──────────┬───────────┘   └───────────────┘  │
`│             │                         │                         │                                  │
`└─────────────┼─────────────────────────┼─────────────────────────┼──────────────────────────────────┘
`              │                         │                         │
`     ┌────────▼──────────┐     ┌────────▼─────────┐     ┌─────────▼───────────┐
`     │                   │     │                  │     │                     │
`     │    webfinger      │     │    mastodon      │     │ activitypub/server  │
`     │                   │     │                  │     │                     │
`     └────────────┬──────┘     └─┬─────────┬──────┘     └┬─────────────┬──────┘
`                  │              │         │             │             │
`                  │              │         │             │             │
`                  │              │         │             │             │
`                  │              │         │             │             │
`                  │         ┌────▼─────────┼─────────────▼─────────────┼────────┐
`                  │         │              │                           │        │
`                  │         │   core       │            core           │        │
`                  │         │              │                           │        │
`                  │         │    + ...     │             + ...         │        ├──────────┐
`                  │         │              │                           │        │          │
`                  │         │              │                           │        │          │
`                  │         └────┬─────────┼─────────────┬─────────────┼────────┘          │
`                  │              │         │             │             │                   │
`                  │              │         │             │             │                   │
`                  │              │         │             │             │                   │
`  ┌───────────────▼──────────────▼─────────▼─────────────▼─────────────▼─────────┐   ┌─────▼──────────────────────┐
`  │                                                                              │   │                            │
`  │   db                         db                      db                      │   │                            │
`  │                                                                              │   │    activitypub/client      │
`  │    + users                    + ...                   + ...                  │   │                            │
`  │                                                                              │   │                            │
`  │                                                                              │   │                            │
`  └──────────────────────────────────────────────────────────────────────────────┘   └────────────────────────────┘
```

## webfinger
大概是好的吧？
同时有server~~和client代码~~
没有client
自己管自己的

## controller
gin解析
到gin无关的所需变量
- webfinger
- mastodon  
- activitypub
- moonchan (TODO)



## mastodon/*
```txt
http ---gin---> params --------------> functions
        ┌─────────────────────────┐    ┌────────────────┐
        │                         │    │                │
api ───►│   controller/mastodon   ├───►│   mastodon/*   │
        │                         │    │                │
        └─────────────────────────┘    └────────────────┘
```

~~@[UpdateCredentials](/back/controller/mastodon/accounts.go)~~
~~需要处理auth的header~~
~~需要处理map中的~~做好了

## db/*
挪出来了
[参见](./database.md)

## activitypub/*
关于activitypub的一切，同时有server和client

## 

### account
由于给的param太b了所以mastodon/misc.go里面加几个func处理一下吧。
话术o

# ====

写啥。。
~~想手冲忍住了继续写吧。~~
那个时候怎么性欲这么强啊，现在感觉没啥

## mastodon/*
先把这个部分写完了吧。

## core/*
有啥啊，都忘了

## ~~core/db~~
~~[参见](#db)~~

### core/activitypub

对远程的发送，返回结果

拿到obj，签名对象

不记得了

accepted。
状态转移图。我日，不知道怎么画

https://www.w3.org/TR/activitystreams-core/#intransitiveactivities
https://www.w3.org/TR/activitypub/#inbox
https://www.w3.org/TR/activitystreams-vocabulary/#dfn-accept