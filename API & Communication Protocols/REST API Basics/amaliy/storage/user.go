package storage

import "amaliy/model"



var Users = []model.User{}
var lastID = 0


func CreateUser(user model.User) model.User {
	lastID++
	user.ID = lastID
	Users = append(Users, user)
	return user
}

func GetUser() []model.User{
	return Users
}


func GetUserFiltered(name string, afterID, limit  int) ([]model.User, int) {
	filtered := []model.User{}
	for _, u := range Users{ 
		if  name == "" || u.Name == name {
			if u.ID > afterID {
				filtered = append(filtered, u)
			}
		}
	}
	if len(filtered) > limit {
		filtered = filtered[:limit]
	}
	nextAfter := 0
	if len(filtered) > 0  {
		nextAfter = filtered[len(filtered)-1].ID
	}
	return  filtered, nextAfter
}