package repository

import (
	"rmsServer/model"
	"rmsServer/tools"
)

type UserRepository struct {

}

type IUserRepository interface {
	FindUserPwdByName(userName string)(model.User,error)
}

func (r UserRepository)FindUserPwdByName(userName string)(user model.User, err error) {
	return user,tools.DbEngine.Where("user_name=?",userName).Find(&user).Error
}
