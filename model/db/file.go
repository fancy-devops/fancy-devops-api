package db

const (
	FileType_Doc   = 0
	FileType_Audio = 1
	FileType_Video = 2

	FileStatus_Default = 0
	FileStatus_Deleted = 10
)

type File struct {
	BaseModel

	Name     string `json:"name"`
	Host     string `json:"host"`
	Path     string `json:"path"`
	Size     int64  `json:"size"`
	Duration int64  `json:"duration"`
	Type     int8   `json:"type"`
	Status   int8   `json:"status"`
	DirId    int    `json:"dir_id"`
}

type Dir struct {
	BaseModel

	Name     string `json:"name"`
	FullName string `json:"full_name"`
	ParentId int    `json:"parent_id"`
}
