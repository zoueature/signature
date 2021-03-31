package signature

import (
	"encoding/json"
	"fmt"
	"reflect"
	"sort"
	"strings"
)

// SignatureJson
// 生成json的签名字符串
func SignatureJson(str []byte) (string, error) {
	if len(str) == 0 {
		return "", nil
	}
	container := make(map[string]interface{})
	err := json.Unmarshal(str, &container)
	if err != nil {
		return "", err
	}
	formatStr := SortMap(container)
	return formatStr, nil
}

func SortMap(m map[string]interface{}) string {
	if len(m) == 0 {
		return ""
	}
	keys := make([]string, 0)
	for k, _ := range m {
		keys = append(keys, k)
	}
	sortedKeys := sort.StringSlice(keys)
	sortedKeys.Sort()
	kvArray := make([]string, 0)
	for _, key := range sortedKeys {
		vStr := ""
		v := reflect.ValueOf(m[key])
		kind := v.Type().Kind()
		switch kind {
		case reflect.Slice:
			fallthrough
		case reflect.Array:
			vStr = sortSlice(v.Interface().([]interface{}))
		case reflect.Map:
			vStr = SortMap(v.Interface().(map[string]interface{}))
		default:
			vStr = fmt.Sprintf("%v", v.Interface())
		}
		kvArray = append(kvArray, fmt.Sprintf("%s:%s", key, vStr))
	}
	return "{" + strings.Join(kvArray, ",") + "}"
}

func sortSlice(s []interface{}) string {
	if len(s) == 0 {
		return ""
	}
	strArray := make([]string, 0, len(s))
	for _, v := range s {
		str := ""
		value := reflect.ValueOf(v)
		switch value.Kind() {
		case reflect.Array:
			fallthrough
		case reflect.Slice:
			str = sortSlice(value.Interface().([]interface{}))
		case reflect.Map:
			str = SortMap(value.Interface().(map[string]interface{}))
		default:
			str = fmt.Sprintf("%v", v)
		}
		strArray = append(strArray, str)
	}
	sortedArray := sort.StringSlice(strArray)
	sortedArray.Sort()
	return "[" + strings.Join(sortedArray, ",") + "]"
}
