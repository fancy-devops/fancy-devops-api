package responses

type BaseRes struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type Result struct {
	BaseRes
	Data string `json:"data"`
}

type IdResult struct {
	BaseRes
	Data int `json:"data"`
}
