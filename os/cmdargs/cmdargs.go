package main

import (
	"flag"
	"fmt"
	"os"
)

const version = "1.0.0"

const usage = `Usage:

	%s [command]

Commands:

	greet
	version
`

const greetUsage = `Usage:

	%s greet name [flag]

Positional Arguments:
	name
	the name to greet

Flags:
`

// 중첩된 명령줄에 대한 모든 레벨 저장
type MenuConf struct {
	Goodbye bool
}

// 기본 플래그 설정
func (m *MenuConf) SetupMenu() *flag.FlagSet {
	menu := flag.NewFlagSet("menu", flag.ExitOnError)
	menu.Usage = func() {
		fmt.Printf(usage, os.Args[0])
		menu.PrintDefaults()
	}

	return menu
}

// 하위 메뉴에 대한 플래그 모음 반환
func (m *MenuConf) GetSubMenu() *flag.FlagSet {
	submenu := flag.NewFlagSet("submenu", flag.ExitOnError)
	submenu.BoolVar(&m.Goodbye, "goodbye", false, "Say Goodbye instead of Hello")
	submenu.Usage = func() {
		fmt.Printf(greetUsage, os.Args[0])
		submenu.PrintDefaults()
	}

	return submenu
}

// greet 명령에 의해 호출
func (m *MenuConf) Greet(name string) {
	if m.Goodbye {
		fmt.Println("Goodbye " + name + "!")
	} else {
		fmt.Println("Hello " + name + "!")
	}
}

// 현재 버전 출력
func (m *MenuConf) Version() {
	fmt.Println("Version: " + version)
}
