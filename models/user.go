package models

import (
	"time"
)

type User struct {
	Name         string
	Pwd          string
	CreateDate   time.Time
}

// func (u *User)setObjectId(objectId ){
	
// }