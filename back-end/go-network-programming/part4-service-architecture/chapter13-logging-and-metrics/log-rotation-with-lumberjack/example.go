package main

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"io/ioutil"
	"os"
	"path/filepath"
)

var encoderConfig = zapcore.EncoderConfig{
	MessageKey:   "yj-msg",
	NameKey:      "yj-name",
	LevelKey:     "level",
	EncodeLevel:  zapcore.LowercaseLevelEncoder,
	CallerKey:    "yjCaller",
	EncodeCaller: zapcore.ShortCallerEncoder,
	TimeKey:      "time",
	EncodeTime:   zapcore.ISO8601TimeEncoder,
}

func main() {
	tempDir, err := ioutil.TempDir("", "")
	if err != nil {
		panic(err)
	}
	defer func() { _ = os.RemoveAll(tempDir) }()

	logFilePath := filepath.Join(tempDir, "debug.log")

	zl := zap.New(
		zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderConfig),
			zapcore.AddSync(
				&lumberjack.Logger{
					Filename: filepath.Join(logFilePath),
					// 로그 파일 압축 여부
					Compress: true,
					// 로깅할 어플리케이션 로컬 타임존 사용
					LocalTime: true,
					// 로그 파일 로테이션 최대 날짜, 백업 개수, 최대 크기(mb)
					MaxAge:     7,
					MaxBackups: 5,
					MaxSize:    100,
				},
			),
			zapcore.DebugLevel,
		),
	)
	defer func() { _ = zl.Sync() }()

	zl.Debug("debug message written to the log file")

	logFile, err := os.Open(logFilePath)
	if err != nil {
		panic(err)
	}
	defer func() { _ = logFile.Close() }()

	logData, err := ioutil.ReadAll(logFile)
	if err != nil {
		panic(err)
	}
	fmt.Println("Log Data:\n", string(logData))
}
