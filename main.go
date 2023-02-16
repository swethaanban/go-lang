package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type project struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}
type projects []project

func allprojects(w http.ResponseWriter, r *http.Request) {
	projects := projects{
		project{Title: "gas", Content: "fastest"},
		project{Title: "car resale value", Content: "prediction"},
	}
	fmt.Println("endpoint hit:all project endpoint")
	json.NewEncoder(w).Encode(projects)
}
func testPostprojects(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "test post endpoints worked")
}
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "homepage")

}
func handleRequest() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/projects", allprojects).Methods("GET")
	myRouter.HandleFunc("/projects", testPostprojects).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", myRouter))
}
func main() {
	handleRequest()

}
