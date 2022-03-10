package log4fox

type logger interface {
	log() error
	logf() error
	logln()	error
}

type loggerObject struct {
	name string
	dateFormat string
	color bool
}