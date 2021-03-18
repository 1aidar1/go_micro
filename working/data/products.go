package data

import (
	"encoding/json"
	"io"
)

type Product struct {
	ID uint `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Price float32 `json:"price"`
	SKU string `json:"sku"`
	//CreatedOn string `json:"-"`
	//UpdatedOn string `json:"-"`
	//DeletedOn string `json:"-"`
}

//read json obj
func (p *Product) FromJSON(r io.Reader) error{
	e:=json.NewDecoder(r)
	return e.Decode(p)
}

type Products []*Product

//list to json
func (p *Products) ToJSON(w io.Writer) error  {
	e:=json.NewEncoder(w)
	return e.Encode(p)
}

func GetProducts() Products{
	return productList
}

var productList = Products{
	&Product{
		1,
		"Latte",
		"Milky Coffee",
		2.45,
		"abc234",
		//time.Now().UTC().String(),
		//time.Now().UTC().String(),
		//"",//
	},
	&Product{
		2,
		"Espresso",
		"String coffee without milk",
		1.99,
		"jid567",
		//time.Now().UTC().String(),
		//time.Now().UTC().String(),
		//"",
	},
}
