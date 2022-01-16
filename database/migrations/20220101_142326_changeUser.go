package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type ChangeUser_20220101_142326 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &ChangeUser_20220101_142326{}
	m.Created = "20220101_142326"

	migration.Register("ChangeUser_20220101_142326", m)
}

// Run the migrations
func (m *ChangeUser_20220101_142326) Up() {
	m.SQL("ALTER TABLE users ADD CONSTRAINT unique_username UNIQUE (username)")

}

// Reverse the migrations
func (m *ChangeUser_20220101_142326) Down() {
	m.SQL("ALTER TABLE users DROP CONSTRAINT unique_username")

}
