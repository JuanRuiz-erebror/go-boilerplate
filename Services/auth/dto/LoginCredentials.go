package dto

type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	AccessUuid   string
	RefreshUuid  string
	AtExpires    int64
	RtExpires    int64
}

//Login credential
type LoginCredentials struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}

type User struct {
	ID uint64 `json:"id"`
	//	ID       int    `json:"id" example:"1"`
	Name     string `json:"name" example:"bottle_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}

var UserTest = User{
	ID:       1,
	Name:     "username",
	Email:    "mmmmm",
	Password: "password",
	Phone:    "49123454322", //this is a random number
}

var users = []User{
	{ID: 1, Email: "juan.ruiz@iwo-os.com", Name: "Juan"},
}
