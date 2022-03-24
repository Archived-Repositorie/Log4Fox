package log4fox

import (
	_ "fmt"
	_ "strings"
)

type format interface {
	Date()
	Name() 
	Color() 
	Start()
	Convert() 
}
type formatObj struct {
	text []interface{}
}

func (f *formatObj) Convert() {
	print(f.text)
}