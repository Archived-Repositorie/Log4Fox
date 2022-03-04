package log4fox

import (
	"fmt"
)

func Color(c color, s string) string {
	return string(c) + s + string(C.RESET)
}

func Colorf(c color, s string, k ...interface{}) string {
	return string(c) + fmt.Sprintf(s, k...) + string(C.RESET)
}

func Colorln(c color, s ...interface{}) string {
	return string(c) + fmt.Sprintln(s...) + string(C.RESET)
}