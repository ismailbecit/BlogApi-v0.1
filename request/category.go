package request

type CategoryInsert struct {
	Name string `validate:"required" json:"name"`
}

type CategoryDel struct {
	ID int `validate:"required" json:"id"`
}
