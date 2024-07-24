package validations

type RegistrationData struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6,max=64"`
	Username string `json:"username" binding:"required"`
}

type LoginData struct {
	Identifier string `json:"identifier" binding:"required"`
	Password  string `json:"password" binding:"required"`
}