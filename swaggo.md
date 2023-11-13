# swaggo

## 初始化

以后补。


## go-swagger

### 注释

如何生成注释
注视的格式，必要内容。


swagger的接入参数和go-swagger的参数好像差得还蛮远的.
最好还是有统一的格式, 大概是ymal, 能生成golang的comment

## swag

```go
//  godoc
// @Summary
// @Description
// @Tags mastodon/accounts
// @Produce 
// @Param   resource query string true "search webfinger object by account"
// @Success 200 {object} object{subject=string,links=[]any}
// @Failure 404 {object} object{error=string}
// @Failure 422 {object} object{error=string}
// @Failure 500 {object} object{error=string}
// @Router /.well-known/webfinger [get]
```

- 标题 godoc
- @Summary 会在标题旁边
- @Description 展开之后显示
- @Description 多行
- @Tags 在哪些 group 当中
- @Produce MIME类型
- @Param 参数名 参数类型 数据类型 is强制的? "注释内容" attribute(optional)
- @Success 200 {object} object{subject=string,links=[]any}
- @Failure 404 {object} object{error=string}
- @Failure 422 {object} object{error=string}
- @Failure 500 {object} object{error=string}
- @Router 从`/`开始的 path [get]

### MIME types

|                       |                                   |
| --------------------- | --------------------------------- |
| json                  | application/json                  |
| xml                   | text/xml                          |
| plain                 | text/plain                        |
| html                  | text/html                         |
| mpfd                  | multipart/form-data               |
| x-www-form-urlencoded | application/x-www-form-urlencoded |
| json-api              | application/vnd.api+json          |
| json-stream           | application/x-json-stream         |
| octet-stream          | application/octet-stream          |
| png                   | image/png                         |
| jpeg                  | image/jpeg                        |
| gif                   | image/gif                         |

### param types
- query
- path
- header
- body
- formData

### data types
- string (string)
- integer (int, uint, uint32, uint64)
- number (float32)
- boolean (bool)
- user defined struct

# TODO

~~配置好 swag，让网页能显示 api~~
检查下 webfinger？也能直接捅 db <- 为啥？
写 mastodon 的 api，gin 转换到变量 <- 还差file转成url
写 mastodon 的 api 的处理，可以直接捅 db，其他的交给 core
也不知道能不能编译

写 swagger + gin 处理接受
