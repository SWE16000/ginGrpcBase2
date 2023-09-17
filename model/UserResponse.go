package model

type UserResponse struct {
	Id   int64  `form:"id"`
	Name string  `form:"name"`
	Username string `form:"username"`
	Phone string     `form:"phone"`
	Created_at string `form:"created_at"`
	Updated_at string `form:"updated_at"`
	Roles [] Role
}
