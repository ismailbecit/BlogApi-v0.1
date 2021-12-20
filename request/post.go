package request

type PostInsert struct {
	Title      string `json:"title"`
	Content    string `json:"content"`
	Userfk     uint   `json:"userfk"`
	Categoryfk uint   `json:"categoryfk"`
}
type PostDel struct {
	ID int `json:"id"`
}
