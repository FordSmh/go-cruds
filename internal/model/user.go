package model

type User struct {
	Username string
	Password string
	Role     string
}

var Users = []User{
	{Username: "admin", Password: "admin123", Role: "admin"},
	{Username: "editor", Password: "editor123", Role: "editor"},
}

func ValidateUser(username string, password string) *User {
	for _, u := range Users {
		if u.Username == username && u.Password == password {
			return &u
		}
	}
	return nil
}

func GetUserByUsername(email string) *User {
	for _, u := range Users {
		if u.Username == email {
			return &u
		}
	}
	return nil
}
