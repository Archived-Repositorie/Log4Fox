package log4fox

import (
	"fmt"
	"time"

	c "github.com/PurotoApp/Log4Fox/color"
	"github.com/PurotoApp/Log4Fox/essential"
)

var (
	text = "%v [LOG] %v %v"
	color = c.FONT_CYAN
	name = "go"
)

func (l *LoggerObject) Logln(a ...interface{}) error {
	now := time.Now()
	if l.DateFormat == "" {
		l.DateFormat = "2006.01.02 15:04:05"
	}
	if !l.Color {
		color = c.RESET
	}
	if l.Name == "" {
		l.Name = name
	}
	text = fmt.Sprintf(text, now.Format(l.DateFormat), l.Name, fmt.Sprintln(a...))
	text = c.Color(color, text)
	err := essential.Stderr(text)
	return err
}

func (l *LoggerObject) Log(a ...interface{}) error {
	now := time.Now()
	if l.DateFormat == "" {
		l.DateFormat = "2006.01.02 15:04:05"
	}
	if !l.Color {
		color = c.RESET
	}
	if l.Name == "" {
		l.Name = name
	}
	text = fmt.Sprintf(text, now.Format(l.DateFormat), l.Name, fmt.Sprintln(a...))
	text = c.Color(color, text)
	err := essential.Stderr(text)
	return err
}

func (l *LoggerObject) Logf(a string, u ...interface{}) error {
	now := time.Now()
	if l.DateFormat == "" {
		l.DateFormat = "2006.01.02 15:04:05"
	}
	if !l.Color {
		color = c.RESET
	}
	if l.Name == "" {
		l.Name = name
	}
	text = fmt.Sprintf(text, now.Format(l.DateFormat), l.Name, fmt.Sprintf(a, u...))
	text = c.Color(color, text)
	err := essential.Stderr(text)
	return err
}
