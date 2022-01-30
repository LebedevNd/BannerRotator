package internalhttp

import (
	"fmt"
	"net/http"
	"time"

	"github.com/golang-module/carbon"
)

type Logger struct {
	handler http.Handler
}

func (l *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	rCtx := r.Clone(r.Context())
	l.handler.ServeHTTP(w, rCtx)

	log := r.RemoteAddr + " " +
		carbon.Now().Format("Y-m-d H:i:s") + " " +
		r.Method + " " +
		r.Proto + " " +
		r.RequestURI + " " +
		r.UserAgent() +
		"; request_time: " + time.Since(start).String()
	fmt.Println(log)
}

func NewLogger(handlerToWrap http.Handler) *Logger {
	return &Logger{handlerToWrap}
}
