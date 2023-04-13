package main

import "github.com/sirupsen/logrus"

// Hook logrus 훅 인터페이스 구현체
type Hook struct {
	Id string
}

func (hook *Hook) Fire(entry *logrus.Entry) error {
	entry.Data["id"] = hook.Id
	return nil
}

func (hook *Hook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func LogrusExample() {
	logrus.SetFormatter(&logrus.TextFormatter{})
	logrus.SetLevel(logrus.InfoLevel)
	logrus.AddHook(&Hook{Id: "yj-hook"})

	fields := logrus.Fields{}
	fields["success"] = true
	fields["complex_struct"] = struct {
		Event string
		When  string
	}{"Some Event Occurred", "Just Now"}

	x := logrus.WithFields(fields)
	x.Warn("Warning!")
	x.Error("Error!")
}
