package repository

import (
	"context"
	"odin/app/common"
	"odin/app/entity"
	"odin/app/repository/memoryRepo"
	"odin/app/repository/pikaRepo"
	"odin/app/repository/redisRepo"

	logger "github.com/tal-tech/loggerX"
)

//本方法为根据回放标识选择repo
//获取repo方法名统一用Get[EntityName]Repo，具体逻辑、入参自行定义
func GetUserRepo(ctx context.Context) (entity.UserRepo, error) {
	daoType := common.MEMORY
	if IS_PLAYBACK := ctx.Value("IS_PLAYBACK"); IS_PLAYBACK != nil {
		if val, ok := IS_PLAYBACK.(string); ok {
			if val == "1" {
				daoType = common.PIKA
			} else {
				daoType = common.REDIS
			}
		}
	}
	switch daoType {
	case common.MEMORY:
		return memoryRepo.NewUserRepo(), nil
	case common.REDIS:
		return redisRepo.NewUserRepo(), nil
	case common.PIKA:
		return pikaRepo.NewUserRepo(), nil
	default:
		return nil, logger.NewError("不存在该存储:" + daoType)
	}
}
