package controllers

type PageInfo struct {
	CurrentPage int `json:"currentpage" example:"1"`
	NextPage    int `json:"nextpage" example:"102"`
	PrevPage    int `json:"prevpage" example:"0"`
	TotalPage   int `json:"totalpage" example:"1"`
	TotalData   int `json:"totaldata" example:"101"`
}

type TaskResponse struct {
	Success  bool        `json:"success"`
	Message  string      `json:"message"`
	PageInfo PageInfo    `json:"pageInfo,omitempty"`
	Result   interface{} `json:"result,omitempty"`
}

type TaskResponse2 struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Result  interface{} `json:"result,omitempty"`
}

// type UserDB struct {
// 	Id    int    `json:"id"`
// 	Name  string `json:"name"`
// 	Email string `json:"email"`
// }
