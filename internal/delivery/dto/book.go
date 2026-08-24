package dto

type BookCreateRequest struct {
	Title  string `form:"title" binding:"required,alphanumspace"`
	Author string `form:"author" binding:"required,alphanumspace"`
}

type BookListRequest struct {
	Limit int `form:"limit,default=10" binding:"required,numeric"`
	Page  int `form:"page,default=1" binding:"required,numeric"`
}

type BookGetRequest struct {
	Id string `uri:"id" binding:"required,uuid"`
}

type BookUpdateRequest struct {
	Id     string `uri:"id" binding:"required,uuid"`
	Title  string `form:"title" binding:"required,alphanumspace"`
	Author string `form:"author" binding:"required,alphanumspace"`
}

type BookDeleteRequest struct {
	Id string `uri:"id" binding:"required,uuid"`
}

type BookResponse struct {
	Id        string `json:"id"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
