package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	//"mango.api/helper"
	//"mango.api/models"
)

type Plan struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title       string             `json:"title,omitempty" bson:"title,omitempty"`
	Description string             `json:"description" bson:"description,omitempty"`
}

var client *mongo.Client

func createplan(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var plan Plan
	json.NewDecoder(r.Body).Decode(&plan)
	collection := client.Database("plans").Collection("smart work")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result, _ := collection.InsertOne(ctx, plan)
	json.NewEncoder(w).Encode(result)
}
func getplan(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	var plan Plan
	collection := client.Database("plans").Collection("smart work")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	err := collection.FindOne(ctx, Plan{ID: id}).Decode(&plan)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(w).Encode(plan)
}
func getallplan(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("content-type", "application/json")
	var plans []Plan
	collection := client.Database("plans").Collection("smart work")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var plan Plan
		cursor.Decode(&plan)
		plans = append(plans, plan)
	}
	if err := cursor.Err();
	err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(w).Encode(plans)
}

func main() {
	fmt.Println("starting the apllication")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, _ = mongo.Connect(ctx, clientOptions)
	r := mux.NewRouter()
	r.HandleFunc("/plan", createplan).Methods("POST")
	r.HandleFunc("/plan/{id}", getplan).Methods("GET")
	r.HandleFunc("/plan",getallplan).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", r))
}
