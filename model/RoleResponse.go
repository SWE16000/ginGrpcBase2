package model

type RoleResponse struct {
	Id   int64  `form:"id"`
	Name string  `form:"name"`
	Created_at string `form:"created_at"`
	Updated_at string `form:"updated_at"`
}