package handlers

import "net/http"

func init() {
	println("some init in common handlers")
}

type CommonHandler struct {
}

func NewCommonhandler() *CommonHandler {
	return &CommonHandler{}
}

func (ch *CommonHandler) Root(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusNotImplemented)
		return
	}
	w.Write([]byte("Hello World"))
	w.WriteHeader(200)
}

func (ch *CommonHandler) Ping(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusNotImplemented)
		return
	}
	w.Write([]byte("pong"))
	w.WriteHeader(200)
}

func (ch *CommonHandler) Health(msg string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusNotImplemented)
			return
		}
		w.Write([]byte(msg))
		w.WriteHeader(200)
	}
}
