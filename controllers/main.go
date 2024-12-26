package controllers

type PageInfo struct {
	CurrentPage int
	NextPage    int
	PrevPage    int
	TotalPage   int
	TotalData   int
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
