package entity

import "context"

type User struct {
	Id   int
	Name string
	Age  int
	City string
}

type UserRepo interface {
	AddUser(context.Context, *User) (int, error)
	GetUserInfo(ctx context.Context, id int) (*User, error)
	UpdateUser(context.Context, *User) error
}
