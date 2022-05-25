package models

// user struct: this is equivalent to class in other languages
type User struct {
	ID        int
	FirstName string
	LastName  string
}

//declaring variables to be used in other space withing this module
var (
	//users collection container pointers(i.e. '*') to an object of user struct.
	//Pointers ensures that the object of the user itself are not
	// copied around with the app but referenced
	users  []*User
	nextID = 1
)

//return a collection of users pointer
func GetUsers() []*User {
	return users

}

func AddUser(user User) (User, error) {
	user.ID = nextID
	nextID++
	users = append(users, &user)

	return user, nil
}
