package main

import (
	"context"
	"goDock/working/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main(){
	//create log
	l:=log.New(os.Stdout, "product-api", log.LstdFlags)

	//create handlers
	hh:=handlers.NewHello(l)
	bh:=handlers.NewGoodbye(l)
	ph:=handlers.NewProduct(l)

	//assign mux
	sm := http.NewServeMux()
	sm.Handle("/",hh)
	sm.Handle("/bye",bh)
	sm.Handle("/products",ph)

	//create server
	s:=&http.Server{
		Addr: ":8001",
		Handler: sm,
		IdleTimeout: 128*time.Second,
		ReadTimeout: 1*time.Second,
		WriteTimeout: 1*time.Second,
	}
	l.Printf("Starting on %s",s.Addr)
	//start server on goroutine
	go func() {

		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	//listen when server dies and kill gracefully
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan,os.Kill)

	sig := <-sigChan
	l.Println("Recieved terminate, graceful shutdown", sig)
	tc, _:=context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}
