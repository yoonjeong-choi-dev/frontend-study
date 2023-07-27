package main

import (
	"fmt"
	"go.uber.org/zap/zapcore"
)

// encoderConfig zap 로거가 출력을 하기전에 로그 데이터를 인코딩하는데 사용하는 설정
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
	fmt.Println("Example 1. Json Encoder Logger")
	LogWithJson()

	fmt.Println("\nExample 2. Console Encoder Logger")
	LogWithConsole()

	fmt.Println("\nExample 3. Multi Output Logger")
	LogWithInfoFileAndDebugConsole()

	fmt.Println("\nExample 4. Sampling duplicated log entries")
	SamplingLogEntries()

	fmt.Println("\nExample 5. Dynamic Change of Log Level")
	DynamicLevelSetting()
}
