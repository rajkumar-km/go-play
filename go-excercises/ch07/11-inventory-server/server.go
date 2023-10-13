/*
server manages the inventory items

Start the server and run the following commands in another terminal to access the API
$ curl http://localhost:8080/list
$ curl http://localhost:8080/price?item=socks-type1
*/
package main

import (
	"fmt"
	"log"
	"net/http"
)

// database represents a map of item name and its price
type database map[string]int

// list lists the items in inventory with price
func (d *database) list(w http.ResponseWriter, r *http.Request) {
	for item,price := range *d {
		fmt.Fprintf(w, "%q, %d\n", item, price)
	}
}

// price returns the price for the given item
func (d *database) price(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	price, ok := (*d)[item]
	if !ok {
		http.Error(w, "no such item: " + item, http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "%q, %d\n", item, price)
}

func main() {
	db := database{
		"shoe-model1": 700,
		"socks-type1": 100,
	}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)

	addr := "localhost:8080"
	fmt.Printf("Inventory server listening on: %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}