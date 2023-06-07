package service

import (
	"admin/global"
	"admin/model"
	"admin/model/common/request"
	"admin/utils"
	"errors"
	"time"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type UserService struct{}

func (userService *UserService) Register(u model.SysUser) (userInter model.SysUser, err error) {
	var user model.SysUser
	if !errors.Is(global.DB.Where("username = ?", u.Username).First(&user).Error, gorm.ErrRecordNotFound) {
		return userInter, errors.New("用户名已注册")
	}

	u.Password = utils.BcryptHash(u.Password)
	u.UUID = uuid.Must(uuid.NewV4())
	err = global.DB.Create(&u).Error
	return u, err
}

func (userService *UserService) Login(u *model.SysUser) (userInter *model.SysUser, err error) {
	if nil == global.DB {
		return nil, errors.New("数据库连接失败")
	}

	var user model.SysUser
	err = global.DB.Where("username = ?", u.Username).Preload("Authority").First(&user).Error
	if err == nil {
		if ok := utils.BcryptCheck(u.Password, user.Password); !ok {
			return nil, errors.New("密码错误")
		}
	}

	return &user, err
}

func (userService *UserService) ChangePassword(u *model.SysUser, newPassword string) (userInter *model.SysUser, err error) {
	var user model.SysUser
	if err = global.DB.Where("id = ?", u.ID).First(&user).Error; err != nil {
		return nil, err
	}
	if ok := utils.BcryptCheck(u.Password, user.Password); !ok {
		return nil, errors.New("原密码错误")
	}
	user.Password = utils.BcryptHash(newPassword)
	err = global.DB.Save(&user).Error
	return &user, err
}

func (userService *UserService) GetUserInfoList(info request.PageInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.DB.Model(&model.SysUser{})
	var userList []model.SysUser
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Preload("Authority").Find(&userList).Error
	return userList, total, err
}

func (userService *UserService) SetUserAuthority(id uint, authorityID uint) (err error) {
	return global.DB.Transaction(func(tx *gorm.DB) error {
		userAuthority := &model.SysUserAuthority{
			SysUserID:               id,
			SysAuthorityAuthorityID: authorityID,
		}

		txErr := tx.Model(&model.SysUserAuthority{}).Where("sys_user_id = ?", id).Updates(userAuthority).Error

		txErr = tx.Where("id = ?", id).First(&model.SysUser{}).Update("authority_id", authorityID).Error
		if txErr != nil {
			return txErr
		}

		return nil
	})
}

func (userService *UserService) DeleteUser(id int) (err error) {
	var user model.SysUser
	err = global.DB.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return err
	}
	err = global.DB.Delete(&[]model.SysUserAuthority{}, "sys_user_id = ?", id).Error
	return err
}

func (userService *UserService) SetUserInfo(req model.SysUser) error {
	return global.DB.Model(&model.SysUser{}).Select("updated_at", "nick_name", "header_img", "enable").Where("id = ?", req.ID).Updates(map[string]interface{}{
		"updated_at": time.Now(),
		"nick_name":  req.NickName,
		"header_img": req.HeaderImg,
		"enable":     req.Enable,
	}).Error
}

func (userService *UserService) GetUserInfo(uuid uuid.UUID) (user model.SysUser, err error) {
	var reqUser model.SysUser
	err = global.DB.Preload("Authority").First(&reqUser, "uuid = ?", uuid).Error
	if err != nil {
		return reqUser, err
	}
	return reqUser, err
}

func (userService *UserService) FindUserByID(id int) (user *model.SysUser, err error) {
	var u model.SysUser
	err = global.DB.Where("id = ?", id).First(&u).Error
	return &u, err
}

func (userService *UserService) FindUserByUUID(uuid string) (user *model.SysUser, err error) {
	var u model.SysUser
	if err = global.DB.Where("uuid = ?", uuid).First(&u).Error; err != nil {
		return &u, errors.New("用户不存在")
	}
	return &u, err
}

func (userService *UserService) ResetPassword(ID uint) (err error) {
	err = global.DB.Model(&model.SysUser{}).Where("id = ?", ID).Update("passowrd", utils.BcryptHash("123456")).Error
	return err
}
