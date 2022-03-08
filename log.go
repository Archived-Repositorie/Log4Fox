package log4fox

import (
	"github.com/PurotoApp/Log4Fox/color"
	"github.com/PurotoApp/Log4Fox/essential"
)

func LogPrint(s string) {
	a := color.Color(color.FONT_GREEN, s)
	essential.Stderr(a)
}