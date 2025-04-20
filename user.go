package todo

type User struct {
	Id       int    `json:"-"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Passowrd string `json:"passowrd"`
}