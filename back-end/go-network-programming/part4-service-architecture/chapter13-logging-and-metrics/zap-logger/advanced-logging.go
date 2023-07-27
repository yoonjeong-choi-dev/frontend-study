package main

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/fsnotify.v1"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

// SamplingLogEntries 특정 시간동안 중복되는 로그 엔트리들을 샘플링하여 기록
func SamplingLogEntries() {
	zl := zap.New(
		// zap 코어(로거)에 샘플링 기능을 추가
		zapcore.NewSamplerWithOptions(
			zapcore.NewCore(
				zapcore.NewJSONEncoder(encoderConfig),
				zapcore.Lock(os.Stdout),
				zapcore.DebugLevel,
			),
			// 샘플링 간격, 기록할 초기 중복 로그 엔트리, n번쨰 마다 로깅할
			// -> 샘플링 간격 동안 중복해서 발생하는 n 번째 로그 엔트리를 로깅
			time.Second, 1, 3,
		),
	)
	defer func() { _ = zl.Sync() }()

	for i := 0; i < 10; i++ {
		if i == 5 {
			// 샘플링 간격만큼 멈춤
			// => 0~4 와 5~9 메시지의 중복되는 로그들의 출력 형태가 똑같음
			time.Sleep(time.Second)
		}
		zl.Debug(fmt.Sprintf("%d", i))
		zl.Debug("debug message") // 중복되는 로그 엔트리
	}
}

// DynamicLevelSetting 애플리케이션 실행 중간에 로거의 로그 레벨을 동적으로 변경
func DynamicLevelSetting() {
	tempDir, err := ioutil.TempDir("", "")
	if err != nil {
		panic(err)
	}
	defer func() { _ = os.RemoveAll(tempDir) }()

	// 세마포어 파일을 이용하여 로거 동작 변경
	// => 해당 파일이 있는 경우 디버그 레벨의 로그 출력
	debugLevelFile := filepath.Join(tempDir, "level.debug")

	// 로그 레벨에 대한 경쟁 상태를 방지하기 위한 아토믹한 레벨 정보
	atomicLevel := zap.NewAtomicLevel()

	zl := zap.New(
		zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderConfig),
			zapcore.Lock(os.Stdout),
			atomicLevel,
		),
	)
	defer func() { _ = zl.Sync() }()

	// 세마포어 파일 감시
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		panic(err)
	}
	defer func() { _ = watcher.Close() }()

	err = watcher.Add(tempDir)
	if err != nil {
		panic(err)
	}

	// 세마포어 파일을 감시하면서 로그 레벨을 변경하는 고루틴
	ready := make(chan struct{})
	go func() {
		defer close(ready)

		originalLevel := atomicLevel.Level()

		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				// 세마포어 파일 변경 이벤트
				// => 생성되는 경우 디버그 레벨로 변경
				if event.Name == debugLevelFile {
					switch {
					case event.Op&fsnotify.Create == fsnotify.Create:
						atomicLevel.SetLevel(zapcore.DebugLevel)
						ready <- struct{}{}
					case event.Op&fsnotify.Remove == fsnotify.Remove:
						atomicLevel.SetLevel(originalLevel)
						ready <- struct{}{}
					}
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				zl.Error(err.Error())
			}
		}
	}()

	zl.Debug("debug test <- cannot display")

	// 세마포어 파일 생성 -> 디버그 레벨 출력
	file, err := os.Create(debugLevelFile)
	if err != nil {
		panic(err)
	}
	err = file.Close()
	if err != nil {
		panic(err)
	}
	<-ready

	zl.Debug("debug test <- can display!")
	zl.Info("this message is always logged")
}
