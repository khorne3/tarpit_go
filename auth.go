package main

import (
	"html/template"
	"net/http"
)

// AUTHUSER AUTHPW
const (
	AUTHUSER = "User"
	AUTHPW   = "1234"
)

var (
	authenticated = false
)

// LoginHandler ... render login if get request, otherwise hand off to perform authentication
func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("view/login.html")
		t.Execute(w, nil)
	} else {
		authCheck(w, r)
	}
}

// LogoutHandler ... /logout
func logoutHandler(w http.ResponseWriter, r *http.Request) {
	authenticated = false
	enforcer(w, r, authenticated)
}

func enforcer(w http.ResponseWriter, r *http.Request, authed bool) {
	if authed == false {
		http.Redirect(w, r, "/login", http.StatusFound)
	}
}

// eventually modify this method to call out to authentication service
func authCheck(w http.ResponseWriter, r *http.Request) {
	user := r.FormValue("username")
	pw := r.FormValue("password")
	if (user == AUTHUSER) && (pw == AUTHPW) {
		authenticated = true
		http.Redirect(w, r, "/app", http.StatusFound)
	} else {
		http.Redirect(w, r, "/login", http.StatusFound)
	}
}
