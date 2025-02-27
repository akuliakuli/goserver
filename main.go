package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Person struct {
	Name string
	Age int
}

var people []Person

func main(){
	http.HandleFunc("/people", peopleHandler)

	err := http.ListenAndServe("localhost:8080", nil)


	log.Println("server start listening on port 8080")
	if err !=nil{
		log.Fatal(err)
	}
}

func peopleHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method{
	case http.MethodGet:
		getPeople(w, r)
	case http.MethodPost:
		postPerson(w, r)
	
	default:
		http.Error(w, "invalid http method", http.StatusMethodNotAllowed)
	}
}

func getPeople(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(people)
	fmt.Fprint(w, "get '%v", people)

}

func postPerson(w http.ResponseWriter, r *http.Request){
	var person Person
	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	people = append(people, person)

	fmt.Fprint(w, "post new person: '%v", person)
}

