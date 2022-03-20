package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type Initial_20220101_142326 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Initial_20220101_142326{}
	m.Created = "20220101_142326"

	migration.Register("Initial_20220101_142326", m)
}

// Run the migrations
func (m *Initial_20220101_142326) Up() {
	m.SQL(`CREATE TABLE IF NOT EXISTS "users" (
		"created_at" timestamp with time zone NOT NULL,
		"updated_at" timestamp with time zone NOT NULL,
		"archived_at" timestamp with time zone,
		"id" text NOT NULL PRIMARY KEY,
		"first_name" text NOT NULL DEFAULT '' ,
		"last_name" text NOT NULL DEFAULT '' ,
		"username" text NOT NULL DEFAULT ''  UNIQUE,
		"password" text NOT NULL DEFAULT '' ,
		"is_superuser" bool NOT NULL DEFAULT FALSE ,
		"is_staff" bool NOT NULL DEFAULT FALSE
	);`)

	m.SQL(`CREATE TABLE IF NOT EXISTS "tasks" (
		"created_at" timestamp with time zone NOT NULL,
		"updated_at" timestamp with time zone NOT NULL,
		"archived_at" timestamp with time zone,
		"id" text NOT NULL PRIMARY KEY,
		"description" text NOT NULL DEFAULT '' ,
		"scheduled_at" timestamp with time zone NOT NULL,
		"user_id" varchar(255) NOT NULL
	);`)

}

// Reverse the migrations
func (m *Initial_20220101_142326) Down() {
	m.SQL("DROP TABLE tasks")
	m.SQL("DROP TABLE users")

}
