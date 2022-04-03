package main

import (
	"log"
	"net/http"
	"math/rand"
	"strconv"
	"encoding/json"
	"github.com/gorilla/mux"
)

type Order struct {
	ID string `json: "id"`
	Title string `json: "title"`
	User *User `json: "user"`
}

type User struct {
	Firstname string `json: "firstname"`
	Lastname string	`json: "lastname"`
}

var orders []Order

func getOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)
}

func getOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range orders{
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Order{})
}

func createOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var order Order
	_= json.NewDecoder(r.Body).Decode(&order)
	order.ID = strconv.Itoa(rand.Intn(100000))
	orders = append(orders, order)
	json.NewEncoder(w).Encode(order)
}

func updateOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range orders {
		if item.ID  == params["id"] {
			orders = append(orders[:index], orders[index+1:]...)
			var order Order
			_= json.NewDecoder(r.Body).Decode(&order)
			order.ID = params["id"]
			orders = append(orders, order)
			json.NewEncoder(w).Encode(order)
			return
		}
	}
	json.NewEncoder(w).Encode(orders)
}

func deleteOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range orders {
		if item.ID == params["id"] {
			orders = append(orders[:index], orders[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(orders)
}

func main() {
	r := mux.NewRouter()
	orders = append(orders, Order{ID: "1", Title: "Smth to eat", User: &User{Firstname: "Ivan", Lastname: "Ivanov"}})
	orders = append(orders, Order{ID: "2", Title: "Smth to wear", User: &User{Firstname: "John", Lastname: "Doe"}})
	r.HandleFunc("/orders", getOrders).Methods("GET")
	r.HandleFunc("/orders/{id}", getOrder).Methods("GET")
	r.HandleFunc("/orders", createOrder).Methods("POST")
	r.HandleFunc("/orders/{id}", updateOrder).Methods("PUT")
	r.HandleFunc("/orders/{id}", deleteOrder).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", r))
}