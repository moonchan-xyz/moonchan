# golnag

json 可以 marshal 到 struct
那么自己出个data struct，该怎么从 gin.Form marshal过来。
或者从[PostForm](gin.md#applicationx-www-form-urlencoded) marshal 过来
总之，key的形式是 `key[key][key]` 这种的。

首先去参考json怎么做的。
用了`reflect.Value`，然后看不懂了
好像是做了scanner然后赋值之类的，但是具体没看懂。
扫出一段然后找回去赋值大概是这样。
但是这里要做的事情有点不一样，需要扫

遍历自己的keys然后得到tag去找value。
如何处理嵌套的 struct 和 array

原来struct后面的tag也是struct的一部分。属于type吧？

## [reflect](./wiki/reflect.md)

## formData

因为给来的key是literial的。

所以大概不能用json的那个逻辑吧(虽然也不会用)


