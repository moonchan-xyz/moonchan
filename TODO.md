# TODO

配置好swag，让网页能显示api
检查下webfinger？也能直接捅db
写mastodon的api，gin转换到变量
写mastodon的api的测试，用curl，测试写这里就好？
写mastodon的api的处理，可以直接捅db，其他的交给core
也不知道能不能编译
httpsig挪过来了，记得
db可以不用加字段而是加外键来着
外键怎么优雅地加呃呃呃，级联操作怎么来的哇。

大概很久后才会打开。

写 swagger + gin 处理接受

~~auth改个中间件出来改了没啊。~~改出来了
`mastodon/*`的api未实现
*注：流程是`controller/*`写好解析，然后再在`mastodon/*`加上框架，//todo和return nil*
*写完记得文档和`main.go`记一笔*

hint: mastodon api document 带缩进的是不是大包的。

进度看[gin.md](./wiki/backend/gin.md)
