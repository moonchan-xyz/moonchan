# ref

[post json](https://github.com/go-swagger/go-swagger/issues/1772)

[example](https://github.com/swaggo/swag/blob/e769bbe213c871a10c6fc64f64dfce47cddce40f/testdata/simple2/api/api.go)

# 流程

https://goswagger.io/about.html

[follow](https://github.com/swaggo/gin-swagger)

创建了`doc/`文件夹，然后呢

[follow](https://github.com/swaggo/gin-swagger#canonical-example)

https://github.com/swaggo/swag/blob/master/README_zh-CN.md#%E5%A6%82%E4%BD%95%E4%B8%8Egin%E9%9B%86%E6%88%90

扛不住了

大概是main.go没加必要字段
照着文档加一下。下次一定

POST 表格样式要加
`@Accept		multipart/form-data`

gin swagger
接受图片的param怎么写

https://goswagger.io/use/models/schemas.html
难绷，用来干啥的都不知道


http 
query 传参
post 传参
multypart 传参
三种格式

# ref
- 标题 godoc
- @Summary 会在标题旁边
- @Description 展开之后显示
- @Description 多行
- @Tags 在哪些 group 当中
- @Produce MIME类型
- @Param 参数名 参数类型 数据类型 is强制的? "注释内容" attribute(optional)
  - 参数类型
    - query url之后的
    - path {}path中解析的
    - header header中的
    - body body中的(eg: a=1&b=2)
    - formData multiple/form-data
  - 数据类型
    - string 
    - integer
    - number (float)
    - boolean
    - file 文件 这个为啥document没写
- @Success 200 {object} object{subject=string,links=[]any}
  - 这是个自己编object的例子
  - 可以用go代码里的struct
- @Failure 404 {object} object{error=string}
- @Failure 422 {object} object{error=string}
- @Failure 500 {object} object{error=string}
- @Router 从`/`开始的 path [get]


# Q

怎么在MultipartForm上传图片啊
A: Param 第三个参数选file