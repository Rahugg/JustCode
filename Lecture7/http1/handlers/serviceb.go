package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type ServiceB struct {
	l *log.Logger
}

func NewServiceB(l *log.Logger) *ServiceB {
	return &ServiceB{
		l: l,
	}
}

func (s *ServiceB) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	data := getDataFromB()
	fmt.Fprintf(rw, data)
}

func getDataFromB() string {
	resp, err := http.Get("http://localhost:8080/getInfo")
	if err != nil {
		return "Failed to fetch from Service A"
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}
