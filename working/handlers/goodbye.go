package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Goodbye struct {
	l *log.Logger
}

func NewGoodbye(l *log.Logger) *Goodbye{
	return &Goodbye{l: l}
}

func (b *Goodbye) ServeHTTP(w http.ResponseWriter, r *http.Request){
	b.l.Println("Visiting: 'Goodbye'")
	d,err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w,"err",http.StatusBadRequest)
		b.l.Println(err)
		return
	}
	fmt.Fprintf(w,"Bye %s", d)
}

