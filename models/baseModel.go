package models

import (
	"time"
)

type BaseModel struct {
	CreatedAt  *time.Time `orm:"auto_now_add" json:"created_at"`
	UpdatedAt  *time.Time `orm:"auto_now" json:"updated_at"`
	ArchivedAt *time.Time `orm:"null" json:"archived_at"`
}
