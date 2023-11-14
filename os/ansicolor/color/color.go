package color

import "fmt"

// 텍스트 색상
type Color int

const (
	// ColorNone은 기본값
	ColorNone = iota
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
	Black Color = -1
)

// 색상의 문자열 및 색상값 저장
type ColorText struct {
	TextColor Color
	Text      string
}

func (r *ColorText) String() string {
	if r.TextColor == ColorNone {
		return r.Text
	}

	value := 30
	if r.TextColor != Black {
		value += int(r.TextColor)
	}

	return fmt.Sprintf("\033[0;%dm%s\033[0m", value, r.Text)
}
