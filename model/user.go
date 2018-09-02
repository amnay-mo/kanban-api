package model

// User is a user
type User struct {
	Email    string `json:"email"     bson:"email"`
	Password string `json:"password"  bson:"password"`
	//	FirstName string `json:"firstname" bson:"firstname"`
	//	LastName  string `json:"lastname"  bson:"lastname"`
}

// Signup sign up a new user
func Signup(user *User) error {
	return taskDatastore.SaveUser(user)
}

// Authenticate authenticates a user
func Authenticate(user *User) bool {
	u, err := taskDatastore.GetUserByEmail(user.Email)
	if err != nil {
		return false
	}
	return u.Password == user.Password
}
