**# Q

MultipartForm怎么读啊啊啊 
A: https://gin-gonic.com/docs/examples/upload-file/multiple-file/

## 注意

顺便为了同时写好swagger和gin的变量解析
开了个从文档扒关键信息生成的repo
去那边用吧。


# api

这里的进度是gin -> params的进度
mastodon的进度去mastodon里面看

- [x] **apps**
  - 这又是啥
- [x] oauth
  - 不会
- [x] emails
  - 这啥？
- [x] **accounts**
  - [x] Register an account
    - [ ] token 回了个啥啊
    - [ ] key 要确认一下
  - [x] Verify account credentials
    - [x] auth的确认
  - [x] Update account credentials
    - [ ] todo 上传图片
  - [x] Get account
  - [x] Get account’s statuses
    - [ ] error没做
  - [x] Get account’s followers
  - [x] Get account’s following
  - [ ] Get account’s featured tags
    - 跳过，看上去像推荐
  - [ ] Get lists containing this account
    - 跳过
  - [x] Follow account
  - [x] Unfollow account
  - [x] Remove account from followers
  - [x] Block account
  - [x] Unblock account
  - [x] Mute account
  - [x] Unmute account
  - [ ] Feature account on your profile
    - 跳过，看上去像推荐
  - [ ] Unfeature account from profile
    - 跳过，看上去像推荐
  - [x] Set private note on profile
    - 就是备注
  - [x] Check relationships to other accounts
  - [ ] Find familiar followers
    - 跳过
  - [ ] Search for matching accounts
    - 跳过? 但是是搜索
  - [x] Lookup account ID from Webfinger address
  - [ ] (DEPRECATED) Identity proofs
- [x] bookmarks
- [x] favourites
- [x] mutes
- [x] blocks
- [x] domain_blocks
- filters
- [x] reports
- [x] follow_requests
  - [x] View pending follow requests
  - [x] Accept follow request
  - [x] Reject follow request
- [x] endorsements
- featured_tags
- [x] preferences
  - [x] View user preferences
- followed_tags
- suggestions
- tags
- [x] **profile**
  - [x] Delete profile avatar
  - [x] Delete profile header
- **statuses**
  - [ ] Post a new status
  - [ ] View a single status
  - [ ] Delete a status
  - [ ] Get parent and child statuses in context
  - [ ] Translate a status
  - [ ] See who boosted a status
  - [ ] See who favourited a status
  - [ ] Favourite a status
  - [ ] Undo favourite of a status
  - [ ] Boost a status
  - [ ] Undo boost of a status
  - [ ] Bookmark a status
  - [ ] Undo bookmark of a status
  - [ ] Mute a conversation
  - [ ] Unmute a conversation
  - [ ] Pin status to profile
  - [ ] Unpin status from profile
  - [ ] Edit a status
  - [ ] View edit history of a status
  - [ ] View status source
  - [ ] (DEPRECATED) Fetch preview card
- media
- polls
- scheduled_statuses
- **timelines**
- conversations
- lists
- markers
- streaming
- **notifications**
  - [ ] Get all notifications
  - [ ] Get a single notification
  - [ ] Dismiss all notifications
  - [ ] Dismiss a single notification
  - [ ] (REMOVED) Dismiss a single notification
- push
- **search**
- **instance**
 - [ ] View server information
 - [ ] List of connected domains
 - [ ] Weekly activity
 - [ ] List of rules
 - [ ] View moderated servers
 - [ ] View extended description
 - [ ] (DEPRECATED) View server information (V1)
- trends
- directory
- custom_emojis
- [x] announcements
- **admin**
- admin/accounts
- admin/domain_blocks
- admin/reports
- admin/trends
- canonical_email_blocks
- dimensions
- domain_allows
- email_domain_blocks
- ip_blocks
- measures
- retention
- [x] **proofs**
- [x] **oembed**


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


