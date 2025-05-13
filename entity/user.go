package entity

type User struct {
	Id          uint
	Name        string
	PhoneNumber string
	//password always keeps hassed password
	Password string
}
