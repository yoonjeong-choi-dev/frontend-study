package global_log

import (
	"errors"
	"github.com/sirupsen/logrus"
	"os"
	"sync"
)

// 로깅을 위한 패키지 수준 전역 변수
// => 외부 모듈에서는 로그 객체를 직접 사용 불가능
var (
	log     *logrus.Logger
	initLog sync.Once
)

func Init() error {
	// singleton pattern
	err := errors.New("already initialized")
	initLog.Do(func() {
		// 첫 초기화 시에만 err==nil
		err = nil
		log = logrus.New()
		log.Formatter = &logrus.JSONFormatter{}
		log.Out = os.Stdout
		log.Level = logrus.DebugLevel
	})
	return err
}

func SetLog(l *logrus.Logger) {
	log = l
}

func WithField(key string, value interface{}) *logrus.Entry {
	return log.WithField(key, value)
}

func Debug(args ...interface{}) {
	log.Debug(args)
}
