package entity

type Book struct {
	Id        string `json:"id"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}
