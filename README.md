### Description
typenaming is **a static analyzer** that will prevent you from using the "Type" suffix in type names.
### Example
_Bad_
```golang
type UserType struct {}
```

_Good_
```golang
type User struct {}
```
