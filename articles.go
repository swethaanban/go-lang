package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Article struct {
	Id      string `json:"Id"`
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}


var articles []Article

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(articles)
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	for _, article := range articles {
		if article.Id == key {
			json.NewEncoder(w).Encode(article)
		}
	}
}

func Createarticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var article Article
	_ = json.NewDecoder(r.Body).Decode(&article)
	article.Id = strconv.Itoa(rand.Intn(10000))
	articles = append(articles, article)
}

func handleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)
	
	myRouter.HandleFunc("/article", returnAllArticles)
	myRouter.HandleFunc("/article", Createarticle).Methods("POST")
	myRouter.HandleFunc("/article/{id}", returnSingleArticle).Methods("GET")
	log.Fatal(http.ListenAndServe(":10000", myRouter))

}

func main() {
	articles = []Article{
		Article{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		Article{Id: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
		Article{Id: "3", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		Article{Id: "4", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
		Article{Id: "5", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		Article{Id: "6", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
		Article{Id: "7", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		Article{Id: "8", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}

	handleRequests()

}
