// 인터페이스 : 구체화된 객체가 아닌 추상화된 상호작용으로 관계를 표현
// 메서드는 반드시 메서드명이 있어야 함
// 매개변수와 반환이 다르더라도 이름이 같은 메서드는 있을 수 없음
// 인터페이스에서는 메서드 구현을 포함하지 않음
// 추상화란?
// 내부 동작을 감춰서 서비스 제공자와 사용자 모두에게 자유를 주는 방식
// 택배를 보낼 때 Fedex로 보내든, 우체국으로 보내든 내부 발송 과정을 자세히 몰라도 보낼 수 있듯
// 추상화를 하면 의존성을 끊을 수 있음

package main

import "fmt"

type Stringer interface {
	String() string
}

type Student struct {
	Name string
	Age  int
}

func (s Student) String() string {
	return fmt.Sprintf("안녕! 나는 %d살 %s라고 해", s.Age, s.Name)
}

func main() {
	student := Student{"철수", 12}
	// var stringer Stringer

	stringer := student
	fmt.Printf("%s\n", stringer.String())
}
