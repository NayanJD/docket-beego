package models

import (
	"docket-beego/utils"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	BaseModel
	Id          string  `orm:"pk" json:"id"`
	FirstName   *string `json:"first_name,omitempty"`
	LastName    *string `json:"last_name,omitempty"`
	Username    *string `orm:"unique" json:"username,omitempty"`
	Password    *string `json:"-"`
	IsSuperuser *bool   `json:"-"`
	IsStaff     *bool   `json:"-"`
	Tasks       []*Task `orm:"reverse(many)" json:"tasks,omitempty"`
}

func (u User) String() string {
	return fmt.Sprintf("User: %v %v", *u.FirstName, *u.LastName)
}

func (u *User) TableName() string {
	return "users"
}

func (u *User) SetPassword(pwd *string) error {
	hash, err := bcrypt.GenerateFromPassword(
		[]byte(*u.Password),
		bcrypt.MinCost,
	)

	if err != nil {
		return err
	}

	*u.Password = string(hash)

	return nil
}

func (u *User) ComparePassword(pwd *string) bool {
	err := bcrypt.CompareHashAndPassword(
		[]byte(*u.Password),
		[]byte(*pwd),
	)

	if err != nil {
		utils.Log.Error("err", err)
		return false
	}

	return true
}
