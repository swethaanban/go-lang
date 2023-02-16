package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type person struct {
	Name     string `json:"name"`
	Age      string `json:"age"`
	location string `json:"location"`
}

type persondetail []person

func allpersondetails(w http.ResponseWriter, r *http.Request) {
	persondetail := persondetail{
		person{Name: "arun", Age: "21", location: "bengaluru"},
		person{Name: "karthick", Age: "20", location: "salem"},
	}
	fmt.Println("endpoint hit:all project endpoint")
	json.NewEncoder(w).Encode(persondetail)
}
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "homepage")
}

func handleRequest() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/persondetail", allpersondetails).Methods("GET")
	log.Fatal(http.ListenAndServe(":1010", myRouter))

}

func main() {
	handleRequest()

}
