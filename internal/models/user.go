package models

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var DemoUser = User{
	Username: "test",
	Password: "test",
}
