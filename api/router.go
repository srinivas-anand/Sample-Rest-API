package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Router handle the listner port and call different various methods

func Router() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/getone/{id}", getOne)
	myRouter.HandleFunc("/addone", addOne).Methods("POST")
	myRouter.HandleFunc("/getall", getAll)

	log.Fatal(http.ListenAndServe(":9898", myRouter))

}
