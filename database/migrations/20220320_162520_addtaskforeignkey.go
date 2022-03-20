package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type Addtaskforeignkey_20220320_162520 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Addtaskforeignkey_20220320_162520{}
	m.Created = "20220320_162520"

	migration.Register("Addtaskforeignkey_20220320_162520", m)
}

// Run the migrations
func (m *Addtaskforeignkey_20220320_162520) Up() {
	m.SQL("ALTER TABLE tasks ADD CONSTRAINT task_user_fk FOREIGN KEY (user_id) REFERENCES users (id)")
}

// Reverse the migrations
func (m *Addtaskforeignkey_20220320_162520) Down() {
	m.SQL("ALTER TABLE tasks DROP CONSTRAINT task_user_fk")
}
