package domain

import "fmt"

var users = map[int64]*User{
	123: {Id: 123, FirstName: "Hello", LastName: "world", Email: "Hello_world@yahoo.com"},
	124: {Id: 124, FirstName: "Hello", LastName: "world2", Email: "Hello_world2@yahoo.com"},
	125: {Id: 125, FirstName: "Hello", LastName: "world3", Email: "Hello_world3@yahoo.com"},
	126: {Id: 126, FirstName: "Hello", LastName: "world4", Email: "Hello_world4@yahoo.com"},
}

func GetUserById(userId int64) (*User, error) {
	var user *User
	for key, value := range users {
		if userId == key {
			user = value
		}
	}
	if user == nil {
		return nil, fmt.Errorf("there is no user with id= %v", userId)
	}
	return user, nil
}
