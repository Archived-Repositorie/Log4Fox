//Library for easier managing logging stuff in GO
package log4fox

//Base interface for every logger 
type Logger interface {
	Log() error
	Logf() error
	Logln()	error
}

/*
Base structure for every logger
Using `Name` property you can change how name of app is shown 
Using `DateFormat` property you can change how date is shown
Using `Color` property you can change if you want use color or not
*/
type LoggerObject struct {
	Name string
	DateFormat string
	Color bool
}