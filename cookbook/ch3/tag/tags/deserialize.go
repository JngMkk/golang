package tags

import (
	"errors"
	"reflect"
	"strings"
)

// 사용자 정의 직렬화 포맷을 사용해 직렬화된 문자열을 구조체로 변환
func DeSerializeStructStrings(s string, res interface{}) error {
	r := reflect.TypeOf(res)

	if r.Kind() != reflect.Ptr {
		return errors.New("res must be a pointer")
	}

	// 포인터 역참조
	r = r.Elem()
	value := reflect.ValueOf(res).Elem()

	// 직렬화된 문자열을 분할해 맵에 저장
	vals := strings.Split(s, ";")
	valMap := make(map[string]string)

	for _, v := range vals {
		keyval := strings.Split(v, ":")
		if len(keyval) != 2 {
			continue
		}
		valMap[keyval[0]] = keyval[1]
	}

	for i := 0; i < r.NumField(); i++ {
		field := r.Field(i)

		if serialize, ok := field.Tag.Lookup("serialize"); ok {
			if serialize == "-" {
				continue
			}
			if val, ok := valMap[serialize]; ok {
				value.Field(i).SetString(val)
			}
		} else if val, ok := valMap[field.Name]; ok {
			value.Field(i).SetString(val)
		}
	}

	return nil
}
