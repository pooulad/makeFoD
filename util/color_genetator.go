package util

import "fmt"

type Color string

const (
	ColorRed   Color = "\u001b[31m"
	ColorGreen Color = "\u001b[32m"
	ColorReset Color = "\u001b[0m"
)

func ColorGenerator(color Color, message string) {
	fmt.Println(string(color), message, string(ColorReset))
}
