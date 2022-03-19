package color

import (
	"fmt"
)

func Color(c string, s ...interface{}) string {
	return c + fmt.Sprint(s...) + RESET
}

func Colorf(c string, s string, k ...interface{}) string {
	return c + fmt.Sprintf(s, k...) + RESET
}

func Colorln(c string, s ...interface{}) string {
	return c + fmt.Sprintln(s...) + RESET
}
