package model

type Book struct {
	Id        string
	Title     string
	AuthorId  string
	CreatedAt string
	UpdatedAt string
}

type BookRequest struct {
	Title    string `form:"title" binding:"required,alphanumspace"`
	AuthorId string `form:"author_id" binding:"required,alphanumspace"`
}

type BookResponse struct {
	Id        string        `json:"id"`
	Title     string        `json:"title"`
	Author    AuthorProfile `json:"author"`
	CreatedAt string        `json:"created_at"`
	UpdatedAt string        `json:"updated_at"`
}

type AuthorProfile struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
