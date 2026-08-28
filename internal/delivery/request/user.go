package request

type UserCreate struct {
	Name     string `form:"name" binding:"required,alphanumspace"`
	Email    string `form:"email" binding:"required,alphanum"`
	Password string `form:"password" binding:"required,alphanumunicode"`
}

type UserLogin struct {
	Email    string `form:"email" binding:"required,alphanum"`
	Password string `form:"password" binding:"required,alphanumunicode"`
}
