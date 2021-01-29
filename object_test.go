package goson

import (
	"reflect"
	"testing"
)

func mockJObject(str string) JObject {
	j, _ := NewJObject([]byte(str))
	return j
}

func TestJObject_GetFloat32(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		j       JObject
		args    args
		want    float32
		wantErr bool
	}{
		{"string float value", mockJObject(`{"v": "12.345"}`), args{"v"}, 12.345, false},
		{"string int value", mockJObject(`{"v": "12"}`), args{"v"}, 12.0, false},
		{"string value", mockJObject(`{"v": "string"}`), args{"v"}, 0, true},
		{"float value", mockJObject(`{"v": 12.345}`), args{"v"}, 12.345, false},
		{"int value", mockJObject(`{"v": 12}`), args{"v"}, 12.0, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.j.GetFloat32(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFloat32() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetFloat32() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJObject_GetFloat32Def(t *testing.T) {
	type args struct {
		key string
		def float32
	}
	tests := []struct {
		name string
		j    JObject
		args args
		want float32
	}{
		{"string float value", mockJObject(`{"v": "12.345"}`), args{"v", -1}, 12.345},
		{"string int value", mockJObject(`{"v": "12"}`), args{"v", -1}, 12.0},
		{"string value", mockJObject(`{"v": "string"}`), args{"v", -1}, -1},
		{"float value", mockJObject(`{"v": 12.345}`), args{"v", -1}, 12.345},
		{"int value", mockJObject(`{"v": 12}`), args{"v", -1}, 12.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.j.GetFloat32Def(tt.args.key, tt.args.def); got != tt.want {
				t.Errorf("GetFloat32Def() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJObject_GetFloat64(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		j       JObject
		args    args
		want    float64
		wantErr bool
	}{
		{"string float value", mockJObject(`{"v": "12.345"}`), args{"v"}, 12.345, false},
		{"string int value", mockJObject(`{"v": "12"}`), args{"v"}, 12.0, false},
		{"string value", mockJObject(`{"v": "string"}`), args{"v"}, 0, true},
		{"float value", mockJObject(`{"v": 12.345}`), args{"v"}, 12.345, false},
		{"int value", mockJObject(`{"v": 12}`), args{"v"}, 12.0, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.j.GetFloat64(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFloat64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetFloat64() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJObject_GetFloat64Def(t *testing.T) {
	type args struct {
		key string
		def float64
	}
	tests := []struct {
		name string
		j    JObject
		args args
		want float64
	}{
		{"string float value", mockJObject(`{"v": "12.345"}`), args{"v", -1}, 12.345},
		{"string int value", mockJObject(`{"v": "12"}`), args{"v", -1}, 12.0},
		{"string value", mockJObject(`{"v": "string"}`), args{"v", -1}, -1},
		{"float value", mockJObject(`{"v": 12.345}`), args{"v", -1}, 12.345},
		{"int value", mockJObject(`{"v": 12}`), args{"v", -1}, 12.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.j.GetFloat64Def(tt.args.key, tt.args.def); got != tt.want {
				t.Errorf("GetFloat64Def() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJObject_GetInt(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		j       JObject
		args    args
		want    int
		wantErr bool
	}{
		{"string float value", mockJObject(`{"v": "12.345"}`), args{"v"}, 0, true},
		{"string int value", mockJObject(`{"v": "12"}`), args{"v"}, 12, false},
		{"string value", mockJObject(`{"v": "string"}`), args{"v"}, 0, true},
		{"float value", mockJObject(`{"v": 12.345}`), args{"v"}, 0, true},
		{"int value", mockJObject(`{"v": 12}`), args{"v"}, 12, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.j.GetInt(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetInt() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJObject_GetInt16(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		j       JObject
		args    args
		want    int16
		wantErr bool
	}{
		{"string float value", mockJObject(`{"v": "12.345"}`), args{"v"}, 0, true},
		{"string int value", mockJObject(`{"v": "12"}`), args{"v"}, 12, false},
		{"string value", mockJObject(`{"v": "string"}`), args{"v"}, 0, true},
		{"float value", mockJObject(`{"v": 12.345}`), args{"v"}, 0, true},
		{"int value", mockJObject(`{"v": 12}`), args{"v"}, 12, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.j.GetInt16(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetInt16() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetInt16() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJObject_GetInt16Def(t *testing.T) {
	type args struct {
		key string
		def int16
	}
	tests := []struct {
		name string
		j    JObject
		args args
		want int16
	}{
		{"string float value", mockJObject(`{"v": "12.345"}`), args{"v", -1}, -1},
		{"string int value", mockJObject(`{"v": "12"}`), args{"v", -1}, 12},
		{"string value", mockJObject(`{"v": "string"}`), args{"v", -1}, -1},
		{"float value", mockJObject(`{"v": 12.345}`), args{"v", -1}, -1},
		{"int value", mockJObject(`{"v": 12}`), args{"v", -1}, 12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.j.GetInt16Def(tt.args.key, tt.args.def); got != tt.want {
				t.Errorf("GetInt16Def() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJObject_GetInt32(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		j       JObject
		args    args
		want    int32
		wantErr bool
	}{
		{"string float value", mockJObject(`{"v": "12.345"}`), args{"v"}, 0, true},
		{"string int value", mockJObject(`{"v": "12"}`), args{"v"}, 12, false},
		{"string value", mockJObject(`{"v": "string"}`), args{"v"}, 0, true},
		{"float value", mockJObject(`{"v": 12.345}`), args{"v"}, 0, true},
		{"int value", mockJObject(`{"v": 12}`), args{"v"}, 12, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.j.GetInt32(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetInt32() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetInt32() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJObject_GetInt32Def(t *testing.T) {
	type args struct {
		key string
		def int32
	}
	tests := []struct {
		name string
		j    JObject
		args args
		want int32
	}{
		{"string float value", mockJObject(`{"v": "12.345"}`), args{"v", -1}, -1},
		{"string int value", mockJObject(`{"v": "12"}`), args{"v", -1}, 12},
		{"string value", mockJObject(`{"v": "string"}`), args{"v", -1}, -1},
		{"float value", mockJObject(`{"v": 12.345}`), args{"v", -1}, -1},
		{"int value", mockJObject(`{"v": 12}`), args{"v", -1}, 12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.j.GetInt32Def(tt.args.key, tt.args.def); got != tt.want {
				t.Errorf("GetInt32Def() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJObject_GetInt64(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		j       JObject
		args    args
		want    int64
		wantErr bool
	}{
		{"string float value", mockJObject(`{"v": "12.345"}`), args{"v"}, 0, true},
		{"string int value", mockJObject(`{"v": "12"}`), args{"v"}, 12, false},
		{"string value", mockJObject(`{"v": "string"}`), args{"v"}, 0, true},
		{"float value", mockJObject(`{"v": 12.345}`), args{"v"}, 0, true},
		{"int value", mockJObject(`{"v": 12}`), args{"v"}, 12, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.j.GetInt64(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetInt64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetInt64() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJObject_GetInt64Def(t *testing.T) {
	type args struct {
		key string
		def int64
	}
	tests := []struct {
		name string
		j    JObject
		args args
		want int64
	}{
		{"string float value", mockJObject(`{"v": "12.345"}`), args{"v", -1}, -1},
		{"string int value", mockJObject(`{"v": "12"}`), args{"v", -1}, 12},
		{"string value", mockJObject(`{"v": "string"}`), args{"v", -1}, -1},
		{"float value", mockJObject(`{"v": 12.345}`), args{"v", -1}, -1},
		{"int value", mockJObject(`{"v": 12}`), args{"v", -1}, 12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.j.GetInt64Def(tt.args.key, tt.args.def); got != tt.want {
				t.Errorf("GetInt64Def() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJObject_GetInt8(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		j       JObject
		args    args
		want    int8
		wantErr bool
	}{
		{"string float value", mockJObject(`{"v": "12.345"}`), args{"v"}, 0, true},
		{"string int value", mockJObject(`{"v": "12"}`), args{"v"}, 12, false},
		{"string value", mockJObject(`{"v": "string"}`), args{"v"}, 0, true},
		{"float value", mockJObject(`{"v": 12.345}`), args{"v"}, 0, true},
		{"int value", mockJObject(`{"v": 12}`), args{"v"}, 12, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.j.GetInt8(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetInt8() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetInt8() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJObject_GetInt8Def(t *testing.T) {
	type args struct {
		key string
		def int8
	}
	tests := []struct {
		name string
		j    JObject
		args args
		want int8
	}{
		{"string float value", mockJObject(`{"v": "12.345"}`), args{"v", -1}, -1},
		{"string int value", mockJObject(`{"v": "12"}`), args{"v", -1}, 12},
		{"string value", mockJObject(`{"v": "string"}`), args{"v", -1}, -1},
		{"float value", mockJObject(`{"v": 12.345}`), args{"v", -1}, -1},
		{"int value", mockJObject(`{"v": 12}`), args{"v", -1}, 12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.j.GetInt8Def(tt.args.key, tt.args.def); got != tt.want {
				t.Errorf("GetInt8Def() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJObject_GetIntDef(t *testing.T) {
	type args struct {
		key string
		def int
	}
	tests := []struct {
		name string
		j    JObject
		args args
		want int
	}{
		{"string float value", mockJObject(`{"v": "12.345"}`), args{"v", -1}, -1},
		{"string int value", mockJObject(`{"v": "12"}`), args{"v", -1}, 12.0},
		{"string value", mockJObject(`{"v": "string"}`), args{"v", -1}, -1},
		{"float value", mockJObject(`{"v": 12.345}`), args{"v", -1}, -1},
		{"int value", mockJObject(`{"v": 12}`), args{"v", -1}, 12.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.j.GetIntDef(tt.args.key, tt.args.def); got != tt.want {
				t.Errorf("GetIntDef() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJObject_GetJArray(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		j       JObject
		args    args
		want    JArray
		wantErr bool
	}{
		{"int value", mockJObject(`{"v": 123}`), args{"v"}, nil, true},
		{"float value", mockJObject(`{"v": 1.23}`), args{"v"}, nil, true},
		{"string value", mockJObject(`{"v": "string"}`), args{"v"}, nil, true},
		{"string array", mockJObject(`{"v": "[1, 2, 3]"}`), args{"v"}, []interface{}{1.0, 2.0, 3.0}, false},
		{"json array", mockJObject(`{"v": [1, 2, 3]}`), args{"v"}, []interface{}{1.0, 2.0, 3.0}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.j.GetJArray(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetJArray() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetJArray() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJObject_GetJObject(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		j       JObject
		args    args
		want    JObject
		wantErr bool
	}{
		{"int value", mockJObject(`{"v": 123}`), args{"v"}, nil, true},
		{"float value", mockJObject(`{"v": 1.23}`), args{"v"}, nil, true},
		{"string value", mockJObject(`{"v": "string"}`), args{"v"}, nil, true},
		{"string object", mockJObject(`{"v": "{\"k\": 123}"}`), args{"v"}, map[string]interface{}{"k": 123.0}, false},
		{"json object", mockJObject(`{"v": {"k": 123}}`), args{"v"}, map[string]interface{}{"k": 123.0}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.j.GetJObject(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetJObject() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetJObject() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJObject_GetString(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		j       JObject
		args    args
		want    string
		wantErr bool
	}{
		{"string float value", mockJObject(`{"v": "12.345"}`), args{"v"}, "12.345", false},
		{"string int value", mockJObject(`{"v": "12"}`), args{"v"}, "12", false},
		{"string value", mockJObject(`{"v": "string"}`), args{"v"}, "string", false},
		{"float value", mockJObject(`{"v": 12.345}`), args{"v"}, "12.345", false},
		{"int value", mockJObject(`{"v": 12}`), args{"v"}, "12", false},
		{"json object", mockJObject(`{"v": {"k": 123}}`), args{"v"}, `{"k":123}`, false},
		{"json array", mockJObject(`{"v": [1, 2, 3]}`), args{"v"}, `[1,2,3]`, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.j.GetString(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetString() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJObject_GetStringDef(t *testing.T) {
	type args struct {
		key string
		def string
	}
	tests := []struct {
		name string
		j    JObject
		args args
		want string
	}{
		{"string float value", mockJObject(`{"v": "12.345"}`), args{"k", ""}, ""},
		{"string int value", mockJObject(`{"v": "12"}`), args{"k", ""}, ""},
		{"string value", mockJObject(`{"v": "string"}`), args{"k", ""}, ""},
		{"float value", mockJObject(`{"v": 12.345}`), args{"k", ""}, ""},
		{"int value", mockJObject(`{"v": 12}`), args{"k", ""}, ""},
		{"json object", mockJObject(`{"v": {"k": 123}}`), args{"k", ""}, ``},
		{"json array", mockJObject(`{"v": [1, 2, 3]}`), args{"k", ""}, ``},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.j.GetStringDef(tt.args.key, tt.args.def); got != tt.want {
				t.Errorf("GetStringDef() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJObject_GetUint16(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		j       JObject
		args    args
		want    uint16
		wantErr bool
	}{
		{"string float value", mockJObject(`{"v": "12.345"}`), args{"v"}, 0, true},
		{"string int value", mockJObject(`{"v": "12"}`), args{"v"}, 12, false},
		{"string value", mockJObject(`{"v": "string"}`), args{"v"}, 0, true},
		{"float value", mockJObject(`{"v": 12.345}`), args{"v"}, 0, true},
		{"int value", mockJObject(`{"v": 12}`), args{"v"}, 12, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.j.GetUint16(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUint16() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetUint16() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJObject_GetUint16Def(t *testing.T) {
	type args struct {
		key string
		def uint16
	}
	tests := []struct {
		name string
		j    JObject
		args args
		want uint16
	}{
		{"string float value", mockJObject(`{"v": "12.345"}`), args{"v", 0}, 0},
		{"string int value", mockJObject(`{"v": "12"}`), args{"v", 0}, 12},
		{"string value", mockJObject(`{"v": "string"}`), args{"v", 0}, 0},
		{"float value", mockJObject(`{"v": 12.345}`), args{"v", 0}, 0},
		{"int value", mockJObject(`{"v": 12}`), args{"v", 0}, 12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.j.GetUint16Def(tt.args.key, tt.args.def); got != tt.want {
				t.Errorf("GetUint16Def() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJObject_GetUint32(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		j       JObject
		args    args
		want    uint32
		wantErr bool
	}{
		{"string float value", mockJObject(`{"v": "12.345"}`), args{"v"}, 0, true},
		{"string int value", mockJObject(`{"v": "12"}`), args{"v"}, 12, false},
		{"string value", mockJObject(`{"v": "string"}`), args{"v"}, 0, true},
		{"float value", mockJObject(`{"v": 12.345}`), args{"v"}, 0, true},
		{"int value", mockJObject(`{"v": 12}`), args{"v"}, 12, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.j.GetUint32(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUint32() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetUint32() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJObject_GetUint32Def(t *testing.T) {
	type args struct {
		key string
		def uint32
	}
	tests := []struct {
		name string
		j    JObject
		args args
		want uint32
	}{
		{"string float value", mockJObject(`{"v": "12.345"}`), args{"v", 0}, 0},
		{"string int value", mockJObject(`{"v": "12"}`), args{"v", 0}, 12},
		{"string value", mockJObject(`{"v": "string"}`), args{"v", 0}, 0},
		{"float value", mockJObject(`{"v": 12.345}`), args{"v", 0}, 0},
		{"int value", mockJObject(`{"v": 12}`), args{"v", 0}, 12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.j.GetUint32Def(tt.args.key, tt.args.def); got != tt.want {
				t.Errorf("GetUint32Def() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJObject_GetUint64(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		j       JObject
		args    args
		want    uint64
		wantErr bool
	}{
		{"string float value", mockJObject(`{"v": "12.345"}`), args{"v"}, 0, true},
		{"string int value", mockJObject(`{"v": "12"}`), args{"v"}, 12, false},
		{"string value", mockJObject(`{"v": "string"}`), args{"v"}, 0, true},
		{"float value", mockJObject(`{"v": 12.345}`), args{"v"}, 0, true},
		{"int value", mockJObject(`{"v": 12}`), args{"v"}, 12, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.j.GetUint64(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUint64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetUint64() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJObject_GetUint64Def(t *testing.T) {
	type args struct {
		key string
		def uint64
	}
	tests := []struct {
		name string
		j    JObject
		args args
		want uint64
	}{
		{"string float value", mockJObject(`{"v": "12.345"}`), args{"v", 0}, 0},
		{"string int value", mockJObject(`{"v": "12"}`), args{"v", 0}, 12},
		{"string value", mockJObject(`{"v": "string"}`), args{"v", 0}, 0},
		{"float value", mockJObject(`{"v": 12.345}`), args{"v", 0}, 0},
		{"int value", mockJObject(`{"v": 12}`), args{"v", 0}, 12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.j.GetUint64Def(tt.args.key, tt.args.def); got != tt.want {
				t.Errorf("GetUint64Def() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJObject_GetUint8(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		j       JObject
		args    args
		want    uint8
		wantErr bool
	}{
		{"string float value", mockJObject(`{"v": "12.345"}`), args{"v"}, 0, true},
		{"string int value", mockJObject(`{"v": "12"}`), args{"v"}, 12, false},
		{"string value", mockJObject(`{"v": "string"}`), args{"v"}, 0, true},
		{"float value", mockJObject(`{"v": 12.345}`), args{"v"}, 0, true},
		{"int value", mockJObject(`{"v": 12}`), args{"v"}, 12, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.j.GetUint8(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUint8() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetUint8() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJObject_GetUint8Def(t *testing.T) {
	type args struct {
		key string
		def uint8
	}
	tests := []struct {
		name string
		j    JObject
		args args
		want uint8
	}{
		{"string float value", mockJObject(`{"v": "12.345"}`), args{"v", 0}, 0},
		{"string int value", mockJObject(`{"v": "12"}`), args{"v", 0}, 12},
		{"string value", mockJObject(`{"v": "string"}`), args{"v", 0}, 0},
		{"float value", mockJObject(`{"v": 12.345}`), args{"v", 0}, 0},
		{"int value", mockJObject(`{"v": 12}`), args{"v", 0}, 12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.j.GetUint8Def(tt.args.key, tt.args.def); got != tt.want {
				t.Errorf("GetUint8Def() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewJsonObject(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name    string
		args    args
		want    JObject
		wantErr bool
	}{
		{"mock test", args{[]byte(`{"v": 123}`)}, map[string]interface{}{"v": 123.0}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewJObject(tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewJObject() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewJObject() got = %v, want %v", got, tt.want)
			}
		})
	}
}
