package main

import (
	"flag"
	"fmt"

	"docket-beego/models"

	"github.com/google/uuid"
)

func CreateSuperuser(flagSet *flag.FlagSet, commands []string) {
	firstName := flagSet.String("first-name", "first_name", "First name of user")
	lastName := flagSet.String("last-name", "last_name", "Last name of user")
	username := flagSet.String("username", "username@user.com", "Username (email) of user")
	password := flagSet.String("password", "password", "Password of account")
	isSuperuser := true
	isStaff := false
	flagSet.Parse(commands)

	newId := uuid.New().String()
	newUser := models.User{
		Id:          newId,
		FirstName:   firstName,
		LastName:    lastName,
		Username:    username,
		Password:    password,
		IsSuperuser: &isSuperuser,
		IsStaff:     &isStaff,
	}

	_, err := models.Orm.Insert(&newUser)

	if err != nil {
		fmt.Println(err)
	}

	newUser.SetPassword(password)

	_, err = models.Orm.Update(&newUser)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("New superuser created")
}
