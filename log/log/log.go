/*
log.Fatal() 함수는 무언가 잘못된 일이 발생했을 때 상황을 보고한 다음 즉시 프로그램을 종료하는 데 사용.
log.Fatal()을 호출하면 log.Fatal()이 호출된 지점에서 에러 메시지를 출력한 다음 프로그램이 종료됨.
대부분의 경우 에러 메시지는 파일 접근 불가, 매개변수 부족 등이 될 수 있음.
또한 log.Fatal()을 호출하면 0이 아닌 종료 코드를 반환하는데, 이는 유닉스에서 에러로 인식함.

log.Panic()은 이전에 접근한 파일에 다시 접근할 수 없거나 디스크 공간이 부족할 때와 같이 정말 예상치 못한 일이 일어난 것을 알려줌.
log.Fatal()처럼 log.Panic()도 에러 메시지를 출력하고 즉시 프로그램을 종료함.

log.Panic()은 log.Print() 호출 이후에 panic()을 호출하는 것과 동일함.
panic()은 Go의 기본 함수로, 현재 함수의 실행을 멈추고 패닉에 빠짐. 그런 다음 함수가 호출된 곳으로 돌아감.
반면 log.Fatal()은 log.Print()를 호출한 다음 os.Exit(1)을 호출함.
*/

package main

import (
	"log"
	"os"
)

func main() {
	if len(os.Args) != 1 {
		log.Fatal("Fatal!")
	}
	log.Panic("Panic!")
}
