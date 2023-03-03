package common

import (
	"mime/multipart"
	"strings"

	"github.com/fancy-devops/fancy-devops-api/model/db"
)

type File struct {
}

func NewFile() *File {
	return &File{}
}

func (obj *File) GetFileInfo(f *multipart.FileHeader) (size int64, duration int64, t int8, err error) {
	size = f.Size
	duration = 0
	var ct = f.Header.Get("Content-Type")
	if strings.Contains(ct, "audio") {
		t = db.FileType_Audio
	} else if strings.Contains(ct, "video") {
		t = db.FileType_Video
	} else {
		t = db.FileType_Doc
	}
	return size, duration, t, err
}
