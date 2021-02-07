package controller

import (
	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"
	"rmsServer/model"
	"rmsServer/service"
)

type UserController struct {
	IService service.IUserService
}

func (c UserController) UserLogin(ctx iris.Context)  {
	user := model.User{}
	if err:=ctx.ReadJSON(&user);err !=nil{
		_, _ = ctx.JSON(iris.Map{"code": iris.StatusBadRequest, "msg": "readJson fail"})
	}
	isOk:=c.IService.CheckUserPwd(user)
	if isOk{
		token := jwt.NewTokenWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"UserName": user.UserName,
		})

		// 签名生成jwt字符串
		tokenString, _ := token.SignedString([]byte("My Secret"))
		_, _ = ctx.JSON(iris.Map{"code":iris.StatusOK,"msg":"login success","token":tokenString,"userName":user.UserName})
	}else {
		_, _ = ctx.JSON(iris.Map{"code": iris.StatusBadRequest, "msg": "login fail,userName or pwd is err"})
	}
}

func (c UserController) ShowJwt(ctx iris.Context) {
	_, _ = ctx.JSON(iris.Map{"code": iris.StatusOK, "msg": "ok"})
}
