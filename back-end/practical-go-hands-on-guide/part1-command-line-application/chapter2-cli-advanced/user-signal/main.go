package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"
)

func createContextWithTimeout(d time.Duration) (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), d)
	return ctx, cancel
}

// 사용자가 입력한 사용자 시그널에 대한 핸들러 등록
func registerSignalHandler(w io.Writer, cancelFunc context.CancelFunc) {
	c := make(chan os.Signal, 1)

	// SIGINT: Ctrl+C, SIGTERM: kill command
	// kill -9 [PID] : Killed 9
	// kill -15 [PID} : terminated
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		s := <-c
		fmt.Fprintf(w, "Got signal: %v\n", s)

		// 위 두 개 시그널 발생 시 컨텍스트 종료
		cancelFunc()
	}()
}

func executeCommand(ctx context.Context, command string, arg string) error {
	// 컨텍스트가 취소되면 에러 반환
	return exec.CommandContext(ctx, command, arg).Run()
}

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stdout, "Usage: %s <command> <argument>\n", os.Args[0])
		os.Exit(1)
	}

	command := os.Args[1]
	arg := os.Args[2]

	// create Context
	commandTimeout := 30 * time.Second
	ctx, cancel := createContextWithTimeout(commandTimeout)

	// Signal Handling
	registerSignalHandler(os.Stdout, cancel)

	// execute external command
	err := executeCommand(ctx, command, arg)
	if err != nil {
		if _, ok := err.(*exec.ExitError); ok {
			fmt.Fprintf(os.Stdout, "Your command must be terminated in %.1f(s)\n", commandTimeout.Seconds())
		}

		fmt.Fprintln(os.Stdout, err)
		os.Exit(1)
	}

	fmt.Fprintln(os.Stdout, "Executed Success")
}

func simpleCommandContextExample() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// context 는 3초 뒤 만료.
	// "sleep 10" 은 10초간 멈추는 멍령어 => context가 3초뒤 만료되어 강제로 프로세스 종료
	if err := exec.CommandContext(ctx, "sleep", "10").Run(); err != nil {
		fmt.Fprintln(os.Stdout, err)
	}
}
