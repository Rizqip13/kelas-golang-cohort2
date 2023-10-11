package main

import (
	"fmt"
	"net/http"
	"text/template"
)

var PORT = ":9090"

type User struct {
	Email    string
	Password string
	Address  string
	Reason   string
}

var users = map[string]User{
	"xever@example.com":     {Email: "xever@example.com", Password: "password123", Address: "66209 Surrey Way", Reason: "Re-engineered reciprocal artificial intelligence"},
	"stanfield@example.com": {Email: "stanfield@example.com", Password: "password123", Address: "90648 Londonderry Place", Reason: "Total solution-oriented open system"},
	"jeromy@example.com":    {Email: "jeromy@example.com", Password: "password123", Address: "506 Schlimgen Plaza", Reason: "Cross-platform clear-thinking software"},
}

func main() {
	http.HandleFunc("/", greet)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/login/invalid", invalidHandler)
	http.HandleFunc("/profile", profileHandler)

	fmt.Printf("Server is running on %v...", PORT)
	http.ListenAndServe(PORT, nil)
}

func greet(w http.ResponseWriter, r *http.Request) {
	msg := "Hi Mom!"
	fmt.Fprintf(w, msg)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// Serve the login page
		loginPage := template.Must(template.ParseFiles("login.html"))
		loginPage.Execute(w, users)
	} else if r.Method == http.MethodPost {
		// Process the login form
		email := r.FormValue("email")
		password := r.FormValue("password")

		user, exists := users[email]

		if !exists || user.Password != password {
			// Invalid password or email
			http.Redirect(w, r, "/login/invalid", http.StatusSeeOther)
		}

		http.Redirect(w, r, "/profile?email="+email, http.StatusSeeOther)
	}
}

func profileHandler(w http.ResponseWriter, r *http.Request) {
	profilePage := template.Must(template.ParseFiles("profile.html"))
	email := r.URL.Query().Get("email")
	profilePage.Execute(w, users[email])
}

func invalidHandler(w http.ResponseWriter, r *http.Request) {
	invalidPage := template.Must(template.ParseFiles("invalid.html"))
	invalidPage.Execute(w, nil)
}
