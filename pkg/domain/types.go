package domain

type GoType int

const (
	GoTypeUnspecified GoType = iota
	GoTypeInt16
	GoTypeInt32
	GoTypeInt64
	GoTypeFloat32
	GoTypeFloat64
	GoTypeString
	GoTypeBool
	GoTypeTime
	GoTypeByteArray
)
