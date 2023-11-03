// flag 패키지는 명령줄 플래그를 Go 애플리케이션에 간단하게 추가할 수 있도록 도와줌.
package main

import (
	"flag"
	"fmt"
)

// 플래그 값을 저장하는 데 사용
type Config struct {
	subject      string
	isAwesome    bool
	howAwesome   int
	countTheWays CountTheWays
}

// 전달되는 플래그 값으로 설정값을 초기화 함
func (c *Config) Setup() {

	// flag 직접 설정
	// var someVar = flag.String("flag_name", "default_val", "description")

	// 구조체에 값을 넣는 것이 일반적으로 더 나음
	flag.StringVar(&c.subject, "subject", "", "subject is a string, it defaults to empty")

	// 축약어
	flag.StringVar(&c.subject, "s", "", "subject is a string, it defaults to empty (shorthand)")
	flag.BoolVar(&c.isAwesome, "isawesome", false, "is it awesome or what?")
	flag.IntVar(&c.howAwesome, "howawesome", 10, "how awesome out of 10?")

	// 커스텀 변수 타입
	flag.Var(&c.countTheWays, "c", "comma separated list of integers")
}

// 모든 내부 설정 변수를 사용해 문장을 반환
func (c *Config) GetMessage() string {
	msg := c.subject
	if c.isAwesome {
		msg += " is awesome"
	} else {
		msg += " is Not awesome"
	}
	msg = fmt.Sprintf("%s with a certainty of %d out of 10. Let me count the ways %s", msg, c.howAwesome, c.countTheWays.String())

	return msg
}
