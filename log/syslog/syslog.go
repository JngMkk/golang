/*
보통 systemd를 이용해 서비스들이 실행되므로 프로그램들은 stdout에 로그를 남겨야 하며 systemd는 이를 저널(systemd의 로그 시스템)에 기록함.
또한 클라우드 네이티브 애플리케이션에서는 stderr에 간단히 로깅하고 컨테이너 시스템이 stderr 스트림을 정해진 목적지로 전달하는 것을 권장함.

log 패키지는 로그 메시지들을 유닉스의 표준 에러로 보냄.
log 패키지 일부인 log/syslog 패키지는 로그 메시지들을 시스로그 서버로 보냄.

기본적으로 log는 표준 에러에 로그 메시지를 쓰지만 log.SetOutput()을 이용하면 동작을 바꿀 수 있음.

시스템 로그를 쓰려면 적절한 매개변수로 syslog.New()를 호출해야 함.
syslog.LOG_SYSLOG 옵션과 함께 syslog.New()를 호출하기만 하면 시스템 로그 파일에 기록할 수 있음.
그 후, log.SetOutput() 함수를 호출해 로그 정보를 새로운 곳에 보낸다고 프로그램에 알려줘야 함.
*/

package main

import (
	"log"
	"log/syslog"
)

func main() {
	sysLog, err := syslog.New(syslog.LOG_SYSLOG, "syslog.go")
	if err != nil {
		log.Println(err)
		return
	} else {
		log.SetOutput(sysLog)
		log.Printf("Everything is Fine!")
	}
}
