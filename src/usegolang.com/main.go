package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"usegolang.com/controllers"
	"usegolang.com/models"
)

const (
	host   = "localhost"
	port   = 5432
	user   = "dahmpark"
	dbname = "jamespark_test"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable",
		host, port, user, dbname)
	us, err := models.NewUserService(psqlInfo)
	if err != nil {
		panic(err)
	}
	defer us.Close()
	us.AutoMigrate()

	staticC := controllers.NewStatic()
	usersC := controllers.NewUsers(us)

	r := mux.NewRouter()
	r.Handle("/", staticC.Home).Methods("GET")
	r.Handle("/contact", staticC.Contact).Methods("GET")
	r.Handle("/FAQ", staticC.FAQ).Methods("GET")
	r.HandleFunc("/signup", usersC.New).Methods("GET")
	r.HandleFunc("/signup", usersC.Create).Methods("POST")
	fmt.Println("starting Service")
	http.ListenAndServe(":3000", r)
}
