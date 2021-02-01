package gson

import (
	"encoding/json"
	"math"
	"strconv"
)

// parse json value to int
func toInt(i interface{}) (int, error) {
	// all value int json map will become float64
	switch i.(type) {
	case float64:
		v := i.(float64)
		if v > math.MaxInt64 {
			return 0, ErrValueOverflow
		}

		rv := int(v)
		if float64(rv) != v {
			return 0, ErrTypeError
		}
		return rv, nil
	case string:
		v, err := strconv.Atoi(i.(string))
		if err != nil {
			return 0, ErrTypeError
		}

		return v, nil
	default:
		return 0, ErrTypeNotSupport
	}
}

// parse json value to int64
func toInt64(i interface{}) (int64, error) {
	// all value int json map will become float64
	switch i.(type) {
	case float64:
		v := i.(float64)
		if v > math.MaxInt64 {
			return 0, ErrValueOverflow
		}

		rv := int64(v)
		if float64(rv) != v {
			return 0, ErrTypeError
		}
		return rv, nil
	case string:
		v, err := strconv.ParseInt(i.(string), 10, 64)
		if err != nil {
			return 0, ErrTypeError
		}

		return v, nil
	default:
		return 0, ErrTypeNotSupport
	}
}

// parse json value to int32
func toInt32(i interface{}) (int32, error) {
	// all value int json map will become float64
	switch i.(type) {
	case float64:
		v := i.(float64)
		if v > math.MaxInt32 {
			return 0, ErrValueOverflow
		}

		rv := int32(v)
		if float64(rv) != v {
			return 0, ErrTypeError
		}
		return rv, nil
	case string:
		v, err := strconv.ParseInt(i.(string), 10, 32)
		if err != nil {
			return 0, ErrTypeError
		}

		return int32(v), nil
	default:
		return 0, ErrTypeNotSupport
	}
}

// parse json value to int16
func toInt16(i interface{}) (int16, error) {
	// all value int json map will become float64
	switch i.(type) {
	case float64:
		v := i.(float64)
		if v > math.MaxInt16 {
			return 0, ErrValueOverflow
		}

		rv := int16(v)
		if float64(rv) != v {
			return 0, ErrTypeError
		}
		return rv, nil
	case string:
		v, err := strconv.ParseInt(i.(string), 10, 16)
		if err != nil {
			return 0, ErrTypeError
		}

		return int16(v), nil
	default:
		return 0, ErrTypeNotSupport
	}
}

// parse json value to int8
func toInt8(i interface{}) (int8, error) {
	// all value int json map will become float64
	switch i.(type) {
	case float64:
		v := i.(float64)
		if v > math.MaxInt8 {
			return 0, ErrValueOverflow
		}

		rv := int8(v)
		if float64(rv) != v {
			return 0, ErrTypeError
		}
		return rv, nil
	case string:
		v, err := strconv.ParseInt(i.(string), 10, 8)
		if err != nil {
			return 0, ErrTypeError
		}

		return int8(v), nil
	default:
		return 0, ErrTypeNotSupport
	}
}

// parse json value to uint64
func toUint64(i interface{}) (uint64, error) {
	// all value int json map will become float64
	switch i.(type) {
	case float64:
		v := i.(float64)
		if v > math.MaxUint64 {
			return 0, ErrValueOverflow
		}

		rv := uint64(v)
		if float64(rv) != v {
			return 0, ErrTypeError
		}
		return rv, nil
	case string:
		v, err := strconv.ParseUint(i.(string), 10, 64)
		if err != nil {
			return 0, ErrTypeError
		}

		return v, nil
	default:
		return 0, ErrTypeNotSupport
	}
}

// parse json value to uint32
func toUint32(i interface{}) (uint32, error) {
	// all value int json map will become float64
	switch i.(type) {
	case float64:
		v := i.(float64)
		if v > math.MaxUint32 {
			return 0, ErrValueOverflow
		}

		rv := uint32(v)
		if float64(rv) != v {
			return 0, ErrTypeError
		}
		return rv, nil
	case string:
		v, err := strconv.ParseUint(i.(string), 10, 32)
		if err != nil {
			return 0, ErrTypeError
		}

		return uint32(v), nil
	default:
		return 0, ErrTypeNotSupport
	}
}

// parse json value to int16
func toUint16(i interface{}) (uint16, error) {
	// all value int json map will become float64
	switch i.(type) {
	case float64:
		v := i.(float64)
		if v > math.MaxUint16 {
			return 0, ErrValueOverflow
		}

		rv := uint16(v)
		if float64(rv) != v {
			return 0, ErrTypeError
		}
		return rv, nil
	case string:
		v, err := strconv.ParseUint(i.(string), 10, 16)
		if err != nil {
			return 0, ErrTypeError
		}

		return uint16(v), nil
	default:
		return 0, ErrTypeNotSupport
	}
}

// parse json value to int8
func toUint8(i interface{}) (uint8, error) {
	// all value int json map will become float64
	switch i.(type) {
	case float64:
		v := i.(float64)
		if v > math.MaxUint8 {
			return 0, ErrValueOverflow
		}

		rv := uint8(v)
		if float64(rv) != v {
			return 0, ErrTypeError
		}
		return rv, nil
	case string:
		v, err := strconv.ParseUint(i.(string), 10, 8)
		if err != nil {
			return 0, ErrTypeError
		}

		return uint8(v), nil
	default:
		return 0, ErrTypeNotSupport
	}
}

// parse json value to float32
func toFloat32(i interface{}) (float32, error) {
	// all value int json map will become float64
	switch i.(type) {
	case float64:
		v := i.(float64)

		rv := float32(v)
		// float32 to float 64 may cause accuracy problem to let it not equal,
		// so float32 parser not check value
		//if float64(rv) != v {
		//	return 0, ErrTypeError
		//}
		return rv, nil
	case string:
		v, err := strconv.ParseFloat(i.(string), 32)
		if err != nil {
			return 0, ErrTypeError
		}

		return float32(v), nil
	default:
		return 0, ErrTypeNotSupport
	}
}

// parse json value to float64
func toFloat64(i interface{}) (float64, error) {
	// all value int json map will become float64
	switch i.(type) {
	case float64:
		v := i.(float64)
		return v, nil
	case string:
		v, err := strconv.ParseFloat(i.(string), 64)
		if err != nil {
			return 0, ErrTypeError
		}

		return v, nil
	default:
		return 0, ErrTypeNotSupport
	}
}

// parse json value to string, also parse json object and json array to string
func toString(i interface{}) (string, error) {
	switch i.(type) {
	case float64:
		s := strconv.FormatFloat(i.(float64), 'f', -1, 64)
		return s, nil
	case string:
		return i.(string), nil
	case map[string]interface{}, []interface{}:
		b, _ := json.Marshal(i)
		return string(b), nil
	default:
		return "", ErrTypeNotSupport
	}
}

func toJObject(i interface{}) (JObject, error) {
	switch i.(type) {
	case string:
		j := JObject{}
		err := json.Unmarshal([]byte(i.(string)), &j)
		if err != nil {
			return nil, ErrTypeError
		}

		return j, nil
	case map[string]interface{}:
		return i.(map[string]interface{}), nil
	default:
		return nil, ErrTypeNotSupport
	}
}

func toJArray(i interface{}) (JArray, error) {
	switch i.(type) {
	case string:
		j := JArray{}
		err := json.Unmarshal([]byte(i.(string)), &j)
		if err != nil {
			return nil, ErrTypeError
		}

		return j, nil
	case []interface{}:
		return i.([]interface{}), nil
	default:
		return nil, ErrTypeNotSupport
	}
}