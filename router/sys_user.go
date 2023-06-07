package router

import (
	"admin/api"

	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (s *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("user")
	baseApi := api.ApiGroupApp.BaseApi
	{
		// userRouter.POST("adminRegister", baseApi.Register)        //管理员注册
		userRouter.POST("changePassword", baseApi.ChangePassword) //用户修改密码
		userRouter.POST("setUserAuth", baseApi.SetUserAuthority)  // 设置用户权限
		userRouter.POST("deleteUser", baseApi.DeleteUser)         // 删除用户
		userRouter.POST("setUserInfo", baseApi.SetUserInfo)       // 设置用户信息
		userRouter.POST("resetPassword", baseApi.ResetPassword)   //重置密码
		userRouter.POST("getUserList", baseApi.GetUserList)       //分页获取用户列表
		userRouter.POST("getUserInfo", baseApi.GetUserInfo)       //获取用户信息
	}
}
