package service

import (
	"gitlab.chad122.top/fancy-devops/fancy-devops-api/database"
	"gitlab.chad122.top/fancy-devops/fancy-devops-api/model/db"
)

type Git struct {
}

func NewGit() *Git {
	return &Git{}
}

func (obj *Git) GetDirectories(where interface{}) (directories []db.Dir) {
	database.DBInstance.Where(where).Find(&directories)
	return
}
