package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
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
	zl := zap.New(
		zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderConfig),
			zapcore.Lock(os.Stdout),
			zapcore.DebugLevel,
		),
	)
	defer func() { _ = zl.Sync() }()

	ts := httptest.NewServer(
		WideEventLogMiddleware(zl, http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				defer func(r io.ReadCloser) {
					_, _ = io.Copy(ioutil.Discard, r)
					_ = r.Close()
				}(r.Body)

				_, _ = w.Write([]byte("wide logging test"))
			},
		)),
	)
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		zl.Fatal(err.Error())
	}
	_ = res.Body.Close()
}
