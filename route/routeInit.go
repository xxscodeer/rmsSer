package route

import (
	"github.com/kataras/iris/v12"
	"rmsServer/controller"
	"rmsServer/middleware"
	"rmsServer/repository"
	"rmsServer/service"
)

func InitRouter()(app *iris.Application)  {
	app = iris.New()
	app.Logger().SetLevel("debug")
	user := controller.UserController{IService: service.UserService{IRepository: repository.UserRepository{}}}
	drive := controller.DriveController{IDriveService: service.DriveService{IDriveRepository: repository.DriveRepository{}}}
	j := middleware.JwtServer()
	userGroup := app.Party("/user",middleware.Cors)
	userGroup.Post("/login",user.UserLogin)

	driveGroup:=app.Party("/api/v1/drive",j.Serve,middleware.Cors)
	driveGroup.Get("/{userName:string}",drive.QueryDrive)
	driveGroup.Get("/all",drive.GetAllDrives)
	driveGroup.Delete("/del/{userName:string}",drive.DelDrive)
	driveGroup.Put("/update",drive.UpdateDrive)
	driveGroup.Post("/add",drive.AddDriveInfo)
	return
}
