package main

import (
	"built-in-logger/writer"
	"bytes"
	"fmt"
	"log"
	"os"
)

func LogLevelExample() {
	// 디버그 로거는 표준 출력만 사용
	lDebug := log.New(os.Stdout, "[Debug]", log.Lshortfile)

	// 에러 로거는 디버그 로거의 출력 및 인메모리 버퍼에 모두 전달
	inMemory := new(bytes.Buffer)
	w := writer.NewSustainedMultiWriter(inMemory, lDebug.Writer())
	lError := log.New(w, "[Error]", log.LstdFlags|log.Lshortfile)

	fmt.Println("Standard output:")
	lError.Println("failed to connect with database")
	lDebug.Println("can we connect to the database...?")

	fmt.Println("\nIn-memory Log:")
	fmt.Println(inMemory.String())
}
