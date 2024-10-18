package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/ZhoraIp/ShelfShare/handlers"
)

func main() {
	l := log.New(os.Stdout, "ZhoraServ: ", log.LstdFlags)

	hh := handlers.NewHello(l)
	gh := handlers.NewGoodbye(l)
	ph := handlers.NewProfile(l)

	sm := http.NewServeMux()

	sm.Handle("/", ph)
	sm.Handle("/hello", hh)
	sm.Handle("/goodbye", gh)

	s := &http.Server{
		Addr:         ":8080",
		Handler:      sm,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, os.Kill)

	sig := <-sigChan
	fmt.Println("Recieved terminate, graceful shutdown!", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}
