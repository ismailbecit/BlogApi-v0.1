package request

type UserInsert struct {
	Name     string `validate:"required" json:"name"`
	Surname  string `validate:"required" json:"surname"`
	Email    string `validate:"required" json:"email"`
	Password string `validate:"required" json:"password"`
	Age      int    `validate:"required" json:"age"`
}

type UserLogin struct {
	Email    string `validate:"required" json:"email"`
	Password string `validate:"required" json:"password"`
}

type UserDel struct {
	Id uint `validate:"required" json:"id"`
}
