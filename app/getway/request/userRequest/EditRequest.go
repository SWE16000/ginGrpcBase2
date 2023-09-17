package userRequest

type EditRequest struct {
	Name string  `form:"name" validate:"required,min=6,max=12" required_msg:"用户名称不能为空" min_msg:"用户名称长度要在6-12个字符之间" max_msg:"用户名称要在6-12个字符之间"`
	Username string `form:"username" validate:"required,min=6,max=12" required_msg:"用户账号不能为空" min_msg:"账号长度要在6-12个字符之间" max_msg:"账号长度要在6-12个字符之间"`
	Password string  `form:"password" validate:"required,min=6,max=18" required_msg:"用户密码不能为空" min_msg:"用户密码长度要在6-18个字符之间" max_msg:"用户密码长度要在6-18个字符之间"`
	Phone string     `form:"phone" validate:"required" required_msg:"用户手机号不能为空"`
	RoleId []int64   `form:"roleId" validate:"required" required_msg:"用户角色不能为空"`
}

