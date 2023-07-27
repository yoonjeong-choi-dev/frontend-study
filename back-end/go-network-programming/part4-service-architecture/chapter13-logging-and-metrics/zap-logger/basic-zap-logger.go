package main

import (
	"bytes"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"runtime"
)

func LogWithJson() {
	zl := zap.New(
		zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderConfig),
			// 동시성에 안전하게 해주는 함수
			zapcore.Lock(os.Stdout),
			zapcore.DebugLevel),
		zap.AddCaller(),
		zap.Fields(
			zap.String("version", runtime.Version()),
		),
	)

	// 함수 종료 시, 버퍼링된 모든 로그 데이터 출력
	defer func() { _ = zl.Sync() }()

	// 로거의 이름 설정 => 로거 설정의 NameKey 키에 대한 값
	ex := zl.Named("example")
	ex.Debug("debug logging")
	ex.Info("info logging")
	ex.Error("error logging")
}

func LogWithConsole() {
	zl := zap.New(
		zapcore.NewCore(
			zapcore.NewConsoleEncoder(encoderConfig),
			zapcore.Lock(os.Stdout),
			zapcore.InfoLevel),
	)

	defer func() { _ = zl.Sync() }()

	ex := zl.Named("[console]")
	ex.Debug("debug logging") // 출력 X
	ex.Info("info logging")
	ex.Error("error logging")
}

// LogWithInfoFileAndDebugConsole 특정 엔트리에 대한 아웃풋을 설정하는 예시
func LogWithInfoFileAndDebugConsole() {
	file := new(bytes.Buffer)
	zl := zap.New(
		zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderConfig),
			// zapcore.AddSync io.Writer 구조체에 zapcore.WriteSyncer 인터페이스 구현
			zapcore.Lock(zapcore.AddSync(file)),
			zapcore.InfoLevel),
	)
	defer func() { _ = zl.Sync() }()

	zl.Info("this message is logged before redefining zap logger")
	zl.Info("Now we redefine zap logger with options")

	zl = zl.WithOptions(
		zap.WrapCore(
			func(c zapcore.Core) zapcore.Core {
				newConfig := encoderConfig

				// 새로운 설정에서는 레벨 정보가 대문자로 출력
				newConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

				// NewTee : io.MultiWriter 처럼 여러 코어(로거)들로 동시에 출력하게 하는 로거 생성
				return zapcore.NewTee(
					c,
					zapcore.NewCore(
						zapcore.NewConsoleEncoder(newConfig),
						zapcore.Lock(os.Stdout),
						zapcore.DebugLevel,
					),
				)
			},
		),
	)

	fmt.Println("Standard output:")
	zl.Debug("this message is logged only for console encoding")
	zl.Info("this message is logged for both console & file")

	fmt.Println("\nFile Log:")
	fmt.Println(file.String())
}
