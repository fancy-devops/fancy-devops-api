package service

import (
	"crypto/md5"
	"errors"
	"fmt"

	"gitlab.chad122.top/fancy-devops/fancy-devops-api/database"
	"gitlab.chad122.top/fancy-devops/fancy-devops-api/model/db"
	"gitlab.chad122.top/fancy-devops/fancy-devops-api/pkg/util"
)

type User struct {
}

func NewUser() *User {
	return &User{}
}

func (obj *User) GetUsers(skip int, size int, where interface{}) (users []db.User) {
	database.DBInstance.Where(where).Offset(skip).Limit(size).Find(&users)
	return
}

func (obj *User) GetUserTotal(where interface{}) (count int64) {
	database.DBInstance.Model(&db.User{}).Where(where).Count(&count)
	return
}

func (obj *User) GetUser(id int) (user db.User) {
	database.DBInstance.First(&user, id)
	return
}

func (obj *User) GetUserByEmail(email string) (user db.User) {
	database.DBInstance.First(&user, "email = ?", email)
	return
}

func (obj *User) GetUserByPassword(email string, pwd string) (user db.User) {
	md5Pwd := fmt.Sprintf("%x", md5.Sum([]byte(pwd)))
	database.DBInstance.First(&user, "email = ? and password = ?", email, md5Pwd)
	return
}

func (obj *User) CreateUser(name string, email string, pwd string, role int) int {
	md5Pwd := fmt.Sprintf("%x", md5.Sum([]byte(pwd)))
	newUser := db.User{
		Name:     name,
		Email:    email,
		Password: md5Pwd,
		Secret:   util.NewGoogleAuth().GetSecret(),
		Role:     role,
	}
	database.DBInstance.Create(&newUser)
	return newUser.ID
}

func (obj *User) UpdateUserPwd(email string, pwd string) {
	md5Pwd := fmt.Sprintf("%x", md5.Sum([]byte(pwd)))
	database.DBInstance.Model(db.User{}).Where("email = ?", email).UpdateColumns(db.User{Password: md5Pwd})
}

func (obj *User) UpdateUserSecret(email string) {
	secret := util.NewGoogleAuth().GetSecret()
	database.DBInstance.Model(db.User{}).Where("email = ?", email).UpdateColumns(db.User{Secret: secret})
}

func (obj *User) SendSecretEmail(email string) error {
	user := obj.GetUserByEmail(email)
	if user.ID == 0 {
		return errors.New("email not exist")
	}
	if user.Secret == "" {
		obj.UpdateUserSecret(email)
		user = obj.GetUserByEmail(email)
	}
	util.NewEmail().SendEmail(user.Email, "【fancy-devops】Google Authenticator", "您的Secret为："+user.Secret+"<br />二维码地址："+util.NewGoogleAuth().GetQrcodeUrl(user.Email, user.Secret))
	return nil
}
