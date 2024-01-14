package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Models
type Cource struct {
	CourceID string  `json:"cource_id"`
	Name     string  `json:"name"`
	Price    int     `json:"price"`
	Language string  `json:"language"`
	Author   *Author `json:"author"`
}

type Author struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// Fake DataBase
var cources []Cource

// Middleware and Helper
func (c *Cource) isEmpty() bool {
	return c.CourceID == "" && c.Name == ""
}
func main() {
	fmt.Println("API Routes")

	r := mux.NewRouter()

	cources = append(cources, Cource{
		CourceID: "1",
		Name:     "Go",
		Price:    10,
		Language: "Go",
		Author: &Author{
			Name: "John",
			Age:  30,
		},
	})
	cources = append(cources, Cource{
		CourceID: "2",
		Name:     "Python",
		Price:    20,
		Language: "Python",
		Author: &Author{
			Name: "Jane",
			Age:  25,
		},
	})
	r.HandleFunc("/", serverhome).Methods("GET")
	r.HandleFunc("/cources", getAllCorces).Methods("GET")
	r.HandleFunc("/cource/{id}", getACource).Methods("GET")
	r.HandleFunc("/cource", createACource).Methods("POST")
	r.HandleFunc("/cource/{id}", updateOneCource).Methods("PUT")
	r.HandleFunc("/cource/{id}", deleteCource).Methods("DELETE")
	// Listen Port
	log.Fatal(http.ListenAndServe(":8000", r))
}

// Controllers
//////// Server Home Route

func serverhome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("API is running"))
}

func getAllCorces(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get All Cources")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cources)
}

func getACource(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting a Cource")
	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	params := mux.Vars(r)

	for _, cource := range cources {
		if cource.CourceID == params["id"] {
			json.NewEncoder(w).Encode(cource)
			return
		}

	}
	json.NewEncoder(w).Encode("Cource not found")
	return
}

func createACource(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create a Cource")
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please Send some data")
	}
	var cource Cource
	_ = json.NewDecoder(r.Body).Decode(&cource)
	if cource.isEmpty() {
		json.NewEncoder(w).Encode("Please Send some data")
		return
	}
	cources = append(cources, cource)
	json.NewEncoder(w).Encode(cource)
	return
}


func updateOneCource(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update a Cource")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	// Loop to get it 
	for index, cource := range cources  {
		if(cource.CourceID == params["id"]){
			cources  = append(cources[:index], cources[index+1:]...)
			_=json.NewDecoder(r.Body).Decode(&cource)
			cource.CourceID = params["id"]
			cources = append(cources, cource)
			json.NewEncoder(w).Encode(cource)
			return
		}
	}
	json.NewEncoder(w).Encode("Cource not found")
	return
}


func deleteCource (w http.ResponseWriter, r *http.Request){
	fmt.Println("Delete a Cource")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, cource := range cources {
		if cource.CourceID == params["id"]{
			cources = append(cources[:index],cources[index+1:]...)
			json.NewEncoder(w).Encode(cources )
			return;
			
		}
	}
	json.NewEncoder(w).Encode("Cource not found")
	return
	
}
func deleteCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, course := range cources {
		if course.CourceID == params["id"] {
			cources = append(cources[:index], cources[index+1:]...)
			json.NewEncoder(w).Encode(cources)
			return
		}
	}

	json.NewEncoder(w).Encode("Course not found")
	return
}