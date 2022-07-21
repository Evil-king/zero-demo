package logic

import (
	"context"

	"zero-demo/user-rpc/internal/svc"
	"zero-demo/user-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoLogic) GetUserInfo(in *pb.GetUserInfoReq) (*pb.GetUserInfoResp, error) {
	nickName := "unKnow"
	m := map[int64]string{
		1: "张三 from rpc",
		2: "赵六 from rpc",
	}
	if value, ok := m[in.Id]; ok {
		nickName = value
	}
	userModel := pb.UserModel{
		Id:       in.Id,
		Nickname: nickName,
	}
	return &pb.GetUserInfoResp{
		UserModer: &userModel,
	}, nil
}
