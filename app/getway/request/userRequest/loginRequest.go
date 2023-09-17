package userRequest

// 定义请求
type LoginRequest struct {
	Username string `form:"username"  validate:"required,min=6,max=12" required_msg:"用户账号不能为空" min_msg:"账号长度要在6-12个字符之间" max_msg:"账号长度要在6-12个字符之间"`
	Password string  `form:"password"  validate:"required" required_msg:"密码不能为空"`
}



