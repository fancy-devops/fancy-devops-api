package db

type Collection struct {
	BaseModel

	Name string `json:"name"`
}

type Collection_File struct {
	BaseModel

	FileId       int `json:"file_id"`
	CollectionId int `json:"collection_id"`
}

type User_Collection struct {
	BaseModel

	UserId       int `json:"user_id"`
	CollectionId int `json:"collection_id"`
}
