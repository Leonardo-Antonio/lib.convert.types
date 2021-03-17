package helper

import "errors"

var (
	ErrString = errors.New("value is not a string type")
	ErrBool = errors.New("value is not a bool type")
	ErrInt = errors.New("value is not a int type")
	ErrUInt = errors.New("value is not a uint type")
	ErrFloat32 = errors.New("value is not a float32 type")
	ErrFloat64 = errors.New("value is not a float64 type")
	ErrByte = errors.New("value is not a byte type")
)

