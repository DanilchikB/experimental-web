package web

import (
	"fmt"
	"net/http"

	"github.com/danilbushkov/experimental-web/web/requests"
)

//Server run server
func Server() {

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)
	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "About Page")
	})
	http.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Contact Page")
	})
	requests.UserRequests()

	http.ListenAndServe("localhost:5000", nil)

}
