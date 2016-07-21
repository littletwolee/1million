package controllers

import (
	"1million/models"
	"1million/modules"
	"1million/tools"
	"net/http"
//	"log"
//	"github.com/gorilla/mux"
)

var chanr chan []models.User

// Operations about object
type UserController struct {}

func init(){
	chanr = make(chan []models.User)
}

// @Title GetAllUsers
// @Description get all users
// @Success 200 {result} models.Result
// @Failure 500 
// @router /users[GET]
func (u *UserController) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	result := &models.Result{}
	go modules.UserModule.GetAllUsers(chanr)
	// for list := range chanr {
	// 	log.Println(list)
	// 	result.StatusCode = http.StatusOK
	// 	result.Data = list
	// 	tools.RH.GetResult(w, result)
	// }
	for {
		select {
		case list, ok := <- chanr :
			if ok {
				result.StatusCode = http.StatusOK
				result.Data = list
				tools.RH.GetResult(w, result)
			}
			return
		}
	}
}


// func NewWorker(workerPool chan chan models.UserJob) models.UserWorker {
// 	return models.UserWorker{
// 		WorkerPool: workerPool,
// 		JobChannel: make(chan models.UserJob),
// 		Quit:       make(chan bool)}
// }

// // Start method starts the run loop for the worker, listening for a quit channel in
// // case we need to stop it
// func (w models.UserWorker) Start() {
// 	go func() {
// 		for {
// 			// register the current worker into the worker queue.
// 			w.WorkerPool <- w.JobChannel

// 			select {
// 			case job := <-w.JobChannel:
// 				// we have received a work request.
// 				if err := job.Payload.UploadToS3(); err != nil {
// 					log.Errorf("Error uploading to S3: %s", err.Error())
// 				}

// 			case <-w.quit:
// 				// we have received a signal to stop
// 				return
// 			}
// 		}
// 	}()
// }

// // Stop signals the worker to stop listening for work requests.
// func (w models.UserWorker) Stop() {
// 	go func() {
// 		w.Quit <- true
// 	}()
// }
