package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/mail"

	"github.com/gorilla/mux"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func validMailAddress(address string) (string, bool) {
	addr, err := mail.ParseAddress(address)
	if err != nil {
		return "", false
	}
	return addr.Address, true
}

func mailChecker(w http.ResponseWriter, r *http.Request) {
	var newEmail string
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	json.Unmarshal(reqBody, &newEmail)
	_, ok := validMailAddress(newEmail)
	if ok {
		fmt.Fprintf(w, "accepted email")
	} else {
		fmt.Fprint(w, "rejected")
	}
	w.WriteHeader(http.StatusCreated)
}

func maxSubstr(w http.ResponseWriter, r *http.Request) {

}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/rest/email", mailChecker)
	router.HandleFunc("/rest/substr/", maxSubstr)
	log.Fatal(http.ListenAndServe(":8080", router))
}
