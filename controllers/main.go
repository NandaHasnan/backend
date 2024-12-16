package controllers

type TaskResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}

type User struct {
	Id       int    `json:"id"`
	Fullname string `json:"fullname" form:"fullname"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

var users []User = []User{
	{
		Id:       1,
		Fullname: "fazz",
		Email:    "fazz@mail.com",
		Password: "$argon2id$v=19$m=65536,t=1,p=2$Zm9vYmFyYmF6$FOMYAEWsUSM4Cmj2YNQ8Xyoa6fw11J/EysRQp2/RIMw", //  "1234"
	},
	{
		Id:       2,
		Fullname: "track",
		Email:    "track@mail.com",
		Password: "$argon2id$v=19$m=65536,t=1,p=2$Zm9vYmFyYmF6$RPeB4d/todvjVR4QFvP3qOK2cAHpAjDZspXeEKMaJvU", //  "5678"
	},
	{
		Id:       3,
		Fullname: "endra",
		Email:    "endra@mail.com",
		Password: "$argon2id$v=19$m=65536,t=1,p=2$Zm9vYmFyYmF6$FOMYAEWsUSM4Cmj2YNQ8Xyoa6fw11J/EysRQp2/RIMw", //  "1234"
	},
	{
		Id:       4,
		Fullname: "adiv",
		Email:    "adiv@mail.com",
		Password: "$argon2id$v=19$m=65536,t=1,p=2$Zm9vYmFyYmF6$RPeB4d/todvjVR4QFvP3qOK2cAHpAjDZspXeEKMaJvU", //  "5678"
	},
	{
		Id:       5,
		Fullname: "rinaldi",
		Email:    "rinaldi@mail.com",
		Password: "$argon2id$v=19$m=65536,t=1,p=2$Zm9vYmFyYmF6$FOMYAEWsUSM4Cmj2YNQ8Xyoa6fw11J/EysRQp2/RIMw", //  "1234"
	},
	{
		Id:       6,
		Fullname: "rama",
		Email:    "rama@mail.com",
		Password: "$argon2id$v=19$m=65536,t=1,p=2$Zm9vYmFyYmF6$RPeB4d/todvjVR4QFvP3qOK2cAHpAjDZspXeEKMaJvU", //  "5678"
	},
	{
		Id:       7,
		Fullname: "alwi",
		Email:    "alwi@mail.com",
		Password: "$argon2id$v=19$m=65536,t=1,p=2$Zm9vYmFyYmF6$FOMYAEWsUSM4Cmj2YNQ8Xyoa6fw11J/EysRQp2/RIMw", //  "1234"
	},
	{
		Id:       8,
		Fullname: "joko",
		Email:    "joko@mail.com",
		Password: "$argon2id$v=19$m=65536,t=1,p=2$Zm9vYmFyYmF6$RPeB4d/todvjVR4QFvP3qOK2cAHpAjDZspXeEKMaJvU", //  "5678"
	},
}
