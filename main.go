package log4fox

type logger interface {
	log() error
	logf() error
	logln()	error
}

type loggerObject struct {
	Name string
	DateFormat string
	Color bool
}