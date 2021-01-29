package goson

import "fmt"

type JError int32

const (
	ErrKeyNotFound JError = iota
	ErrTypeError
	ErrTypeNotSupport
	ErrIndexOutOfRange
	ErrValueOverflow
)

var errorMsg = [...]string{
	ErrKeyNotFound:     "key not found in this json.",
	ErrTypeError:       "type parse fail.",
	ErrTypeNotSupport:  "json type is not support.",
	ErrIndexOutOfRange: "index out of range.",
	ErrValueOverflow:   "value is overflow, please check type of result.",
}

func (e JError) Error() string {
	n := int32(e)
	if n < 0 || n >= int32(len(errorMsg)) {
		return fmt.Sprint("json error", n)
	}

	return errorMsg[n]
}
