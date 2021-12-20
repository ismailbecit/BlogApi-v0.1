package request

type CategoryInsert struct {
	Name string `json:"name"`
}

type CategoryDel struct {
	ID int `json:"id"`
}
