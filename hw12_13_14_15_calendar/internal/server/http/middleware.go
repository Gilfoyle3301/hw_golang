package internalhttp

import (
	"net/http"
	"strconv"
	"time"

	"github.com/fixme_my_friend/hw12_13_14_15_calendar/internal/app"
)

type wrapper struct {
	http.ResponseWriter
	statusCode int
}

func loggingMiddleware(logger app.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			logger.Info("Incoming request: " + r.Method + " " + r.URL.Path)
			wrapperWrite := &wrapper{ResponseWriter: w, statusCode: http.StatusOK}
			next.ServeHTTP(wrapperWrite, r)
			duration := time.Since(start)
			logger.Info("Response: " + http.StatusText(wrapperWrite.statusCode) + "(" + strconv.Itoa(wrapperWrite.statusCode) + ")" + duration.String())
		})
	}

}
