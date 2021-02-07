package controller

import (
	"github.com/kataras/iris/v12"
	"rmsServer/model"
	"rmsServer/service"
)

type DriveController struct {
	IDriveService service.IDriveService
}
//AddDriveInfo 添加设备
func (dc DriveController)AddDriveInfo(ctx iris.Context){
	driveInfo := model.UserDrive{}
	if err := ctx.ReadJSON(&driveInfo);err !=nil{
		_, _ = ctx.JSON(iris.Map{"code": iris.StatusBadRequest, "msg": "read json fail"})
	}
	isOk,msg:= dc.IDriveService.AddDrive(driveInfo)
	if isOk{
		_, _ = ctx.JSON(iris.Map{"code": iris.StatusOK, "msg": msg})
	}else {
		_, _ = ctx.JSON(iris.Map{"code": iris.StatusBadRequest, "msg": msg})
	}
}
//QueryDrive 查询设备
func (dc DriveController) QueryDrive(ctx iris.Context) {
	userName :=ctx.Params().GetString("userName")
	info,msg:= dc.IDriveService.QueryDriveByUserName(userName)
	_, _ = ctx.JSON(iris.Map{"code": iris.StatusOK, "msg": msg, "data": info})
}
//UpdateDrive 更新设备
func (dc DriveController)UpdateDrive(ctx iris.Context)  {
	drive := model.UserDrive{}
	if err := ctx.ReadJSON(&drive);err !=nil{
		_, _ = ctx.JSON(iris.Map{"code": iris.StatusBadRequest, "msg": "read json fail"})
	}
	isOK,msg:= dc.IDriveService.UpdateDriveByName(drive)
	if isOK{
		ctx.JSON(iris.Map{"code":iris.StatusOK,"msg":msg})
	}else {
		ctx.JSON(iris.Map{"code":iris.StatusBadRequest,"msg":msg})
	}
}
//DelDrive 删除设备
func (dc DriveController) DelDrive(ctx iris.Context)  {
	userName := ctx.Params().GetString("userName")
	isOk,msg:= dc.IDriveService.DelDriveByUserName(userName)
	if isOk{
		ctx.JSON(iris.Map{"code":iris.StatusOK,"msg":msg})
	}else {
		ctx.JSON(iris.Map{"code":iris.StatusBadRequest,"msg":msg})
	}
}

func (dc DriveController) GetAllDrives(ctx iris.Context)  {
	drives,msg:=dc.IDriveService.QueryAllDrive()
	_, _ = ctx.JSON(iris.Map{"code": iris.StatusOK, "msg": msg, "data": drives})
}