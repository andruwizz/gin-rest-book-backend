package request

type BookCreate struct {
	Title  string `form:"title" binding:"required,alphanumspace"`
	Author string `form:"author" binding:"required,alphanumspace"`
}

type BookList struct {
	Limit int `form:"limit,default=10" binding:"required,numeric"`
	Page  int `form:"page,default=1" binding:"required,numeric"`
}

type BookGet struct {
	Id string `uri:"id" binding:"required,uuid"`
}

type BookUpdate struct {
	Id     string `uri:"id" binding:"required,uuid" json:"-"`
	Title  string `form:"title" binding:"required,alphanumspace"`
	Author string `form:"author" binding:"required,alphanumspace"`
}

type BookDelete struct {
	Id string `uri:"id" binding:"required,uuid"`
}
