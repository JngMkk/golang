package main

import (
	"fmt"

	"github.com/JngMkk/cookbook/ch2/ansicolor/color"
)

func main() {
	r := color.ColorText{
		TextColor: color.Red,
		Text:      "It's Red",
	}
	fmt.Println(r.String())

	r.TextColor = color.Green
	r.Text = "Now, It's Green"
	fmt.Println(r.String())

	r.TextColor = color.ColorNone
	r.Text = "Back to noraml..."
	fmt.Println(r.String())
}
