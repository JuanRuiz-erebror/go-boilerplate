package dto

//Login credential
type LoginCredentials struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}

type User struct {
	ID    int    `json:"id" example:"1"`
	Name  string `json:"name" example:"bottle_name"`
	Email string `json:"email"`
}

var users = []User{
	{ID: 1, Email: "juan.ruiz@iwo-os.com", Name: "Juan"},
}
