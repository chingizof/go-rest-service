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

func HomeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func ValidMailAddress(address string) (string, bool) {
	addr, err := mail.ParseAddress(address)
	if err != nil {
		return "", false
	}
	return addr.Address, true
}

func MailChecker(w http.ResponseWriter, r *http.Request) {
	var newEmail []string
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	json.Unmarshal(reqBody, &newEmail)
	for _, v := range newEmail {
		_, ok := ValidMailAddress(v)
		if ok {
			v += " "
			fmt.Fprintf(w, v)
		}
	}
	w.WriteHeader(http.StatusCreated)
}

func LongestSubstr(str string) string {
	str_len := len(str)

	arr := make([]bool, 125)

	max := 1
	maxstr := ""
	count := 0
	curr := ""

	// If string length is zero
	if str_len == 0 {

		max = 0

	} else {
		for i := 0; i < str_len; i++ {
			curr = ""
			for j := 0; j < 125; j++ {
				arr[j] = false
			}

			arr[str[i]] = true
			curr += string(str[i])
			count = 1

			for k := i + 1; k < str_len; k++ {

				if arr[str[k]] {
					break
				}

				arr[str[k]] = true
				count = count + 1
				curr += string(str[k])

				if max < count {
					max = count
					maxstr = curr
				}

			}
		}
	}
	return maxstr
}

func MaxSubstr(w http.ResponseWriter, r *http.Request) {
	var word string

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	json.Unmarshal(reqBody, &word)
	fmt.Println(word)
	v := LongestSubstr(word)
	fmt.Fprintf(w, v)
	w.WriteHeader(http.StatusCreated)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", HomeLink)
	router.HandleFunc("/rest/email/check", MailChecker) //task 2
	router.HandleFunc("/rest/substr/find", MaxSubstr)   //task1
	log.Fatal(http.ListenAndServe(":8080", router))
}
