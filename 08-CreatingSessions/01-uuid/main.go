package main

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("session")
	if err != nil {
		id := uuid.New()
		cookie = &http.Cookie{
			Name:  "session",
			Value: id.String(),
			// Secure: true,
			HttpOnly: true, // not accessible with JavaScripte -> more secure
		}
		http.SetCookie(w, cookie)
	}
	fmt.Println(cookie)
}
