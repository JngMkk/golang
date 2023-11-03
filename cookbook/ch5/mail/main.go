package main

import (
	"fmt"
	"io"
	"log"
	"net/mail"
	"os"
	"strings"
)

const msg string = `Date: Thu, 24 Jul 2022 08:00:00 -0700
From: JngMkk <JngMkk@example.com>
To: YooHyun <YooHyun@example.com>
Subject: Hello !!

Hi Hello nihao Annyung
`

// 헤더 정보 추출
func printHeaderInfo(header mail.Header) {
	// 단일 주소라는 것을 알기 때문에 다음 코드가 동작
	// 단일 주소가 아닌 경우 ParseAddressList 사용
	toAddr, err := mail.ParseAddress(header.Get("To"))
	if err == nil {
		fmt.Printf("To: %s <%s>\n", toAddr.Name, toAddr.Address)
	}

	fromAddr, err := mail.ParseAddress(header.Get("From"))
	if err == nil {
		fmt.Printf("From: %s <%s>\n", fromAddr.Name, fromAddr.Address)
	}
	fmt.Println("Subject:", header.Get("Subject"))

	// 다음 코드는 유효한 RFC5322 날짜에 대해 동작
	// d := header.Get("Date")
	// mail.ParseDate(d) 실행

	if date, err := header.Date(); err == nil {
		fmt.Println("Date:", date)
	}

	fmt.Println(strings.Repeat("=", 40))
	fmt.Println()
}

func main() {
	r := strings.NewReader(msg)
	m, err := mail.ReadMessage(r)
	if err != nil {
		log.Fatal(err)
	}

	printHeaderInfo(m.Header)

	if _, err := io.Copy(os.Stdout, m.Body); err != nil {
		log.Fatal(err)
	}
}
