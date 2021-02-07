package service

import (
	"rmsServer/model"
	"rmsServer/repository"
)

type UserService struct {
	IRepository repository.IUserRepository
}

type IUserService interface {
	CheckUserPwd(user model.User)bool
}
//CheckUserPwd loginCheck
func (s UserService) CheckUserPwd(user model.User)bool{
	userInfo,err:=s.IRepository.FindUserPwdByName(user.UserName)
	if err !=nil{
		return false}
	if userInfo.UserPwd == user.UserPwd{
		return true
	}
	return false
}