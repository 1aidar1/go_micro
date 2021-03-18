package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello  {
	return &Hello{l: l}
}

func (h *Hello) ServeHTTP(w http.ResponseWriter, r	 *http.Request){
	h.l.Println("Visiting: 'hello'")
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w,"err",http.StatusBadRequest)
		h.l.Println(err)
		return
	}
	fmt.Fprintf(w,"hello %s",d)
}
