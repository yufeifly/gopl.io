package fucklog

import (
	"context"
	"errors"
	"github.com/sirupsen/logrus"
)

var logger2 = logrus.New()

func init() {
	logger2.SetLevel(logrus.DebugLevel)
}

func Fuck(ctx context.Context, logger *logrus.Logger) {
	header := "client.SendContainerCreate"

	logentry := logger2.WithFields(logrus.Fields{
		"name": "sucker",
	})

	logentry.Debug("duck you")

	//logrus.SetLevel(logrus.DebugLevel)
	logrus.WithFields(logrus.Fields{
		"Cmd": string("/bin/sh"),
	}).Debug("the command is")

	logger.WithFields(logrus.Fields{
		"DestUrl": "127.0.0.1",
	}).Info("SendContainerCreate")
	logger.WithContext(ctx).WithFields(logrus.Fields{
		"animal": "yufei",
		"fucker": "alan",
	}).Info("i am logger")

	logger.WithFields(logrus.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

	logger.WithFields(logrus.Fields{
		"omg":    true,
		"number": 122,
	}).Warn("The group's number increased tremendously!")

	logger.WithFields(logrus.Fields{
		"number": 123,
	}).Error("The group's number increased tremendously!")

	logger.WithFields(logrus.Fields{
		"number": 124,
	}).Panic("The group's number increased tremendously!")
	logger.Errorf("%s: new request %v", header, errors.New("fucker"))
	logger.WithFields(logrus.Fields{
		"omg":    true,
		"number": 100,
	}).Fatal("The ice breaks!")

	// A common pattern is to re-use fields between logging statements by re-using
	// the logrus.Entry returned from WithFields()
	contextLogger := logger.WithFields(logrus.Fields{
		"common": "this is a common field",
		"other":  "I also should be logged always",
	})

	logger.Printf("err: %v\n", "fuck")
	contextLogger.Info("I'll be logged with common and other field")
	contextLogger.Info("Me too")
}
