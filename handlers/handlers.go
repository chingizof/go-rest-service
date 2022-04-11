package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/mail"
	"time"

	"github.com/chingizof/go-rest-service/db"
	"github.com/chingizof/go-rest-service/redisconn"
)

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		red := redisconn.GetRedisConnection()
		res, err := red.Get("user").Result()
		if err != nil {
			http.Error(w, "Error!", 404)
			return

		}
		w.Write([]byte(res))
	case "POST":
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Internal server error", 500)
			return
		}
		defer r.Body.Close()

		red := redisconn.GetRedisConnection()
		_, err = red.SetNX("user", body, 60*time.Second).Result()
		if err != nil {
			http.Error(w, "Internal server error", 500)
			return

		}
		w.Write([]byte(body))

	}

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

func HomeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "I love TSARKA!")
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

func AddUser(w http.ResponseWriter, r *http.Request) {
	forms := []string{}
	db := db.SqlConnect()
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	json.Unmarshal(reqBody, &forms)
	name, surname := forms[0], forms[1]

	insert, err := db.Query("INSERT INTO user VALUES(0, '" + name + "', '" + surname + "');")

	if err != nil {
		fmt.Println(err)
	}

	defer insert.Close()
}
