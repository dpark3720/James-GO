package controllers

import (
	"fmt"
	"net/http"
	"usegolang.com/views"
)

// New users is used to create a new Users controller.
// This function will panic if the templates are not
// parsed correctly, and should only be used during initial setup.
func NewUsers() *Users {
	return &Users{
	NewView: views.NewView("bootstrap", "users/new"),
	}
}

// New is used to render the form where a user can
// create a new user account
//
// GET /signup
func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	u.NewView.Render(w, nil)
}

// Create is used to process the signup form when a user
// submits it. This is used to create a new user account.
//
// POST /signup
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	form := SignupForm{}
	if err := parseForm(r, &form); err != nil {
		panic(err)
	}
	fmt.Fprintln(w, "Email is", form.Email)
	fmt.Fprintln(w, "Password is", form.Password)
}
type Users struct {
	NewView *views.View
}
type SignupForm struct {
	Email string `schema:"email"`
	Password string `schema:"password"`
}