package proto

type SayHelloRequest struct {
	Greeting string
}

type SayHelloResponse struct {
	Reply string
}

type UserInfoRequest struct {
	Id int
}

type UserInfoResponse struct {
	Name string
	Age  int
	City string
}

type AddUserRequest struct {
	Name string
	Age  int
	City string
}

type AddUserResponse struct {
	Id int
}

type UpdateUserRequest struct {
	Id   int
	Name string
	Age  int
	City string
}

type UpdateUserResponse struct {
}
