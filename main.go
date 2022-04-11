package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/chingizof/go-rest-service/handlers"
	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
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

func AddUserHandler(db *sql.DB, AddUser func(w http.ResponseWriter, r *http.Request, db *sql.DB)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { AddUser(w, r, db) })
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", handlers.HomeLink)
	router.HandleFunc("/rest/email/check", handlers.MailChecker) //task 2
	router.HandleFunc("/rest/iin/check", handlers.IinChecker)    //task 2 harder
	router.HandleFunc("/rest/substr/find", handlers.MaxSubstr)   //task1
	router.HandleFunc("/rest/redisusers", handlers.UsersHandler) //task 3 partly
	router.HandleFunc("/rest/user", handlers.AddUser)            //task 4, add in format ["name", "surname"]

	log.Fatal(http.ListenAndServe(":8080", router))
}
