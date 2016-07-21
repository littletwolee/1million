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
	chanr = make(chan []models.User, 10)
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
		case list, _ := <- chanr :
			result.StatusCode = http.StatusOK
			result.Data = list
			tools.RH.GetResult(w, result)
			return
		}
	}
}
