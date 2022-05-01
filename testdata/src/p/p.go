package p

type UserType struct{} // want "do not use suffix `Type` in Go"

type userType string // want "do not use suffix `Type` in Go"
