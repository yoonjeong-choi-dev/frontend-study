package writer

import (
	"go.uber.org/multierr"
	"io"
)

// SustainedMultiWriter io.MultiWriter 인터페이스 개선
type SustainedMultiWriter struct {
	writers []io.Writer
}

func (w *SustainedMultiWriter) Write(p []byte) (n int, err error) {
	for _, writer := range w.writers {
		i, wErr := writer.Write(p)
		n += i

		// 예외를 처리하는 대신, 각 io.Writer 구조체에서 반환한 에러를 누적
		// => 특정 writer.Write() 처리가 실패해도 이후 writer 들에 대한 처리 진행
		err = multierr.Append(err, wErr)
	}

	return n, err
}

func NewSustainedMultiWriter(writers ...io.Writer) io.Writer {
	mw := &SustainedMultiWriter{
		writers: make([]io.Writer, 0, len(writers)),
	}

	for _, w := range writers {
		if m, ok := w.(*SustainedMultiWriter); ok {
			// SustainedMultiWriter 구조체가 인자로 들어온 경우, 구조체의 writer 객체들을 모두 등록
			mw.writers = append(mw.writers, m.writers...)
			continue
		}

		// 일반적인 writer 구조체 등록
		mw.writers = append(mw.writers, w)
	}

	return mw
}
