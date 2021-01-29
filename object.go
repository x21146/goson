package goson

import "encoding/json"

type JObject map[string]interface{}

func NewJObject(b []byte) (JObject, error) {
	var j JObject
	err := json.Unmarshal(b, &j)
	if err != nil {
		return nil, err
	}

	return j, nil
}

func (j JObject) GetInt(key string) (int, error) {
	v, ok := j[key]
	if !ok {
		return 0, ErrKeyNotFound
	}

	return toInt(v)
}

func (j JObject) GetIntDef(key string, def int) int {
	if v, err := j.GetInt(key); err != nil {
		return def
	} else {
		return v
	}
}

func (j JObject) GetInt64(key string) (int64, error) {
	v, ok := j[key]
	if !ok {
		return 0, ErrKeyNotFound
	}

	return toInt64(v)
}

func (j JObject) GetInt64Def(key string, def int64) int64 {
	if v, err := j.GetInt64(key); err != nil {
		return def
	} else {
		return v
	}
}

func (j JObject) GetInt32(key string) (int32, error) {
	v, ok := j[key]
	if !ok {
		return 0, ErrKeyNotFound
	}

	return toInt32(v)
}

func (j JObject) GetInt32Def(key string, def int32) int32 {
	if v, err := j.GetInt32(key); err != nil {
		return def
	} else {
		return v
	}
}

func (j JObject) GetInt16(key string) (int16, error) {
	v, ok := j[key]
	if !ok {
		return 0, ErrKeyNotFound
	}

	return toInt16(v)
}

func (j JObject) GetInt16Def(key string, def int16) int16 {
	if v, err := j.GetInt16(key); err != nil {
		return def
	} else {
		return v
	}
}

func (j JObject) GetInt8(key string) (int8, error) {
	v, ok := j[key]
	if !ok {
		return 0, ErrKeyNotFound
	}

	return toInt8(v)
}

func (j JObject) GetInt8Def(key string, def int8) int8 {
	if v, err := j.GetInt8(key); err != nil {
		return def
	} else {
		return v
	}
}

func (j JObject) GetUint64(key string) (uint64, error) {
	v, ok := j[key]
	if !ok {
		return 0, ErrKeyNotFound
	}

	return toUint64(v)
}

func (j JObject) GetUint64Def(key string, def uint64) uint64 {
	if v, err := j.GetUint64(key); err != nil {
		return def
	} else {
		return v
	}
}

func (j JObject) GetUint32(key string) (uint32, error) {
	v, ok := j[key]
	if !ok {
		return 0, ErrKeyNotFound
	}

	return toUint32(v)
}

func (j JObject) GetUint32Def(key string, def uint32) uint32 {
	if v, err := j.GetUint32(key); err != nil {
		return def
	} else {
		return v
	}
}

func (j JObject) GetUint16(key string) (uint16, error) {
	v, ok := j[key]
	if !ok {
		return 0, ErrKeyNotFound
	}

	return toUint16(v)
}

func (j JObject) GetUint16Def(key string, def uint16) uint16 {
	if v, err := j.GetUint16(key); err != nil {
		return def
	} else {
		return v
	}
}

func (j JObject) GetUint8(key string) (uint8, error) {
	v, ok := j[key]
	if !ok {
		return 0, ErrKeyNotFound
	}

	return toUint8(v)
}

func (j JObject) GetUint8Def(key string, def uint8) uint8 {
	if v, err := j.GetUint8(key); err != nil {
		return def
	} else {
		return v
	}
}

func (j JObject) GetFloat64(key string) (float64, error) {
	v, ok := j[key]
	if !ok {
		return 0, ErrKeyNotFound
	}

	return toFloat64(v)
}

func (j JObject) GetFloat64Def(key string, def float64) float64 {
	if v, err := j.GetFloat64(key); err != nil {
		return def
	} else {
		return v
	}
}

func (j JObject) GetFloat32(key string) (float32, error) {
	v, ok := j[key]
	if !ok {
		return 0, ErrKeyNotFound
	}

	return toFloat32(v)
}

func (j JObject) GetFloat32Def(key string, def float32) float32 {
	if v, err := j.GetFloat32(key); err != nil {
		return def
	} else {
		return v
	}
}

func (j JObject) GetString(key string) (string, error) {
	v, ok := j[key]
	if !ok {
		return "", ErrKeyNotFound
	}

	return toString(v)
}

func (j JObject) GetStringDef(key string, def string) string {
	if v, err := j.GetString(key); err != nil {
		return def
	} else {
		return v
	}
}

func (j JObject) GetJObject(key string) (JObject, error) {
	v, ok := j[key]
	if !ok {
		return nil, ErrKeyNotFound
	}

	return toJObject(v)
}

func (j JObject) GetJArray(key string) (JArray, error) {
	v, ok := j[key]
	if !ok {
		return nil, ErrKeyNotFound
	}

	return toJArray(v)
}