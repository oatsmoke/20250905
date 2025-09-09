package logger

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"runtime"
)

func New() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))
}

func Info(msg string) {
	slog.Info(fmt.Sprintf("[%s]: %s", getFuncName(), msg))
}

func HttpError(w http.ResponseWriter, err error, status int) {
	slog.Error(fmt.Sprintf("[%s]: %v", getFuncName(), err))
	http.Error(w, err.Error(), status)
}

func getFuncName() string {
	pc, _, _, _ := runtime.Caller(2)
	funcName := runtime.FuncForPC(pc).Name()
	return funcName
}
