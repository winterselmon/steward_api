package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http" // import package net-http เข้ามา

	"github.com/gorilla/mux"
)

func HomePage(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "Welcome to the HomePage!")
}

func HandleRequest() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", HomePage)
	myRouter.HandleFunc("/getAddress", getAddressBookAll)
	http.Handle("/", myRouter)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getAddressBookAll(w http.ResponseWriter, _ *http.Request) {
	addBook := addressBook{
		Firstname: "Chaiyarin",
		Lastname:  "Niamsuwan",
		Code:      1993,
		Phone:     "0870940955",
	}
	json.NewEncoder(w).Encode(addBook)
}
