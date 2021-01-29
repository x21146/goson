package goson

import "encoding/json"

type JArray []interface{}

func NewJArray(b []byte) (JArray, error) {
	var j JArray
	err := json.Unmarshal(b, &j)
	if err != nil {
		return nil, err
	}

	return j, nil
}

func (a JArray) checkIndex(i int) bool {
	return i >= 0 && i < len(a)
}

func (a JArray) ToIntSlice() ([]int, error) {
	s := make([]int, len(a))
	for i, v := range a {
		var err error
		s[i], err = toInt(v)
		if err != nil {
			return nil, err
		}
	}

	return s, nil
}

func (a JArray) ToInt64Slice() ([]int64, error) {
	s := make([]int64, len(a))
	for i, v := range a {
		var err error
		s[i], err = toInt64(v)
		if err != nil {
			return nil, err
		}
	}

	return s, nil
}

func (a JArray) ToInt32Slice() ([]int32, error) {
	s := make([]int32, len(a))
	for i, v := range a {
		var err error
		s[i], err = toInt32(v)
		if err != nil {
			return nil, err
		}
	}

	return s, nil
}

func (a JArray) ToInt16Slice() ([]int16, error) {
	s := make([]int16, len(a))
	for i, v := range a {
		var err error
		s[i], err = toInt16(v)
		if err != nil {
			return nil, err
		}
	}

	return s, nil
}

func (a JArray) ToInt8Slice() ([]int8, error) {
	s := make([]int8, len(a))
	for i, v := range a {
		var err error
		s[i], err = toInt8(v)
		if err != nil {
			return nil, err
		}
	}

	return s, nil
}

func (a JArray) ToUint64Slice() ([]uint64, error) {
	s := make([]uint64, len(a))
	for i, v := range a {
		var err error
		s[i], err = toUint64(v)
		if err != nil {
			return nil, err
		}
	}

	return s, nil
}

func (a JArray) ToUint32Slice() ([]uint32, error) {
	s := make([]uint32, len(a))
	for i, v := range a {
		var err error
		s[i], err = toUint32(v)
		if err != nil {
			return nil, err
		}
	}

	return s, nil
}

func (a JArray) ToUint16Slice() ([]uint16, error) {
	s := make([]uint16, len(a))
	for i, v := range a {
		var err error
		s[i], err = toUint16(v)
		if err != nil {
			return nil, err
		}
	}

	return s, nil
}

func (a JArray) ToUint8Slice() ([]uint8, error) {
	s := make([]uint8, len(a))
	for i, v := range a {
		var err error
		s[i], err = toUint8(v)
		if err != nil {
			return nil, err
		}
	}

	return s, nil
}

func (a JArray) ToFloat64Slice() ([]float64, error) {
	s := make([]float64, len(a))
	for i, v := range a {
		var err error
		s[i], err = toFloat64(v)
		if err != nil {
			return nil, err
		}
	}

	return s, nil
}

func (a JArray) ToFloat32Slice() ([]float32, error) {
	s := make([]float32, len(a))
	for i, v := range a {
		var err error
		s[i], err = toFloat32(v)
		if err != nil {
			return nil, err
		}
	}

	return s, nil
}

func (a JArray) ToStringSlice() ([]string, error) {
	s := make([]string, len(a))
	for i, v := range a {
		var err error
		s[i], err = toString(v)
		if err != nil {
			return nil, err
		}
	}

	return s, nil
}

func (a JArray) ToJObjectSlice() ([]JObject, error) {
	s := make([]JObject, len(a))
	for i, v := range a {
		var err error
		s[i], err = toJObject(v)
		if err != nil {
			return nil, err
		}
	}

	return s, nil
}

func (a JArray) ToJArraySlice() ([]JArray, error) {
	s := make([]JArray, len(a))
	for i, v := range a {
		var err error
		s[i], err = toJArray(v)
		if err != nil {
			return nil, err
		}
	}

	return s, nil
}

func (a JArray) GetInt(i int) (int, error) {
	if !a.checkIndex(i) {
		return 0, ErrIndexOutOfRange
	}

	return toInt(a[i])
}

func (a JArray) GetInt64(i int) (int64, error) {
	if !a.checkIndex(i) {
		return 0, ErrIndexOutOfRange
	}

	return toInt64(a[i])
}

func (a JArray) GetInt32(i int) (int32, error) {
	if !a.checkIndex(i) {
		return 0, ErrIndexOutOfRange
	}

	return toInt32(a[i])
}

func (a JArray) GetInt16(i int) (int16, error) {
	if !a.checkIndex(i) {
		return 0, ErrIndexOutOfRange
	}

	return toInt16(a[i])
}

func (a JArray) GetInt8(i int) (int8, error) {
	if !a.checkIndex(i) {
		return 0, ErrIndexOutOfRange
	}

	return toInt8(a[i])
}

func (a JArray) GetUint64(i int) (uint64, error) {
	if !a.checkIndex(i) {
		return 0, ErrIndexOutOfRange
	}

	return toUint64(a[i])
}

func (a JArray) GetUint32(i int) (uint32, error) {
	if !a.checkIndex(i) {
		return 0, ErrIndexOutOfRange
	}

	return toUint32(a[i])
}

func (a JArray) GetUint16(i int) (uint16, error) {
	if !a.checkIndex(i) {
		return 0, ErrIndexOutOfRange
	}

	return toUint16(a[i])
}

func (a JArray) GetUint8(i int) (uint8, error) {
	if !a.checkIndex(i) {
		return 0, ErrIndexOutOfRange
	}

	return toUint8(a[i])
}

func (a JArray) GetFloat64(i int) (float64, error) {
	if !a.checkIndex(i) {
		return 0, ErrIndexOutOfRange
	}

	return toFloat64(a[i])
}

func (a JArray) GetFloat32(i int) (float32, error) {
	if !a.checkIndex(i) {
		return 0, ErrIndexOutOfRange
	}

	return toFloat32(a[i])
}

func (a JArray) GetString(i int) (string, error) {
	if !a.checkIndex(i) {
		return "", ErrIndexOutOfRange
	}

	return toString(a[i])
}

func (a JArray) GetJObject(i int) (JObject, error) {
	if !a.checkIndex(i) {
		return nil, ErrIndexOutOfRange
	}

	return toJObject(a[i])
}

func (a JArray) GetJArray(i int) (JArray, error) {
	if !a.checkIndex(i) {
		return nil, ErrIndexOutOfRange
	}

	return toJArray(a[i])
}