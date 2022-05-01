package p

type UserType struct{} // want "trim suffix `Type` from type name"

type userType string // want "trim suffix `Type` from type name"
