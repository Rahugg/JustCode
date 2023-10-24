package main

import (
	"context"
	"http2/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	sm := http.NewServeMux()

	gi := handlers.NewGetInfo(&log.Logger{})
	sa := handlers.NewServiceA(&log.Logger{})

	sm.Handle("/getInfo", gi)
	sm.Handle("/serviceA", sa)

	s := &http.Server{
		Addr:         ":8080",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
	go func() {
		log.Printf("Listening on port%s", s.Addr)
		err := s.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	log.Println("Receieved terminate, graceful shutdown", sig)
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}
