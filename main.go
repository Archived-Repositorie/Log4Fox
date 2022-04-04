package log4fox

type Logger interface {
	Log() error
	Logf() error
	Logln() error
}

type LoggerObject struct {
	Format string
}
