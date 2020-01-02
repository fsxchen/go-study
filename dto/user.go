package dto

type UserCreateDto struct {
	Username string `form:"username" json:"username" valid:"Required;MaxSize(100)"`
	Password string `form:"password" json:"password" valid:"Required;MaxSize(100)"`
}
