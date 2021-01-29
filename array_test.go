package goson

import (
	"reflect"
	"testing"
)

func mockJArray(str string) JArray {
	a, _ := NewJArray([]byte(str))
	return a
}

func TestJArray_GetFloat32(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name    string
		a       JArray
		args    args
		want    float32
		wantErr bool
	}{
		{"test float value", mockJArray(`[3.5, 4.5]`), args{0}, 3.5, false},
		{"test int value", mockJArray(`[3, 4]`), args{0}, 3.0, false},
		{"test string int value", mockJArray(`["3", "4"]`), args{0}, 3.0, false},
		{"test string", mockJArray(`["non-value", "test"]`), args{0}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.a.GetFloat32(tt.args.i)
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

func TestJArray_GetFloat64(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name    string
		a       JArray
		args    args
		want    float64
		wantErr bool
	}{
		{"test float value", mockJArray(`[3.5, 4.5]`), args{1}, 4.5, false},
		{"test int value", mockJArray(`[3, 4]`), args{0}, 3.0, false},
		{"test string int value", mockJArray(`["3", "4"]`), args{0}, 3.0, false},
		{"test string", mockJArray(`["non-value", "test"]`), args{0}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.a.GetFloat64(tt.args.i)
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

func TestJArray_GetInt(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name    string
		a       JArray
		args    args
		want    int
		wantErr bool
	}{
		{"test float value", mockJArray(`[3.5, 4.5]`), args{1}, 0, true},
		{"test int value", mockJArray(`[3, 4]`), args{0}, 3, false},
		{"test string int value", mockJArray(`["3", "4"]`), args{0}, 3, false},
		{"test string", mockJArray(`["non-value", "test"]`), args{0}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.a.GetInt(tt.args.i)
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

func TestJArray_GetInt16(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name    string
		a       JArray
		args    args
		want    int16
		wantErr bool
	}{
		{"test float value", mockJArray(`[3.5, 4.5]`), args{1}, 0, true},
		{"test int value", mockJArray(`[3, 4]`), args{0}, 3, false},
		{"test big int value", mockJArray(`[3000000000, 4]`), args{0}, 0, true},
		{"test string int value", mockJArray(`["3", "4"]`), args{0}, 3, false},
		{"test string big int value", mockJArray(`["3000000000", "4"]`), args{0}, 0, true},
		{"test string", mockJArray(`["non-value", "test"]`), args{0}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.a.GetInt16(tt.args.i)
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

func TestJArray_GetInt32(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name    string
		a       JArray
		args    args
		want    int32
		wantErr bool
	}{
		{"test float value", mockJArray(`[3.5, 4.5]`), args{1}, 0, true},
		{"test int value", mockJArray(`[3, 4]`), args{0}, 3, false},
		{"test big int value", mockJArray(`[3000000000, 4]`), args{0}, 0, true},
		{"test string int value", mockJArray(`["3", "4"]`), args{0}, 3, false},
		{"test string big int value", mockJArray(`["3000000000", "4"]`), args{0}, 0, true},
		{"test string", mockJArray(`["non-value", "test"]`), args{0}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.a.GetInt32(tt.args.i)
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

func TestJArray_GetInt64(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name    string
		a       JArray
		args    args
		want    int64
		wantErr bool
	}{
		{"test float value", mockJArray(`[3.5, 4.5]`), args{1}, 0, true},
		{"test int value", mockJArray(`[3, 4]`), args{0}, 3, false},
		{"test big int value", mockJArray(`[30000000000000000000000000000000000, 4]`), args{0}, 0, true},
		{"test string int value", mockJArray(`["3", "4"]`), args{0}, 3, false},
		{"test string big int value", mockJArray(`["30000000000000000000000000000000000", "4"]`), args{0}, 0, true},
		{"test string", mockJArray(`["non-value", "test"]`), args{0}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.a.GetInt64(tt.args.i)
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

func TestJArray_GetInt8(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name    string
		a       JArray
		args    args
		want    int8
		wantErr bool
	}{
		{"test float value", mockJArray(`[3.5, 4.5]`), args{1}, 0, true},
		{"test int value", mockJArray(`[3, 4]`), args{0}, 3, false},
		{"test big int value", mockJArray(`[3000000000, 4]`), args{0}, 0, true},
		{"test string int value", mockJArray(`["3", "4"]`), args{0}, 3, false},
		{"test string big int value", mockJArray(`["3000000000", "4"]`), args{0}, 0, true},
		{"test string", mockJArray(`["non-value", "test"]`), args{0}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.a.GetInt8(tt.args.i)
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

func TestJArray_GetJArray(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name    string
		a       JArray
		args    args
		want    JArray
		wantErr bool
	}{
		{"test value slice", mockJArray(`[1, 2, 3]`), args{0}, nil, true},
		{"test int slice", mockJArray(`[[1, 2, 3], [4, 5, 6]]`), args{0}, mockJArray(`[1, 2, 3]`), false},
		{"test float slice", mockJArray(`[[1.1, 2.2, 3.3], [4.4, 5.5, 6.6]]`), args{0}, mockJArray(`[1.1, 2.2, 3.3]`), false},
		{"test string slice", mockJArray(`[["hello", "world"], ["this", "is", "json"]]`), args{0}, mockJArray(`["hello", "world"]`), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.a.GetJArray(tt.args.i)
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

func TestJArray_GetJObject(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name    string
		a       JArray
		args    args
		want    JObject
		wantErr bool
	}{
		{"test obj slice", mockJArray(`[{"k": 1}]`), args{0}, mockJObject(`{"k": 1}`), false},
		{"test int slice", mockJArray(`[1, 2, 3]`), args{0}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.a.GetJObject(tt.args.i)
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

func TestJArray_GetString(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name    string
		a       JArray
		args    args
		want    string
		wantErr bool
	}{
		{"test obj slice", mockJArray(`[{"k": 1}]`), args{0}, `{"k":1}`, false},
		{"test int slice", mockJArray(`[1, 2, 3]`), args{0}, "1", false},
		{"test float slice", mockJArray(`[1.1, 2.2, 3.3]`), args{0}, "1.1", false},
		{"test string slice", mockJArray(`["this", "is", "json"]`), args{2}, "json", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.a.GetString(tt.args.i)
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

func TestJArray_GetUint16(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name    string
		a       JArray
		args    args
		want    uint16
		wantErr bool
	}{
		{"test float value", mockJArray(`[3.5, 4.5]`), args{1}, 0, true},
		{"test int value", mockJArray(`[3, 4]`), args{0}, 3, false},
		{"test big int value", mockJArray(`[3000000000, 4]`), args{0}, 0, true},
		{"test string int value", mockJArray(`["3", "4"]`), args{0}, 3, false},
		{"test string big int value", mockJArray(`["3000000000", "4"]`), args{0}, 0, true},
		{"test string", mockJArray(`["non-value", "test"]`), args{0}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.a.GetUint16(tt.args.i)
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

func TestJArray_GetUint32(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name    string
		a       JArray
		args    args
		want    uint32
		wantErr bool
	}{
		{"test float value", mockJArray(`[3.5, 4.5]`), args{1}, 0, true},
		{"test int value", mockJArray(`[3, 4]`), args{0}, 3, false},
		{"test big int value", mockJArray(`[8000000000, 4]`), args{0}, 0, true},
		{"test string int value", mockJArray(`["3", "4"]`), args{0}, 3, false},
		{"test string big int value", mockJArray(`["8000000000", "4"]`), args{0}, 0, true},
		{"test string", mockJArray(`["non-value", "test"]`), args{0}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.a.GetUint32(tt.args.i)
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

func TestJArray_GetUint64(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name    string
		a       JArray
		args    args
		want    uint64
		wantErr bool
	}{
		{"test float value", mockJArray(`[3.5, 4.5]`), args{1}, 0, true},
		{"test int value", mockJArray(`[3, 4]`), args{0}, 3, false},
		{"test big int value", mockJArray(`[30000000000000000000000, 4]`), args{0}, 0, true},
		{"test string int value", mockJArray(`["3", "4"]`), args{0}, 3, false},
		{"test string big int value", mockJArray(`["30000000000000000000000", "4"]`), args{0}, 0, true},
		{"test string", mockJArray(`["non-value", "test"]`), args{0}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.a.GetUint64(tt.args.i)
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

func TestJArray_GetUint8(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name    string
		a       JArray
		args    args
		want    uint8
		wantErr bool
	}{
		{"test float value", mockJArray(`[3.5, 4.5]`), args{1}, 0, true},
		{"test int value", mockJArray(`[3, 4]`), args{0}, 3, false},
		{"test big int value", mockJArray(`[3000000000, 4]`), args{0}, 0, true},
		{"test string int value", mockJArray(`["3", "4"]`), args{0}, 3, false},
		{"test string big int value", mockJArray(`["3000000000", "4"]`), args{0}, 0, true},
		{"test string", mockJArray(`["non-value", "test"]`), args{0}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.a.GetUint8(tt.args.i)
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

func TestJArray_ToFloat32Slice(t *testing.T) {
	tests := []struct {
		name    string
		a       JArray
		want    []float32
		wantErr bool
	}{
		{"test string value", mockJArray(`["3.5", "4.5"]`), []float32{3.5, 4.5}, false},
		{"test float value", mockJArray(`[3.5, 4.5]`), []float32{3.5, 4.5}, false},
		{"test int value", mockJArray(`[1, 2, 3]`), []float32{1, 2, 3}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.a.ToFloat32Slice()
			if (err != nil) != tt.wantErr {
				t.Errorf("ToFloat32Slice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToFloat32Slice() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJArray_ToFloat64Slice(t *testing.T) {
	tests := []struct {
		name    string
		a       JArray
		want    []float64
		wantErr bool
	}{
		{"test string value", mockJArray(`["3.5", "4.5"]`), []float64{3.5, 4.5}, false},
		{"test float value", mockJArray(`[3.5, 4.5]`), []float64{3.5, 4.5}, false},
		{"test int value", mockJArray(`[1, 2, 3]`), []float64{1, 2, 3}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.a.ToFloat64Slice()
			if (err != nil) != tt.wantErr {
				t.Errorf("ToFloat64Slice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToFloat64Slice() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJArray_ToInt16Slice(t *testing.T) {
	tests := []struct {
		name    string
		a       JArray
		want    []int16
		wantErr bool
	}{
		{"test string value", mockJArray(`["3", "4"]`), []int16{3, 4}, false},
		{"test float value", mockJArray(`[3.5, 4.5]`), nil, true},
		{"test int value", mockJArray(`[1, 2, 3]`), []int16{1, 2, 3}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.a.ToInt16Slice()
			if (err != nil) != tt.wantErr {
				t.Errorf("ToInt16Slice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToInt16Slice() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJArray_ToInt32Slice(t *testing.T) {
	tests := []struct {
		name    string
		a       JArray
		want    []int32
		wantErr bool
	}{
		{"test string value", mockJArray(`["3", "4"]`), []int32{3, 4}, false},
		{"test float value", mockJArray(`[3.5, 4.5]`), nil, true},
		{"test int value", mockJArray(`[1, 2, 3]`), []int32{1, 2, 3}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.a.ToInt32Slice()
			if (err != nil) != tt.wantErr {
				t.Errorf("ToInt32Slice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToInt32Slice() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJArray_ToInt64Slice(t *testing.T) {
	tests := []struct {
		name    string
		a       JArray
		want    []int64
		wantErr bool
	}{
		{"test string value", mockJArray(`["3", "4"]`), []int64{3, 4}, false},
		{"test float value", mockJArray(`[3.5, 4.5]`), nil, true},
		{"test int value", mockJArray(`[1, 2, 3]`), []int64{1, 2, 3}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.a.ToInt64Slice()
			if (err != nil) != tt.wantErr {
				t.Errorf("ToInt64Slice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToInt64Slice() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJArray_ToInt8Slice(t *testing.T) {
	tests := []struct {
		name    string
		a       JArray
		want    []int8
		wantErr bool
	}{
		{"test string value", mockJArray(`["3", "4"]`), []int8{3, 4}, false},
		{"test float value", mockJArray(`[3.5, 4.5]`), nil, true},
		{"test int value", mockJArray(`[1, 2, 3]`), []int8{1, 2, 3}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.a.ToInt8Slice()
			if (err != nil) != tt.wantErr {
				t.Errorf("ToInt8Slice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToInt8Slice() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJArray_ToIntSlice(t *testing.T) {
	tests := []struct {
		name    string
		a       JArray
		want    []int
		wantErr bool
	}{
		{"test string value", mockJArray(`["3", "4"]`), []int{3, 4}, false},
		{"test float value", mockJArray(`[3.5, 4.5]`), nil, true},
		{"test int value", mockJArray(`[1, 2, 3]`), []int{1, 2, 3}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.a.ToIntSlice()
			if (err != nil) != tt.wantErr {
				t.Errorf("ToIntSlice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToIntSlice() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJArray_ToJArraySlice(t *testing.T) {
	tests := []struct {
		name    string
		a       JArray
		want    []JArray
		wantErr bool
	}{
		{"test string value", mockJArray(`[["3", "4"]]`), []JArray{[]interface{}{"3", "4"}}, false},
		{"test float value", mockJArray(`[[3.5, 4.5]]`), []JArray{[]interface{}{3.5, 4.5}}, false},
		{"test int value", mockJArray(`[[1, 2, 3]]`), []JArray{[]interface{}{1.0, 2.0, 3.0}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.a.ToJArraySlice()
			if (err != nil) != tt.wantErr {
				t.Errorf("ToJArraySlice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToJArraySlice() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJArray_ToJObjectSlice(t *testing.T) {
	tests := []struct {
		name    string
		a       JArray
		want    []JObject
		wantErr bool
	}{
		{"test JObject", mockJArray(`[{"k": 1}]`), []JObject{map[string]interface{}{"k": 1.0}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.a.ToJObjectSlice()
			if (err != nil) != tt.wantErr {
				t.Errorf("ToJObjectSlice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToJObjectSlice() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJArray_ToStringSlice(t *testing.T) {
	tests := []struct {
		name    string
		a       JArray
		want    []string
		wantErr bool
	}{
		{"test string value", mockJArray(`["3", "4"]`), []string{"3", "4"}, false},
		{"test float value", mockJArray(`[3.5, 4.5]`), []string{"3.5", "4.5"}, false},
		{"test int value", mockJArray(`[1, 2, 3]`), []string{"1", "2", "3"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.a.ToStringSlice()
			if (err != nil) != tt.wantErr {
				t.Errorf("ToStringSlice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToStringSlice() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJArray_ToUint16Slice(t *testing.T) {
	tests := []struct {
		name    string
		a       JArray
		want    []uint16
		wantErr bool
	}{
		{"test string value", mockJArray(`["3", "4"]`), []uint16{3, 4}, false},
		{"test float value", mockJArray(`[3.5, 4.5]`), nil, true},
		{"test int value", mockJArray(`[1, 2, 3]`), []uint16{1, 2, 3}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.a.ToUint16Slice()
			if (err != nil) != tt.wantErr {
				t.Errorf("ToUint16Slice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToUint16Slice() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJArray_ToUint32Slice(t *testing.T) {
	tests := []struct {
		name    string
		a       JArray
		want    []uint32
		wantErr bool
	}{
		{"test string value", mockJArray(`["3", "4"]`), []uint32{3, 4}, false},
		{"test float value", mockJArray(`[3.5, 4.5]`), nil, true},
		{"test int value", mockJArray(`[1, 2, 3]`), []uint32{1, 2, 3}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.a.ToUint32Slice()
			if (err != nil) != tt.wantErr {
				t.Errorf("ToUint32Slice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToUint32Slice() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJArray_ToUint64Slice(t *testing.T) {
	tests := []struct {
		name    string
		a       JArray
		want    []uint64
		wantErr bool
	}{
		{"test string value", mockJArray(`["3", "4"]`), []uint64{3, 4}, false},
		{"test float value", mockJArray(`[3.5, 4.5]`), nil, true},
		{"test int value", mockJArray(`[1, 2, 3]`), []uint64{1, 2, 3}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.a.ToUint64Slice()
			if (err != nil) != tt.wantErr {
				t.Errorf("ToUint64Slice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToUint64Slice() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJArray_ToUint8Slice(t *testing.T) {
	tests := []struct {
		name    string
		a       JArray
		want    []uint8
		wantErr bool
	}{
		{"test string value", mockJArray(`["3", "4"]`), []uint8{3, 4}, false},
		{"test float value", mockJArray(`[3.5, 4.5]`), nil, true},
		{"test int value", mockJArray(`[1, 2, 3]`), []uint8{1, 2, 3}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.a.ToUint8Slice()
			if (err != nil) != tt.wantErr {
				t.Errorf("ToUint8Slice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToUint8Slice() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJArray_checkIndex(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		a    JArray
		args args
		want bool
	}{
		{"success", []interface{}{1, 2, 3}, args{0}, true},
		{"failed", []interface{}{1, 2, 3}, args{3}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.checkIndex(tt.args.i); got != tt.want {
				t.Errorf("checkIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewJArray(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name    string
		args    args
		want    JArray
		wantErr bool
	}{
		{"json array", args{[]byte(`[1, 2, 3]`)}, []interface{}{1.0, 2.0, 3.0}, false},
		{"json object", args{[]byte(`{"k": 123}`)}, nil, true},
		{"non-json", args{[]byte(`"non-json string"`)}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewJArray(tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewJArray() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewJArray() got = %v, want %v", got, tt.want)
			}
		})
	}
}