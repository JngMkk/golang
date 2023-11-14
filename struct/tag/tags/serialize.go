package tags

import "reflect"

// 구조체를 사용자 정의 직렬화(serialization) 포맷으로 변환.
// 문자열 타입에 대한 구조체 태그를 직렬화하는 방식을 사용.
func SerializeStructStrings(s interface{}) (string, error) {
	res := ""

	r := reflect.TypeOf(s)
	val := reflect.ValueOf(s)

	// 구조체에 대한 포인터가 전달된 경우, 적절하게 처리함
	if r.Kind() == reflect.Ptr {
		r = r.Elem()
		val = val.Elem()
	}

	for i := 0; i < r.NumField(); i++ {
		field := r.Field(i)

		// 구조체 태그가 발견된 경우
		key := field.Name
		if serialize, ok := field.Tag.Lookup("serialize"); ok {
			// -를 무시. 그렇지 않으면 전체 값이 직렬화 키가 됨
			if serialize == "-" {
				continue
			}
			key = serialize
		}
		switch val.Field(i).Kind() {
		case reflect.String:
			res += key + ":" + val.Field(i).String() + ";"
		default:
			continue
		}
	}

	return res, nil
}
