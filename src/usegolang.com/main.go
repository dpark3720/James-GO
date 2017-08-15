package main

import  (
	"net/http"
	"usegolang.com/controllers"
	"github.com/gorilla/mux"
)

func main() {
	staticC := controllers.NewStatic()
	usersC := controllers.NewUsers()

	r := mux.NewRouter()
	r.Handle("/", staticC.Home) .Methods("GET")
	r.Handle("/contact", staticC.Contact) .Methods("GET")
	r.Handle("/FAQ", staticC.FAQ) .Methods("GET")
	r.HandleFunc("/signup", usersC.New) .Methods("GET")
	r.HandleFunc("/signup", usersC.Create) .Methods("POST")
	http.ListenAndServe(":3000", r)
}