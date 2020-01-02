package model

type User struct {
	Model
	Username string `json:"username"`
	Password string `json:"password"`
}

func (user *User) Save() error {
	return DB.Create(user).Error
}

func ListAccount() (users []*User, err error) {
	err = DB.Find(&users).Error
	return
}

func DeleteUser(id string) bool {
	DB.Where("id = ?", id).Delete(User{})
	return true
}

func UserLogin(username, password string) (user User, success bool) {
	var auth User
	err := DB.Select("*").Where(User{Username: username, Password: password}).First(&auth).Error

	if err == nil {
		return auth, true
	}
	return auth, false
}
