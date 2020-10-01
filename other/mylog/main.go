package main

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/gopl.io/other/mylog/fucklog"
	"os"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	//log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	//log.SetOutput(os.Stdout)
	Log.Out = os.Stdout

	// Only log the warning severity or above.
	logrus.SetLevel(logrus.DebugLevel)
}

var Log = logrus.New()
var ctx = context.Background()

func main() {

	fucklog.Fuck(ctx, Log)
}
