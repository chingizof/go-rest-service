package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/chingizof/go-rest-service/handlers"
	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
)

func CounterAdd(w http.ResponseWriter, r *http.Request, client *redis.Client) {
	value, err := client.Get("value").Result()
	if err != nil {
		fmt.Println(err)
	}

	err = client.Set("value", value, 0).Err()

	if err != nil {
		fmt.Println(err)
	}
}

func CounterCheck(w http.ResponseWriter, r *http.Request) {

}

func CounterMinus(w http.ResponseWriter, r *http.Request) {

}

func AddHandler(client *redis.Client, CounterAdd func(w http.ResponseWriter, r *http.Request, client *redis.Client)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { CounterAdd(w, r, client) })
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", handlers.HomeLink)
	router.HandleFunc("/rest/email/check", handlers.MailChecker) //task 2
	router.HandleFunc("/rest/substr/find", handlers.MaxSubstr)   //task1
	router.HandleFunc("/rest/users", handlers.UsersHandler)

	log.Fatal(http.ListenAndServe(":8080", router))
}
