package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

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
	fmt.Println("Starting server")
	handleRequest()

}

var err error

//ArWorldPost - Gets the AR world map from the ios device
func ArWorldPost(w http.ResponseWriter, r *http.Request) {
	var arMap WorldMapWrapper
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

	WriteToFile(arMap)

}

// ArWorldGet - Sends ARMap to iOS Device
func ArWorldGet(w http.ResponseWriter, r *http.Request) {
	println("Map get Hit")
	var arMap WorldMapWrapper

	decodeJSONConfig(&arMap, "mapFile.json")

	json.NewEncoder(w).Encode(arMap)
}

// WriteToFile - puts json into the file
func WriteToFile(arMap WorldMapWrapper) {
	f, err := os.Create("mapFile.json")
	PrintFatalError(err)
	defer f.Close()

	err = json.NewEncoder(f).Encode(&arMap)

	PrintFatalError(err)

}

// decodeJSONConfig - takes Data from the json file
func decodeJSONConfig(v interface{}, filename string) {

	fmt.Println("Decoding JSON")
	file, err := os.Open(filename)
	PrintFatalError(err)
	if err != nil {
		// return v, err
	}
	err = json.NewDecoder(file).Decode(&v)
	PrintFatalError(err)
}

// PrintFatalError - Prints an error
func PrintFatalError(err error) {
	if err != nil {
		log.Fatal("Error happened while processing file", err)
	}
}
