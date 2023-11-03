// 인터페이스 사이즈
// 인스턴스 메모리 주소 + 타입 정보
package main

import (
	"fmt"
	"unsafe"
)

type Stringer interface {
	String() string
}

type Student struct {
	Name string
}

func (s *Student) String() string {
	return s.Name
}

type User struct {
	Name string
	Age  int
}

func (u User) String() string {
	return u.Name
}

func main() {
	var stringer1 Stringer
	fmt.Printf("type:%T size:%d\n", stringer1, unsafe.Sizeof(stringer1))

	student := &Student{"AAA"}
	stringer1 = student
	fmt.Printf("type:%T size:%d\n", stringer1, unsafe.Sizeof(stringer1))

	var stringer2 Stringer
	fmt.Printf("type:%T size:%d\n", stringer2, unsafe.Sizeof(stringer2))

	user := User{"BBB", 20}
	stringer2 = user
	fmt.Printf("type:%T size:%d\n", stringer2, unsafe.Sizeof(stringer2))
}
