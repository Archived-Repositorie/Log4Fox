package log4fox

import (
	"fmt"
	"time"

	_ "github.com/PurotoApp/Log4Fox/color"
	"github.com/PurotoApp/Log4Fox/essential"
)



func (l loggerObject) Logln(a ...interface{}) error {
	fmt.Println(l.color)
	fmt.Println(l.dateFormat)
	fmt.Println(l.name)
	s := fmt.Sprintf("%v [LOG] %v", time.Now().Format("2006.01.02 15:04:05"), a)
	err := essential.Stderr(fmt.Sprintln(s))
	return err
}

func (l loggerObject) Log(a ...interface{}) error {
	err := essential.Stderr(fmt.Sprint(a...))
	return err
} 

func (l loggerObject) Logf(a string, c ...interface{}) error {
	err := essential.Stderr(fmt.Sprintf(a, c...))
	return err
} 