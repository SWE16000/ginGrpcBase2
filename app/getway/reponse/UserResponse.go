package reponse

type UserResponse struct {
	User interface{} `json:"user"`
	Token string `json:"token"`
}
