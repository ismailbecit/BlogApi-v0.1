package request

type CategoryInsert struct {
	Name string `validate:"required" json:"name"`
}

type CategoryDel struct {
	ID uint `validate:"required" json:"id"`
}
