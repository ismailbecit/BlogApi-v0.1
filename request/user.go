package request

type UserInsert struct {
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Age      int    `json:"age"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
