package reponse

type ListReposnse struct{
	Data interface{} `json:"data"`
	Total int64 `json:"total"`
}