package dataconv

import "fmt"

// 인터페이스 타입을 기반으로 출력
func CheckType(s interface{}) {
	switch s.(type) {
	case string:
		fmt.Println("It's a string")
	case int:
		fmt.Println("It's a int")
	default:
		fmt.Println("not sure what it is...")
	}
}

// 익명의 인터페이스를 다른 타입으로 형 변환하는 방법
func Interfaces() {
	CheckType("test")
	CheckType(1)
	CheckType(false)

	var i interface{} = "test"

	if val, ok := i.(string); ok {
		fmt.Println("val is", val)
	}

	if _, ok := i.(int); !ok {
		fmt.Println("?")
	}
}
