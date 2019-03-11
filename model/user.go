package model

import (
	"fmt"
	"github.com/Tizeen/go-restful-example/pkg/auth"
	"github.com/Tizeen/go-restful-example/pkg/constvar"
	"gopkg.in/go-playground/validator.v9"
)

type UserModel struct {
	BaseModel
	Username string `json:"username" gorm:"column:username;not null" binding:"required" validate:"min=1,max=32"`
	Password string `json:"password" gorm:"column:password;not null" binding:"required" validate:"min=5,max=128"`
}

func (u *UserModel) TableName() string {
	return "tb_users"
}

func GetUser(username string) (*UserModel, error) {
	u := &UserModel{}
	d := DB.Self.Where("username = ?", username).First(&u)
	return u, d.Error
}

func (u *UserModel) Create() error {
	return DB.Self.Create(&u).Error
}

// 加密密码
func (u *UserModel) Encrypt() (err error) {
	u.Password, err = auth.Encrypt(u.Password)
	return
}

// 验证数据结构是否正确
func (u *UserModel) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}

func ListUser(username string, offset, limit int) ([]*UserModel, uint64, error) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}

	users := make([]*UserModel, 0)
	var count uint64

	where := fmt.Sprintf("username like '%%%s%%'", username)
	if err := DB.Self.Model(&UserModel{}).Where(where).Count(&count).Error; err != nil {
		return users, count, err
	}

	if err := DB.Self.Where(where).Offset(offset).Limit(limit).Order("id desc").Find(&users).Error; err != nil {
		return users, count, err
	}

	return users, count, nil

}

func DeleteUser(id uint64) error {
	user := UserModel{}
	user.BaseModel.Id = id
	return DB.Self.Delete(&user).Error
}

func (u *UserModel) Update() error {
	return DB.Self.Model(u).Updates(UserModel{Username: u.Username, Password: u.Password}).Error
	//return DB.Self.Save(&u).Error
}
