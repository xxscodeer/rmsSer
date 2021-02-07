package service

import (
	"fmt"
	"rmsServer/model"
	"rmsServer/repository"
)

type DriveService struct {
	IDriveRepository repository.IDriveRepository
}

type IDriveService interface {
	AddDrive(drive model.UserDrive)(bool,string)
	DelDriveByUserName(userName string)(bool,string)
	UpdateDriveByName(drive model.UserDrive)(bool,string)
	QueryDriveByUserName(userName string)(model.UserDrive,string)
	QueryAllDrive()([]model.UserDrive,string)
}

func (s DriveService)AddDrive(drive model.UserDrive)(bool,string) {
	if err :=s.IDriveRepository.CreateDrive(drive);err !=nil{
		return false,"添加失败"
	}
	return true,"添加成功"
}

func (s DriveService) DelDriveByUserName(userName string)(bool,string)  {
	if err := s.IDriveRepository.DelDriveByUserName(userName);err !=nil{
		return false,fmt.Sprintln("删除失败",err)
	}
	return true,"删除成功"
}

func (s DriveService) UpdateDriveByName(drive model.UserDrive)(bool,string)  {
	if err :=s.IDriveRepository.UpdateDriveById(drive);err !=nil{
		return false,fmt.Sprintln("删除失败",err)
	}
	return true,"更新成功"
}

func (s DriveService) QueryDriveByUserName(userName string)(model.UserDrive,string) {
	info,err:=s.IDriveRepository.FindDriveByUserName(userName)
	if err !=nil{
		return model.UserDrive{},fmt.Sprintln("查询失败",err)
	}
	return info,"查询成功"
}

func (s DriveService) QueryAllDrive()([]model.UserDrive,string)  {
	drive,err:=s.IDriveRepository.FindAllDrive()
	if err !=nil{
		return nil,fmt.Sprintln("查询失败",err)
	}
	return drive,"查询成功"
}