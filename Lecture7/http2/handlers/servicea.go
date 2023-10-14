package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type ServiceA struct {
	l *log.Logger
}

func NewServiceA(l *log.Logger) *ServiceA {
	return &ServiceA{
		l: l,
	}
}

func (s *ServiceA) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	data := getDataFromA()
	fmt.Fprintf(rw, data)
}

func getDataFromA() string {
	resp, err := http.Get("http://localhost:8081/getInfo")
	if err != nil {
		return "Failed to fetch from Service A"
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}
