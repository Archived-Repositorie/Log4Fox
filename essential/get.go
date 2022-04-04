package essential

import (
	"path"
	"runtime"
)

func Locate(skip int) (fileName, funcName string, line int, ok bool) {
	pc, file, line, ok := runtime.Caller(skip)
	fileName = path.Base(file)
	funcName = runtime.FuncForPC(pc).Name()
	return
}
