### Description
typenaming is **a static analyzer** that will prevent you from using the ugly "Type" suffix in type names.

It's written for educational purposes.

_Bad_
```golang
type UserType struct {}
```

_Good_
```golang
type User struct {}
```
### Install
```bash
go install github.com/armantarkhanian/typenaming@v1.1.1
```
### Usage
##### For whole directory
```bash
typenaming ./...
```
##### For certain file
```bash
typenaming main.go
```
