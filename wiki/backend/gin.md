# Q

MultipartForm怎么读啊啊啊 
A: https://gin-gonic.com/docs/examples/upload-file/multiple-file/

# overall
- [x] Register an account
  - [ ] token 回了个啥啊
  - [ ] key 要确认一下
- [x] Verify account credentials
  - [ ] auth的确认
- [x] Update account credentials
  - [ ] todo 上传图片
- [x] Get account
- [x] Get account’s statuses
  - [ ] error没做
- [x] Get account’s followers
- [x] Get account’s following
- [ ] Get account’s featured tags
  - 跳过
- [ ] Get lists containing this account
  - 跳过
- [ ] Follow account
- [ ] Unfollow account
- [ ] Remove account from followers
- [ ] Block account
- [ ] Unblock account
- [ ] Mute account
- [ ] Unmute account
- [ ] Feature account on your profile
  - 跳过
- [ ] Unfeature account from profile
  - 跳过
- [x] Set private note on profile
  - 就是备注
- [x] Check relationships to other accounts
- [ ] Find familiar followers
  - 跳过
- [ ] Search for matching accounts
  - 跳过? 但是是搜索
- [x] Lookup account ID from Webfinger address
- [ ] (DEPRECATED) Identity proofs

## 
因为 upload 需要空掉没有上传的内容，所以用pointer

success 之后那个 array 可能并不对应
需要改成  {array} some.Obj

patch = update selected
post = update all ?

accept在哪里。