package proto

type SayHelloRequest struct {
	Greeting string
}

type SayHelloResponse struct {
	Reply string
}

type MatrixHelloRequest struct {
	Greeting string
}

type MatrixHelloResponse struct {
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

type RedisSetRequest struct {
}

type RedisSetResponse struct {
	Reply []string
}

type RedisGetRequest struct {
}

type RedisGetResponse struct {
	Reply []string
}

type MysqlInsertRequest struct {
}

type MysqlInsertResponse struct {
}

type MysqlSelectRequest struct {
}

type MysqlSelectResponse struct {
	Reply []interface{}
}

type MysqlDeleteRequest struct {
}

type MysqlDeleteResponse struct {
}
