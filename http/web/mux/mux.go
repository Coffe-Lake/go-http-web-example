package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func productHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	response := fmt.Sprintf("id=%s", id)
	fmt.Fprint(w, response)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Index PAGE")
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/products/{id:[0-9]+}", productHandler)
	router.HandleFunc("/articles/{id:[0-9]+}", productHandler)
	router.HandleFunc("/", indexHandler)
	http.Handle("/", router)

	fmt.Println("Server is listening...")
	http.ListenAndServe(":8181", nil)
}
