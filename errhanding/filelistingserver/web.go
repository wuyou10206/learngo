package main

import (
	"learngo/errhanding/filelistingserver/filelisting"
	"net/http"
	_ "net/http/pprof"
	"os"

	"github.com/gpmgo/gopm/modules/log"
)

type appHandler func(http.ResponseWriter, *http.Request) error

type UserError interface {
	error
	Message() string
}

func errWrapper(handler appHandler) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			r := recover()
			if r != nil {
				log.Warn("Painc :%v", r)
				http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()
		err := handler(writer, request)
		if err != nil {
			log.Warn("Error handling request:%s", err.Error())
			if userErr, ok := err.(UserError); ok {
				http.Error(writer, userErr.Message(), http.StatusBadRequest)
				return
			}
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer, http.StatusText(code), code)

		}
	}
}
func main() {
	http.HandleFunc("/", errWrapper(filelisting.HandleFileList2))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
