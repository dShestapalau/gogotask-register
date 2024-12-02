package request

type CreateUserRequest struct {
	Login    string `json:"login" binding:"required,notEmpty,min=8,max=15"`
	Password string `json:"password" binding:"required,notEmpty,min=8,max=15"`
}
