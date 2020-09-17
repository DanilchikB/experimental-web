package requests

import (
	"fmt"
	"net/http"

	"github.com/DanilchikB/helper-for-university/conf"
	"github.com/DanilchikB/helper-for-university/web/user"
)

//UserRequests add Handle for user
func UserRequests() {
	http.HandleFunc("/auth", func(w http.ResponseWriter, r *http.Request) {
		user.Login(w, r)
		http.Redirect(w, r, "/userpage", http.StatusSeeOther)
	})

	http.HandleFunc("/auth1", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		fmt.Fprintln(w, r.PostForm)
	})

	http.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) {
		user.Logout(w, r)
		http.Redirect(w, r, "/userpage", http.StatusSeeOther)
	})

	http.HandleFunc("/userpage", func(w http.ResponseWriter, r *http.Request) {
		session, _ := conf.Store.Get(r, "user")

		var id int
		// Check if user is authenticated
		if auth := session.Values["auth"].(bool); !auth {
			fmt.Fprintln(w, "We are not auth")
			return
		}
		id = session.Values["id"].(int)

		// Print secret message
		fmt.Fprintf(w, "You are user nomber %v!", id)
	})
}
