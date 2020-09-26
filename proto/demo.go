package proto

type RedisSetRequest struct {
}
type RedisSetResponse struct {
	Reply []interface{}
}

type RedisGetRequest struct {
}
type RedisGetResponse struct {
	Reply []interface{}
}

type MysqlInsertRequest struct {
}
type MysqlInsertResponse struct {
	Reply []interface{}
}

type MysqlSelectRequest struct {
}
type MysqlSelectResponse struct {
	Reply []interface{}
}

type MysqlDeleteRequest struct {
}
type MysqlDeleteResponse struct {
	Reply []interface{}
}

