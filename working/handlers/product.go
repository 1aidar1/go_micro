package handlers

import (
	"goDock/working/data"
	"log"
	"net/http"
)

//define obj
type Products struct {
	l *log.Logger
}

//create handler with logger
func NewProduct(l *log.Logger) *Products {
	return &Products{l: l}
}

func (p *Products) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	// respond to GET
	if r.Method == http.MethodGet {
		p.getProducts(w,r)
		return
	}
	//respond to POST
	if r.Method == http.MethodPost {
		p.addProduct(w,r)
		return
	}
	//catch non supported
	w.WriteHeader(http.StatusMethodNotAllowed)
}


//GET for /products
func (p *Products) getProducts(w http.ResponseWriter, r * http.Request){
	p.l.Println("Vising.. /products | GET")
	lp := data.GetProducts()
	err:= lp.ToJSON(w)
	if err != nil {
		p.l.Println(err)
		http.Error(w,"Error with json",http.StatusInternalServerError)
	}
}

// POST for /products
func (p *Products) addProduct(w http.ResponseWriter, r *http.Request){
	p.l.Println("Vising.. /products | POST")
	product := &data.Product{}
	err:=product.FromJSON(r.Body)
	if err!=nil {
		p.l.Printf("%s", err)
		http.Error(w,"unable to decode json",http.StatusBadRequest)
	}
	p.l.Printf("Product:%#v",product)
}