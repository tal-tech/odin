package serviceImpl

import (
	"context"
	"math/rand"
	"time"

	"odin/app/repository/mysqlRepo"

	"odin/proto"

	"github.com/spf13/cast"
	logger "github.com/tal-tech/loggerX"
	dbdao "github.com/tal-tech/torm"
	redisdao "github.com/tal-tech/xredis"
)

type DemoService struct {
}

func NewDemoService() *DemoService {
	return new(DemoService)
}

const REDIS_KEY_PREFIX = "odin_demo_%v"

func (this *DemoService) RedisSet(ctx context.Context, req *proto.RedisSetRequest, resp *proto.RedisSetResponse) error {
	instance := redisdao.NewSimpleXesRedis(ctx, "redis")
	for i := 0; i < 100; i++ {
		value := rand.Intn(10000)
		_, err := instance.Set(REDIS_KEY_PREFIX, []interface{}{i}, value, 0)
		if err != nil {
			logger.Ex(ctx, "RedisSet", "error:%v")
			return err
		}
		resp.Reply = append(resp.Reply, cast.ToString(value))
	}
	return nil
}

func (this *DemoService) RedisGet(ctx context.Context, req *proto.RedisGetRequest, resp *proto.RedisGetResponse) error {
	instance := redisdao.NewSimpleXesRedis(ctx, "redis")
	for i := 0; i < 100; i++ {
		value, err := instance.Get(REDIS_KEY_PREFIX, []interface{}{i})
		if err != nil {
			logger.Ex(ctx, "RedisGet", "error:%v")
			return err
		}
		resp.Reply = append(resp.Reply, cast.ToString(value))
	}
	return nil
}

func (this *DemoService) MysqlInsert(ctx context.Context, req *proto.MysqlInsertRequest, resp *proto.MysqlInsertResponse) error {
	dao := mysqlRepo.NewStudentDao()
	s := mysqlRepo.Student{
		Sex:        1,
		Name:       "student-" + cast.ToString(rand.Intn(100000)),
		City:       "beijing",
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	_, err := dao.Create(&s)
	if err != nil {
		logger.Ex(ctx, "MysqlInsert", "error:%v", err)
		return err
	}
	return nil
}

func (this *DemoService) MysqlSelect(ctx context.Context, req *proto.MysqlSelectRequest, resp *proto.MysqlSelectResponse) error {
	ins := dbdao.GetDbInstance("godemo", "reader")
	dao := mysqlRepo.NewStudentDao(ins.Engine)
	var param dbdao.ParamDesc
	param = true
	stulist, err := dao.Get(param)
	if err != nil {
		logger.Ex(ctx, "MysqlSelect", "error:%v", err)
		return err
	}
	for _, stu := range stulist {
		resp.Reply = append(resp.Reply, stu)
	}
	return nil
}

func (this *DemoService) MysqlDelete(ctx context.Context, req *proto.MysqlDeleteRequest, resp *proto.MysqlDeleteResponse) error {
	session := dbdao.GetDbInstance("godemo", "writer").GetSession()
	defer session.Close()
	dao := mysqlRepo.NewStudentDao(session)
	//全表扫，注意不要这样使用，传入id或范围
	stulist, err := dao.Get(nil)
	if err != nil {
		logger.Ex(ctx, "MysqlDelete", "select error:%v", err)
		return err
	}
	for _, stu := range stulist {
		_, err := dao.Delete(&stu)
		if err != nil {
			logger.Ex(ctx, "MysqlDelete", "delete error:%v", err)
			return err
		}
	}
	e := session.Commit()
	if e != nil {
		logger.Ex(ctx, "MysqlDelete", "commit error:%v", err)
		session.Rollback()
		return e
	}

	return nil
}
