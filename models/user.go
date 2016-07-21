package models

import (
	"time"
)

type UserJob struct {
	User User
}

var UserJobQueue chan UserJob

// Worker represents the worker that executes the job
type UserWorker struct {
	WorkerPool  chan chan UserJob
	JobChannel  chan UserJob
	Quit        chan bool
}

type User struct {
	Name         string
	Pwd          string
	CreateDate   time.Time
}

// func (u *User)setObjectId(objectId ){
	
// }
