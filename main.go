package main

import {
	"encoding/json"
	"log"
	"net/http"
	"math/rand"
	"strconv"
	"github.com/gorilla/mux"
}

// Order struct
type Order struct {
	Id 						string 	`json:"id"`,
	Quantity 			string 	`json:"quantity"`,
	ProductName 	string	`json:"productName"`
}

// Init Orders
var orders []Order

// -----------
// api methods
// -----------
func getAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(books)
}

func getById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r) // get req params

	for _, item := range orders {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	json.NewEncoder(w).Encode(&Order{})
}

func create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var order Order
	_ = json.NewEncoder(w).Decode(&order)
	order.Id = strconv.Itoa(rand.Intn(100000)) // mock Id - just a random generator
	orders = append(orders, orders)

	json.NewEncoder(w).Encode(book)
}

func update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r) // put req params

	for index, item := range orders {
		if item.Id == params["id"] {
			books = append(books[:index], books[index+1:]...)

			var order Order
			_ = json.NewEncoder(w).Decode(&order)
			order.Id = params["id"]
			orders = append(orders, orders)

			json.NewEncoder(w).Encode(book)
			return
		}
	}
}

func delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r) // delete req params

	for index, item := range orders {
		if item.Id == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode(books)
}

// -------------
// setup methods
// -------------
func handleRequest() {

	// init router
	myRouter := mux.NewRouter().StrictSlash(true)

	// mock data
	books = append(books, Order(Id: "1", Quantity: "2", ProductName: "Tooth Paste"))
	books = append(books, Order(Id: "2", Quantity: "1", ProductName: "Water"))

	// route handlers / endpoints
	myRouter.HandleFunc("/api/orders", getAll).Methods("GET")
	myRouter.HandleFunc("/api/orders/{id}", getById).Methods("GET")
	myRouter.HandleFunc("/api/orders", create).Methods("POST")
	myRouter.HandleFunc("/api/orders/{id}", update).Methods("PUT")
	myRouter.HandleFunc("/api/orders/{id}", delete).Methods("GET")

	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main() {
	handleRequest()
}