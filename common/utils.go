package common

import (
	"github.com/golang/glog"
	"os"
	"runtime"
)

//returns the first non empty string
//TODO; should probably rename it to something more obvious
func WithDefault(args ...string) string {
	for _, arg := range args {
		if arg != "" {
			return arg
		}
	}
	return ""
}

//prints an informational message in the log
func LogInfo(args ...interface{}) {
	glog.Infoln(args...)
}

//logs stack and error if there is an error
//also stops the app by calling os.Exit(-1)
func LogFatal(err error, args ...interface{}) {
	LogError(err, args...)
	if err != nil {
		os.Exit(-1)
	}
}

//logs stack and error if there is an error
func LogError(err error, args ...interface{}) {
	if err == nil {
		return
	}

	buf := make([]byte, 4096)
	buf = buf[:runtime.Stack(buf, false)]

	glog.Errorln(args...)
	glog.Errorf("ERROR: %q", err)
	glog.Errorf("STACK: %s", buf)
}
