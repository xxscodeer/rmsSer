package repository

import (
	"rmsServer/model"
	"rmsServer/tools"
)

type DriveRepository struct {

}

type IDriveRepository interface {
	CreateDrive(drive model.UserDrive)error
	FindDriveByUserName(userName string)(model.UserDrive,error)
	DelDriveByUserName(userName string)error
	UpdateDriveById(drive model.UserDrive)error
	FindAllDrive()([]model.UserDrive,error)
}

func (r DriveRepository) CreateDrive(drive model.UserDrive)error {
	return tools.DbEngine.Create(&drive).Error
}

func (r DriveRepository) FindDriveByUserName(userName string)(userDrive model.UserDrive,err error){
	return userDrive,tools.DbEngine.Where("user_name=?",userName).First(&userDrive).Error
}

func (r DriveRepository) DelDriveByUserName(userName string)error{
	return tools.DbEngine.Where("user_name=?",userName).Delete(&model.UserDrive{}).Error
}

func (r DriveRepository) UpdateDriveById(drive model.UserDrive)error {
	return tools.DbEngine.Where("user_name=?",drive.UserName).Updates(drive).Error
}

func (r DriveRepository)FindAllDrive()(drives []model.UserDrive,err error) {
	return drives,tools.DbEngine.Find(&drives).Error
}