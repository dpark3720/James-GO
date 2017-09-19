package main

import (
	"fmt"

	"usegolang.com/models"

	_ "github.com/jinzhu/gorm/dialects/postgres"
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
	us.DestructiveReset()
	user := models.User{
		Name:  "James Park",
		Email: "james.park@capitalone.com",
	}
	if err := us.Create(&user); err != nil {
		panic(err)
	}
	if err := us.Delete(user.ID); err != nil {
		panic(err)
	}
	// user.Email = "dpark@capitalone.com"
	// if err := us.Update(&user); err != nil {
	// 	panic(err)
	// }
	// userByEmail, err := us.ByEmail("dpark@capitalone.com")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(userByEmail)

	userByID, err := us.ByID(user.ID)
	if err != nil {
		panic(err)
	}
	fmt.Println(userByID)
}
