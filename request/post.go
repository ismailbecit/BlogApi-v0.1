package request

type PostInsert struct {
	Title      string `validate:"required" json:"title"`
	Content    string `validate:"required" json:"content"`
	Categoryfk uint   `validate:"required" json:"categoryfk"`
}
type PostDel struct {
	ID int `validate:"required" json:"id"`
}
