package handlers

import (
	"io/ioutil"
	"net/http"
	"time"

	"github.com/chingizof/go-rest-service/redisconn"
)

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
