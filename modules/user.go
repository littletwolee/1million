package modules

import(
	"1million/models"
	"time"
)

var UserModule *User

type User struct{}  

func init(){
	UserModule = new(User)
}

func (u *User)GetAllUsers(chanr chan []models.User) {
	var users []models.User
	//time.Sleep(10 * time.Second);
	chanr <- append(users, models.User{"bruce", "mima", time.Unix(time.Now().Unix(), 0)})
}

