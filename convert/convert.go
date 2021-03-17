package convert

import (
	"github.com/Leonardo-Antonio/lib.convert.types/helper"
)

func ToString(value interface{}) (string, error) {
	switch rsp := value.(type) {
	case string:
		return rsp, nil
	default:
		return "", helper.ErrString
	}
}

func ToBool(value interface{}) (bool, error) {
	switch rsp := value.(type) {
	case bool:
		return rsp, nil
	default:
		return false, helper.ErrBool
	}
}

func ToInt(value interface{}) (int, error) {
	switch rsp := value.(type) {
	case int:
		return rsp, nil
	default:
		return 0, helper.ErrInt
	}
}

func ToFloat32(value interface{}) (float32, error) {
	switch rsp := value.(type) {
	case float32:
		return rsp, nil
	default:
		return 0.0, helper.ErrFloat32
	}
}

func ToFloat64(value interface{}) (float64, error) {
	switch rsp := value.(type) {
	case float64:
		return rsp, nil
	default:
		return 0.0, helper.ErrFloat64
	}
}

func ToUInt(value interface{}) (uint, error) {
	switch rsp := value.(type) {
	case uint:
		return rsp, nil
	default:
		return 0, helper.ErrUInt
	}
}

func ToByte(value interface{}) (byte, error) {
	switch rsp := value.(type) {
	case byte:
		return rsp, nil
	default:
		return 0, helper.ErrByte
	}
}