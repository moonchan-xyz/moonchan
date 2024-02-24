
- [gin](#gin)
  - [accept params](#accept-params)
    - [header](#header)
    - [query](#query)
    - [path](#path)
    - [formData](#formdata)
      - [application/x-www-form-urlencoded](#applicationx-www-form-urlencoded)
      - [multipart/form-data](#multipartform-data)
 
# gin

记录一下怎么从别处弄过来

## accept params

how to get params from different inputs 

### header
- GetHeader
- Cookie
```go
authorization := c.GetHeader("Authorization") 
```
string

### query
- Query
- DefaultQuery
- QueryArray
```go
maxID := c.Query("max_id")
```
string

### path
- Param
```go
id := c.Param("id")
```
string

### formData

#### application/x-www-form-urlencoded
- PostForm
- DefaultPostForm
- PostFormArray
```go

```

#### multipart/form-data 

```go
// Form is a parsed multipart form.
// Its File parts are stored either in memory or on disk,
// and are accessible via the *FileHeader's Open method.
// Its Value parts are stored as strings.
// Both are keyed by field name.
type Form struct {
	Value map[string][]string
	File  map[string][]*FileHeader
}
```
```go
form, err := c.MultipartForm() // todo should check type
if err != nil {
  errorJSON(c, http.StatusBadRequest, newErrorMsg(err.Error()))
}
fe := formExactor{Form: form}
```
