package models

import (
	"time"
)

type Task struct {
	BaseModel
	Id          string     `orm:"pk" json:"id"`
	Description *string    `json:"description"`
	ScheduledAt *time.Time `json:"scheduled_at"`
	User        *User      `json:"user" orm:"rel(fk)"`
}

func (u *Task) TableName() string {
	return "tasks"
}
