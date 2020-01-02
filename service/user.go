package service

import (
	"blog/dao"
	"blog/dto"
	"blog/model"
	"errors"
)

var userDao = dao.User{}

type UserService struct {
}

func (us UserService) List(dto dto.GeneralListDto) (users []model.User, total uint64) {
	return userDao.List(dto)
}

func (us UserService) Create(userDto dto.UserCreateDto) (*model.User, error) {
	// TODO 用户是否存在？

	// password 处理
	userModel := model.User{
		Username: userDto.Username,
		Password: userDto.Password,
	}

	c := userDao.Create(&userModel)
	if c.Error != nil {
		return nil, errors.New("create user failed")
	}

	// TODO 处理role

	return &userModel, nil
}
