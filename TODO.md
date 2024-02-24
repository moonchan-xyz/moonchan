- [ ] mastodon-api-parser
  - [ ] 需要做出swaggger和gin和接口
  - 请关注那个repo中的readme的开头

# core的TODO，早了

# 流程

- [ ] 开发mastodon api
- [ ] 开发activitypub
  - [ ] api
  - [ ] action
- [ ] 开发database
  - [ ] gorm

## mastodon api
写 swagger + gin 处理接受

~~auth改个中间件出来改了没啊。~~改出来了
`mastodon/*`的api未实现
**注：流程是`controller/*`写好解析，然后再在`mastodon/*`加上框架，//todo和return nil**
**写完记得文档和`main.go`记一笔**

## 目前的进度

hint: mastodon api document 带缩进的是不是大包的。

进度看[gin.md](./wiki/backend/gin.md)


## fedi 

总之先搓出新岛

flow

- api
  1. [x] 拉api清单  
  2. [ ] 搓出go文件
     1. [ ] 从json搓出swagger注释
        1. [ ] 需要知道swagger需要啥。
        2. [ ] 把流程做好看点之后改会好些吧。
     2. [ ] 从json搓出gin接受
        1. [ ] header
        2. [ ] query
        3. [ ] 
        4. [ ] 从dataform上传的文件
     3. [ ] 从json搓出golang接受了返回的东西
        1. [ ] 返回enities.Entity, error差不多吧
        2. [ ] 框架先返回nil, nil
        3. [ ] 接受的param和其**类型**
- enity
  1. [x] 拉enities清单
  2. [ ] 搓出go文件
     1. [ ] many to many的关系？
        1. [ ] 新建一个测试一下。
     2. [ ] 



    - [ ] 


目前在mastodon-api-phaser的repo里面工作。
做了些啥啊。。。

## readme

因为分支可以直接放pages上所以开个新分支放在`/`了

- [x] swagger
  - [ ] 但是还没测试,如果测到了再有问题再来订正
- [ ] entities
  - [ ] 导出来能用了没
  - [ ] 几个小的subclass好像还是不行的
- [ ] controller
  - [ ] 要