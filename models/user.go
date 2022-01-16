package models

import (
	"docket-beego/utils"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	BaseModel
	Id          string  `orm:"pk" json:"id"`
	FirstName   *string `json:"first_name"`
	LastName    *string `json:"last_name"`
	Username    *string `json:"username"`
	Password    *string `json:"-"`
	IsSuperuser *bool   `json:"-"`
	IsStaff     *bool   `json:"-"`
}

func (u *User) TableName() string {
	return "users"
}

func (u *User) TableUnique() [][]string {
	return [][]string{
		{"Username"},
	}
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
