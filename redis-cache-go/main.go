package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"redis-cache-go/dbredis"
	"time"

	"github.com/gorilla/mux"
)

func init() {
	dbredis.RedisStart()
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/cache/{id}", func(w http.ResponseWriter, r *http.Request) {

		id := mux.Vars(r)["id"]

		renderJSON := func(w http.ResponseWriter, val interface{}, statusCode int) {
			w.WriteHeader(statusCode)
			_ = json.NewEncoder(w).Encode(&val)
		}

		if !dbredis.Redis.CheckExistKey(id) {
			time.Sleep(5 * time.Second)

			n := 1000000
			total := 0

			for i := 0; i < n; i++ {
				total++
			}

			dbredis.Redis.SetCache(id, total, time.Minute)

			renderJSON(w, &total, http.StatusOK)

			return
		}

		var value interface{}

		dbredis.Redis.GetCache(id, &value)

		renderJSON(w, &value, http.StatusOK)

		return

	})

	srv := &http.Server{
		Handler:      router,
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("Starting server :8080")

	log.Fatal(srv.ListenAndServe())
}
