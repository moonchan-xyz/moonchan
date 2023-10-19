# [accounts API methods](https://docs.joinmastodon.org/entities/Account/)

## [Register an account](https://docs.joinmastodon.org/methods/accounts/#create)
可能做好了，做了个不成样子的

## [Verify account credentials](https://docs.joinmastodon.org/methods/accounts/#verify_credentials)
好像做好了，但是怎么从`Authorization`返回id啊

# [profile API methods](https://docs.joinmastodon.org/methods/profile/) 〆

## [Delete profile avatar](https://docs.joinmastodon.org/methods/profile/#delete-profile-avatar)
```txt
DELETE /api/v1/profile/avatar HTTP/1.1
```

## [Delete profile header](https://docs.joinmastodon.org/methods/profile/#delete-profile-header)

# [notifications API methods](https://docs.joinmastodon.org/methods/notifications/)
这东西有坑。关系之类的大概要补几张表
话说级联操作怎么做啊