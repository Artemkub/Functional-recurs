// Контекст: веб-обработчик с таймаутом 50 мс и значением в контексте
package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type key int

const userIPKey key = 0

type User struct {
	Name  string
	Email string
}

func main() {
	routes()

	log.Println("listener : Started : Listening on: http://localhost:4000")
	http.ListenAndServe(":4000", nil)
}

func routes() {
	http.HandleFunc("/user", findUser)
}

func findUser(rw http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(r.Context(), 50*time.Millisecond)
	defer cancel()

	ctx = context.WithValue(ctx, userIPKey, r.RemoteAddr)

	ch := make(chan *User, 1)

	go func() {
		uip, ok := ctx.Value(userIPKey).(string)
		if ok {
			log.Println("User IP:", uip)
		}

		ch <- readDatabase()

		log.Println("goroutine : Terminating")
	}()

	select {
	case u := <-ch:
		sendResponse(rw, u, http.StatusOK)

		log.Println("Sent StatusOK")
		return

	case <-ctx.Done():
		e := struct{ Error string }{ctx.Err().Error()}

		sendResponse(rw, e, http.StatusRequestTimeout)

		log.Println("Sent StatusRequestTimeout")
		return
	}
}

func readDatabase() *User {
	u := User{
		Name:  "Bill",
		Email: "bill@ardanlabs.com",
	}

	time.Sleep(100 * time.Millisecond)

	return &u
}

func sendResponse(rw http.ResponseWriter, v interface{}, statusCode int) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(statusCode)
	json.NewEncoder(rw).Encode(v)
}
