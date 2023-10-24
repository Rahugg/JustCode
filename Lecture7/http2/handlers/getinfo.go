package handlers

import (
	"fmt"
	"log"
	"net/http"
)

type GetInfo struct {
	l *log.Logger
}

func NewGetInfo(l *log.Logger) *GetInfo {
	return &GetInfo{
		l: l,
	}
}

func (h *GetInfo) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(rw, "hello world from Service B")
}
