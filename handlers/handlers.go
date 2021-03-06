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

type Counter struct {
	Value int `json:"value"`
}

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		red := redisconn.GetRedisConnection()
		res, err := red.Get("value").Result()
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
		_, err = red.SetNX("value", body, 60*time.Second).Result()
		if err != nil {
			http.Error(w, "Internal server error", 500)
			return

		}
		IncOne(body)
		w.Write([]byte(body))

	}

}

func IncOne(body []byte) []byte {
	for i := len(body) - 1; i >= 0; i-- {
		if body[i] != '9' {
			body[i] += 1
			break
		} else {
			continue
		}
	}
	return body
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

func IinChecker(w http.ResponseWriter, r *http.Request) {
	var IIN []string
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}
	json.Unmarshal(reqBody, &IIN)
	ans := ""

	for _, v := range IIN {
		if len(v) != 12 {
			fmt.Println("1")
			continue
		} else if v[6] > '6' {
			fmt.Println("2")
			continue
		} else if v[0] > '3' {
			fmt.Println("3")
			continue
		} else if v[0] == '3' && v[1] > '1' {
			fmt.Println("4")
			continue
		} else if v[2] > '1' {
			fmt.Println("5")
			continue
		} else if v[2] == '1' && v[3] > '2' {
			fmt.Println("6")
			continue
		}
		ans += v + " "
	}
	fmt.Fprintf(w, ans)
} //0309175
