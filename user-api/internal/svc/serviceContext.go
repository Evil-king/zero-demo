package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"zero-demo/user-api/internal/config"
	model "zero-demo/user-api/model"
	"zero-demo/user-rpc/usercenter"
)

type ServiceContext struct {
	Config     config.Config
	UserModel  model.UserModel
	UserClient usercenter.UserCenter
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		UserModel:  model.NewUserModel(sqlx.NewMysql(c.DB.DataSource)),
		UserClient: usercenter.NewUserCenter(zrpc.MustNewClient(c.UserRpcConf)),
	}
}
