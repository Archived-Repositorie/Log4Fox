package log4fox

type logs interface {
	log() error
	logf() error
	logln()	error
}

type logObject struct {
	name string
	dateFormat string
	color bool
}