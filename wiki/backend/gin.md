# gin

是有 accpet array 的方法的，但是要求param是同一个值 例如`param[]`
考虑统统用`map[string]any`接受了试试。
自己里面是不需要考虑顺序的，所以ok的。

怎么收集value啊
- query
  - array
- header 
- body
  - array
- formdata
  - 没有array的


gin的bind不能bind map[string]any
只支持 map[string]string
遇到重复的key会覆盖

**# Q

MultipartForm怎么读啊啊啊 
A: https://gin-gonic.com/docs/examples/upload-file/multiple-file/

## 注意

顺便为了同时写好swagger和gin的变量解析
开了个从文档扒关键信息生成的repo
去那边用吧。

# log
statuses 在 `controller/mastodon/statuses.go`
接口在 `mastodon/statuses.go`

instance 的 struct 遇到了坑。好难写，[工具](https://oktools.net/json2go)

## follow_requests API methods


模板
```go
  // get authorization
	authorizationID := c.MustGet("authorizationID").(string)
	if authorizationID == "-" {
		errorJSON(c, http.StatusUnauthorized, errUnauthorized)
		return
	}

	// handle params 

  // pass params
	resp, err := mastodon.GetPreferences(id)

  // 
	if err != nil {
		errorJSON(c, http.StatusInternalServerError, newErrorMsg(err.Error()))
		return
	}
	c.JSON(http.StatusOK, resp)
```

## 
因为 upload 需要空掉没有上传的内容，所以用pointer

success 之后那个 array 可能并不对应
需要改成  {array} some.Obj

patch = update selected
post = update all ?

accept在哪里。**


