# 完全不记得了。

## 流程
现在大概是用phaser表达出以下内容
收集了
- api
- enity
通过这两种需要去搓出
- api
  - [x] go-swagger对应的注释
    - 差不多好了
  - gin对应的代码
    - ~~这里直接使用map去bind如何？~~
      - 能否通过body，formdata，json三种情况呢。
      - 其实并不知道具体的情况对应了什么的。(需要gin的test。)
      - 做不到
      - 请自己写一份
  - core/mastodon中对应的接口
- enity
  - mastodon需要的数据结构

请查看[todo.md](./TODO.md)

需要注意
- 没有websocket
- 没有oauth

重要：在做接下来的事情之前，请保证controller的部分是好的。
请将上面需要注意的部分都打上勾
