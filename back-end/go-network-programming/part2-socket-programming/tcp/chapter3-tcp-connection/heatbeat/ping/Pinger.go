package ping

import (
	"context"
	"io"
	"time"
)

const pingIntervalDefault = 30 * time.Second

// PingWithInterval : 주기적으로 ping 요청을 보내는 고루틴 용 함수
func PingWithInterval(ctx context.Context, w io.Writer, reset <-chan time.Duration) {
	var interval time.Duration
	select {
	case <-ctx.Done():
		return
	case interval = <-reset:
	default:
		// 연결에 대한 핑 인터벌 초기화
		// : 데이터 송수신이 성공적인 경우, 하트비트 관련 요청을 보낼 필요가 없음
		if interval <= 0 {
			// 채널을 통해 받은 리셋 주기가 0보다 작은 경우에는 디폴트값으로 초기화
			interval = pingIntervalDefault
		}
	}

	timer := time.NewTimer(interval)
	defer func() {
		if !timer.Stop() {
			<-timer.C
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return
		case newInterval := <-reset:
			// 타이머가 끝나지 않은 경우, 타이머 만료 후 리셋
			// => 리셋 시그널을 받으면 "case <-time.C"에 의해 메시지 전송
			if !timer.Stop() {
				<-timer.C
			}

			// 0을 전송한 경우에는 이전에 요청한 reset 시그널에 대한 인터벌 변수를 사용
			// => 0 을 이용하여 쉽게 타이머 초기화 가능
			if newInterval > 0 {
				interval = newInterval
			}
		case <-timer.C:
			// 타이머 만료 시, ping 요청
			if _, err := w.Write([]byte("ping")); err != nil {
				// handling timeout for heartbeat
				return
			}
		}

		// 컨텍스트 취소, 타이머 리셋 시그널, 타이머 만료에 대한 작업(select) 이후, 타이머 초기화
		_ = timer.Reset(interval)
	}
}
