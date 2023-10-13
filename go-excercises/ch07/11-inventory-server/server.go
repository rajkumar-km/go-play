/*
server manages the inventory items
Start the server and run the following commands in another terminal to access the API

	curl "http://localhost:8080/create?item=shoe-model2&price=800"
	curl "http://localhost:8080/list"
	curl "http://localhost:8080/price?item=socks-type1"
	curl "http://localhost:8080/update?item=shoe-model1&price=600"
	curl "http://localhost:8080/list"
	curl "http://localhost:8080/delete?item=shoe-model1"
	curl "http://localhost:8080/list"
*/
package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// database represents a map of item name and its price
type database map[string]int

// list lists the items in inventory with price
func (d *database) list(w http.ResponseWriter, r *http.Request) {
	for item, price := range *d {
		fmt.Fprintf(w, "%q, %d\n", item, price)
	}
}

// price returns the price for the given item
func (d *database) price(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	price, ok := (*d)[item]
	if !ok {
		http.Error(w, "no such item: "+item, http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "%q, %d\n", item, price)
}

// create allows to add a new item in inventory
func (d *database) create(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	if !q.Has("item") {
		http.Error(w, "item not provided", http.StatusBadRequest)
		return
	}
	if !q.Has("price") {
		http.Error(w, "price not provided", http.StatusBadRequest)
		return
	}

	item, priceStr := q.Get("item"), q.Get("price")
	_, ok := (*d)[item]
	if ok {
		http.Error(w, "item already exists: "+item, http.StatusConflict)
		return
	}

	price, err := strconv.Atoi(priceStr)
	if err != nil {
		http.Error(w, "invalid price: "+priceStr, http.StatusBadRequest)
		return
	}

	(*d)[item] = price
	w.WriteHeader(http.StatusCreated)
}

// update sets the new price for a given item
func (d *database) update(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	if !q.Has("item") {
		http.Error(w, "item not provided", http.StatusBadRequest)
		return
	}
	if !q.Has("price") {
		http.Error(w, "price not provided", http.StatusBadRequest)
		return
	}

	item, priceStr := q.Get("item"), q.Get("price")
	_, ok := (*d)[item]
	if !ok {
		http.Error(w, "item not found: "+item, http.StatusNotFound)
		return
	}

	price, err := strconv.Atoi(priceStr)
	if err != nil {
		http.Error(w, "invalid price: "+priceStr, http.StatusBadRequest)
		return
	}

	(*d)[item] = price
	w.WriteHeader(http.StatusOK)
}

// delete deletes an item from inventory
func (d *database) delete(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	if !q.Has("item") {
		http.Error(w, "item not provided", http.StatusBadRequest)
		return
	}

	item := q.Get("item")
	_, ok := (*d)[item]
	if !ok {
		http.Error(w, "item not found: "+item, http.StatusNotFound)
		return
	}

	delete(*d, item)
	w.WriteHeader(http.StatusOK)
}

func main() {
	// Initial inventory
	db := database{
		"shoe-model1": 700,
		"socks-type1": 100,
	}

	// Register http handlers
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/create", db.create)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/delete", db.delete)

	// Start the server
	addr := "localhost:8080"
	fmt.Printf("Inventory server listening on: %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
