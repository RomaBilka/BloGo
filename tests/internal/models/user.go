package models

type UserId int

type User struct {
	Id    UserId
	Name  string
	Email string
	Phone string
}
