package entity

type User struct {
	name     string
	password string
	email    string
	phone    string
}

func GetName(u User) string {
	return u.name
}
func GetPassword(u User) string {
	return u.password
}
func GetPhone(u User) string {
	return u.phone
}
func GetEmail(u User) string {
	return u.email
}
