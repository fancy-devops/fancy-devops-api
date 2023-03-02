package responses

type ListDirRes struct {
	BaseRes
	Data []DirModel `json:"data"`
}

type DirModel struct {
	ID       int        `json:"id"`
	Name     string     `json:"name"`
	FullName string     `json:"full_name"`
	Children []DirModel `json:"children"`
}

type ListFileRes struct {
	BaseRes
	Data []FileModel `json:"data"`
}

type FileModel struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	FullName string `json:"full_name"`
	Url      string `json:"url"`
	Size     int64  `json:"size"`
	Duration int64  `json:"duration"`
	Type     int8   `json:"type"`
	Status   int8   `json:"status"`
	IsDir    bool   `json:"is_dir"`
}
