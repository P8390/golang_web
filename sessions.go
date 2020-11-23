package main

import "net/http"
import "fmt"
import "github.com/gorilla/sessions"

var (
	key   = []byte("Secret-Key")
	store = sessions.NewCookieStore(key)
)

func protected(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		fmt.Fprintln(w, "Forbidden", http.StatusForbidden)
		return
	}
	fmt.Fprintln(w, "Success", http.StatusOK)
}

func login(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")
	session.Values["authenticated"] = true
	session.Save(r, w)
	fmt.Fprintln(w, "Success", http.StatusOK)
}

func logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")
	session.Values["authenticated"] = false
	session.Save(r, w)
	fmt.Fprintln(w, "Success", http.StatusOK)
}

func main() {
	http.HandleFunc("/protected", protected)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)

	http.ListenAndServe(":80", nil)
}
