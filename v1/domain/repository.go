package domain

import "fmt"

type usersList map[int64]*User

var users = usersList{
	123: {Id: 123, FirstName: "Hello", LastName: "world", Email: "Hello_world@yahoo.com"},
	124: {Id: 124, FirstName: "Hello", LastName: "world2", Email: "Hello_world2@yahoo.com"},
	125: {Id: 125, FirstName: "Hello", LastName: "world3", Email: "Hello_world3@yahoo.com"},
	126: {Id: 126, FirstName: "Hello", LastName: "world4", Email: "Hello_world4@yahoo.com"},
}

func (li usersList) findById(userId int64) (bool, *User) {
	user, v := li[userId]
	return v, user
}

func GetUserById(userId int64) (*User, error) {
	found, user := users.findById(userId)
	if !found {
		return nil, fmt.Errorf("there is no user with id= %v", userId)
	}
	return user, nil
}
