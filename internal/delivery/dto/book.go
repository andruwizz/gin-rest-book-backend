package dto

type BookRequest struct {
	Title  string `form:"title" binding:"required,alphanumspace"`
	Author string `form:"author" binding:"required,alphanumspace"`
}

type BookResponse struct {
	Id        string `json:"id"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
