package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type WorldMapWrapper struct {
	WorldMap string `json:"worldMap"`
}

func handleRequest() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/arWorld", ArWorldPost).Methods("POST")
	myRouter.HandleFunc("/arWorld", ArWorldGet).Methods("GET")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main() {
	handleRequest()

	// defer db.Close()
}

var err error
var arMap WorldMapWrapper

func ArWorldPost(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hit")
	if err != nil {
		panic("Error in NewUser")
	}
	fmt.Println("Hit")
	w.Header().Set("Content-Type", "application/json")

	//decodes the user from the body and turns it into data
	if r.Body == nil {
		fmt.Println("Body is nil")
	}

	json.NewDecoder(r.Body).Decode(&arMap)

	b := []byte(arMap.WorldMap)
	println("bytes", len(b))
	fmt.Println(len(arMap.WorldMap))

}

func ArWorldGet(w http.ResponseWriter, r *http.Request) {
	// w http.ResponseWriter, r *http.Request
	// print
	println("Map get Hit")
	b := []byte(arMap.WorldMap)
	println("bytes", len(b))
	json.NewEncoder(w).Encode(arMap)
}
