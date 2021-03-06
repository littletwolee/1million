// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"1million/controllers"
	"github.com/gorilla/mux"
	"net/http"
//	"log"
)

func init() {
	r := mux.NewRouter()
	userController := new(controllers.UserController)
	user := "/users"
	r.HandleFunc(user, userController.GetAllUsers).
		Methods("GET")
	http.Handle("/", r)
}
