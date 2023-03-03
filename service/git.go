package service

import (
	"github.com/fancy-devops/fancy-devops-api/model/db"
	"github.com/fancy-devops/fancy-devops-api/utils/mysql"
)

type Git struct {
}

func NewGit() *Git {
	return &Git{}
}

func (obj *Git) GetDirectories(where interface{}) (directories []db.Dir) {
	mysql.DBInstance.Where(where).Find(&directories)
	return
}
