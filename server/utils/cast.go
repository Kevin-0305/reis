package utils

import (
	"strconv"
	"time"

	jsoniter "github.com/json-iterator/go"
)

func ToExcelData(any interface{}) string {

	if any == nil {
		return ""
	}
	switch value := any.(type) {
	case int:
		return strconv.Itoa(value)
	case int8:
		return strconv.Itoa(int(value))
	case int16:
		return strconv.Itoa(int(value))
	case int32:
		return strconv.Itoa(int(value))
	case int64:
		return strconv.FormatInt(value, 10)
	case uint:
		return strconv.FormatUint(uint64(value), 10)
	case uint8:
		return strconv.FormatUint(uint64(value), 10)
	case uint16:
		return strconv.FormatUint(uint64(value), 10)
	case uint32:
		return strconv.FormatUint(uint64(value), 10)
	case uint64:
		return strconv.FormatUint(value, 10)
	case float32:
		return strconv.FormatFloat(float64(value), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(value, 'f', -1, 64)
	case bool:
		return strconv.FormatBool(value)
	case string:
		return value
	case []byte:
		return string(value)
	case time.Time:
		if value.IsZero() {
			return ""
		}
		return value.String()
	case *time.Time:
		if value == nil {
			return ""
		}
		return value.String()
	case []int:
		var json = jsoniter.ConfigCompatibleWithStandardLibrary
		b, _ := json.Marshal(value)
		return string(b)
	case []int32:
		var json = jsoniter.ConfigCompatibleWithStandardLibrary
		b, _ := json.Marshal(value)
		return string(b)
	case []int16:
		var json = jsoniter.ConfigCompatibleWithStandardLibrary
		b, _ := json.Marshal(value)
		return string(b)
	case []int8:
		var json = jsoniter.ConfigCompatibleWithStandardLibrary
		b, _ := json.Marshal(value)
		return string(b)
	case []int64:
		var json = jsoniter.ConfigCompatibleWithStandardLibrary
		b, _ := json.Marshal(value)
		return string(b)
	case []float64:
		var json = jsoniter.ConfigCompatibleWithStandardLibrary
		b, _ := json.Marshal(value)
		return string(b)
	case []float32:
		var json = jsoniter.ConfigCompatibleWithStandardLibrary
		b, _ := json.Marshal(value)
		return string(b)
	case []uint64:
		var json = jsoniter.ConfigCompatibleWithStandardLibrary
		b, _ := json.Marshal(value)
		return string(b)
	case []uint16:
		var json = jsoniter.ConfigCompatibleWithStandardLibrary
		b, _ := json.Marshal(value)
		return string(b)
	case []string:
		var json = jsoniter.ConfigCompatibleWithStandardLibrary
		b, _ := json.Marshal(value)
		return string(b)
	case []interface{}:
		var json = jsoniter.ConfigCompatibleWithStandardLibrary
		b, _ := json.Marshal(value)
		return string(b)
	case map[string]interface{}:
		var json = jsoniter.ConfigCompatibleWithStandardLibrary
		b, _ := json.Marshal(value)
		return string(b)
	default:
		return ""
	}
}
