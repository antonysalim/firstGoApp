package main

import (
	"net/http"
	"strconv"
	"text/template"

	"github.com/gorilla/mux"
)

// type users struct {
// 	Username string
// 	Password string
// }

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("template/forms.html"))
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}
		username := r.FormValue("user")
		password := r.FormValue("password")

		insertRow(username, password)

		//fmt.Printf("You have submitted user %v with password %v", username, password)
		tmpl.Execute(w, struct{ Success bool }{true})

		queryDB()

		//fmt.Fprintf(w, "This is main page!")
	})

	r.HandleFunc("/delete", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("template/delForms.html"))
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}
		id := r.FormValue("id")
		convID, _ := strconv.ParseInt(id, 0, 64)

		deleteRow(convID)

		//fmt.Printf("You have submitted user %v with password %v", username, password)
		tmpl.Execute(w, struct{ Success bool }{true})

		queryDB()

		//fmt.Fprintf(w, "This is main page!")
	})

	http.ListenAndServe(":8080", r)
}
