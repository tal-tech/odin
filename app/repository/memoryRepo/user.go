package memoryRepo

import (
	"context"
	"odin/app/common"
	"odin/app/entity"
	"sync"
)

type UserRepo struct {
	store map[int]*entity.User
}

var once sync.Once

var repo *UserRepo

func NewUserRepo() *UserRepo {
	if repo != nil {
		return repo
	}
	once.Do(func() {
		this := new(UserRepo)
		this.store = make(map[int]*entity.User)
		repo = this
	})
	return repo
}

func (this *UserRepo) AddUser(ctx context.Context, user *entity.User) (id int, err error) {
	common.CurrentUserID += 1
	user.Id = common.CurrentUserID
	this.store[user.Id] = user
	id = user.Id
	return
}

func (this *UserRepo) GetUserInfo(ctx context.Context, id int) (user *entity.User, err error) {
	user = this.store[id]
	return
}

func (this *UserRepo) UpdateUser(ctx context.Context, user *entity.User) (err error) {
	this.store[user.Id] = user
	return
}
