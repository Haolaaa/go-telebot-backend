package api

import (
	"admin/global"
	"admin/model"
	commonReq "admin/model/common/request"
	"admin/model/common/response"
	"admin/model/request"
	userRes "admin/model/response"
	"admin/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type BaseApi struct{}

func (b *BaseApi) Login(c *gin.Context) {
	var l request.Login
	err := c.ShouldBindJSON(&l)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = utils.Verify(l, utils.LoginVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	u := &model.SysUser{Username: l.Username, Password: l.Password}
	user, err := userService.Login(u)
	if err != nil {
		global.LOG.Error("登陆失败！用户名不存在或者密码错误!", zap.Error(err))
		response.FailWithMessage("用户名不存在或者密码错误", c)
		return
	}
	if user.Enable != 1 {
		global.LOG.Error("登陆失败！用户被禁止登录")
		response.FailWithMessage("用户被禁止登录", c)
		return
	}
	b.TokenNext(c, *user)
	return

}

func (b *BaseApi) TokenNext(c *gin.Context, user model.SysUser) {
	j := &utils.JWT{SigningKey: []byte(global.CONFIG.JWT.SigningKey)}
	claims := j.CreateClaims(request.BaseClaims{
		UUID:        user.UUID,
		ID:          user.ID,
		NickName:    user.NickName,
		Username:    user.Username,
		AuthorityID: user.AuthorityID,
	})
	token, err := j.CreateToken(claims)
	if err != nil {
		global.LOG.Error("获取token失败！", zap.Error(err))
		response.FailWithMessage("获取Token失败", c)
		return
	}

	response.OkWithDetailed(userRes.LoginResponse{
		User:      user,
		Token:     token,
		ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
	}, "登陆成功", c)
	return
}

func (b *BaseApi) Register(c *gin.Context) {
	var r request.Register
	err := c.ShouldBindJSON(&r)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(r, utils.LoginVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	user := &model.SysUser{
		Username:    r.Username,
		NickName:    r.NickName,
		Password:    r.Password,
		HeaderImg:   r.HeaderImg,
		AuthorityID: r.AuthorityID,
		Enable:      r.Enable,
	}
	userReturn, err := userService.Register(*user)
	if err != nil {
		global.LOG.Error("注册失败", zap.Error(err))
		response.FailWithDetailed(userRes.SysUserResponse{User: userReturn}, "注册失败", c)
		return
	}

	response.OkWithDetailed(userRes.SysUserResponse{User: userReturn}, "注册成功", c)
}

func (b *BaseApi) ChangePassword(c *gin.Context) {
	var req request.ChangePassword
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(req, utils.ChangePasswordVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	uid := utils.GetUserID(c)
	u := &model.SysUser{
		MODEL: global.MODEL{ID: uid}, Password: req.Password,
	}
	_, err = userService.ChangePassword(u, req.NewPassword)
	if err != nil {
		global.LOG.Error("修改失败", zap.Error(err))
		response.FailWithMessage("修改失败，原密码与当前账户不符", c)
		return
	}
	response.OkWithMessage("修改成功", c)
}

func (b *BaseApi) GetUserList(c *gin.Context) {
	var pageInfo commonReq.PageInfo
	err := c.ShouldBindJSON(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(pageInfo, utils.PageInfoVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	list, total, err := userService.GetUserInfoList(pageInfo)
	if err != nil {
		global.LOG.Error("获取失败", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

func (b *BaseApi) SetUserAuthority(c *gin.Context) {
	var sua request.SetUserAuthority
	err := c.ShouldBindJSON(&sua)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = userService.SetUserAuthority(sua.ID, sua.AuthorityID)
	if err != nil {
		global.LOG.Error("修改失败", zap.Error(err))
		response.FailWithMessage("修改失败", c)
		return
	}

	response.OkWithMessage("修改成功", c)
}

func (b *BaseApi) DeleteUser(c *gin.Context) {
	var reqID commonReq.GetByID
	err := c.ShouldBindJSON(&reqID)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(reqID, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	jwtID := utils.GetUserID(c)
	if jwtID == uint(reqID.ID) {
		response.FailWithMessage("删除失败", c)
		return
	}
	err = userService.DeleteUser(reqID.ID)
	if err != nil {
		global.LOG.Error("删除失败", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

func (b *BaseApi) SetUserInfo(c *gin.Context) {
	var user request.ChangeUserInfo
	err := c.ShouldBindJSON(&user)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(user, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = userService.SetUserAuthority(user.ID, user.AuthorityID)
	if err != nil {
		global.LOG.Error("设置失败", zap.Error(err))
		response.FailWithMessage("设置失败", c)
		return
	}

	err = userService.SetUserInfo(model.SysUser{
		MODEL: global.MODEL{
			ID: user.ID,
		},
		NickName:  user.NickName,
		HeaderImg: user.HeaderImg,
		Enable:    user.Enable,
	})
	if err != nil {
		global.LOG.Error("设置失败", zap.Error(err))
		response.FailWithMessage("设置失败", c)
		return
	}
	response.OkWithMessage("设置成功", c)
}

func (b *BaseApi) GetUserInfo(c *gin.Context) {
	uuid := utils.GetUserUuid(c)
	reqUser, err := userService.GetUserInfo(uuid)
	if err != nil {
		global.LOG.Error("获取失败", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(
		gin.H{"userInfo": reqUser},
		"获取成功",
		c,
	)
}

func (b *BaseApi) ResetPassword(c *gin.Context) {
	var user model.SysUser
	err := c.ShouldBindJSON(&user)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = userService.ResetPassword(user.ID)
	if err != nil {
		global.LOG.Error("重置失败", zap.Error(err))
		response.FailWithMessage("重置失败"+err.Error(), c)
		return
	}

	response.OkWithMessage("重置成功", c)
}
