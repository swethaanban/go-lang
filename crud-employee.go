package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
    "github.com/gorilla/mux"
)

type Employee struct {
	Id        string `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var employees []Employee

func allemployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employees)
}

func singleemployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, employee := range employees {
		if employee.Id == params["id"] {
			json.NewEncoder(w).Encode((employee))
			return
		}
	}
}
func createemployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	reqBody, _ := ioutil.ReadAll(r.Body)
    var employee Employee
	json.Unmarshal(reqBody, &employee)
	employees = append(employees, employee)
	json.NewEncoder(w).Encode(employee)

}
func deleteemployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	Id :=params["id"]
	for index, Employee := range employees {
		if Employee.Id == Id {
			employees = append(employees[:index], employees[index+1:]...)
			break
		}
	}
}
func updateemployee(w http.ResponseWriter,r *http.Request){
	vars := mux.Vars(r)
    id := vars["id"]
    var updated Employee
    reqBody, _ := ioutil.ReadAll(r.Body)
    json.Unmarshal(reqBody, &updated)
    for i, Employee := range employees {
	     if Employee.Id == id {
			Employee.Firstname = updated.Firstname
			Employee.Lastname = updated.Lastname
			employees[i] = Employee
			json.NewEncoder(w).Encode(employees)
	}
} 

}

func handleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/employee", allemployees).Methods("GET")
	myRouter.HandleFunc("/employee/{id}", singleemployee).Methods("GET")
	myRouter.HandleFunc("/employee", createemployee).Methods("POST")
	myRouter.HandleFunc("/employe/{id}", deleteemployee).Methods("DELETE")
	myRouter.HandleFunc("/employee/{id}",updateemployee).Methods("PUT")
	log.Fatal(http.ListenAndServe(":10000", myRouter))

}
func main() {
	employees = []Employee{
		Employee{Id: "1", Firstname: "swetha", Lastname: "anban"},
		Employee{Id: "2", Firstname: "arun", Lastname: "k"},
	}
	handleRequests()

}
