package logic

import (
	"context"
	"zero-demo/user-api/internal/svc"
	"zero-demo/user-api/internal/types"
	"zero-demo/user-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(req *types.UserInfoReq) (resp *types.UserInfoResp, err error) {
	// todo: add your logic here and delete this line
	logx.Infof("入参打印，req：%#v\n", req)
	//user,err:=l.svcCtx.UserModel.FindOne(l.ctx,req.UserId)
	//if err != nil && err != sqlc.ErrNotFound {
	//	return nil,errors.New("查询数据失败")
	//}
	//if user == nil {
	//	return nil,errors.New("用户名不存在")
	//}
	userInfoResp, err := l.svcCtx.UserClient.GetUserInfo(l.ctx, &pb.GetUserInfoReq{
		Id: req.UserId,
	})
	if err != nil {
		return nil, err
	}
	return &types.UserInfoResp{
		UserId:   userInfoResp.UserModer.Id,
		NickName: userInfoResp.UserModer.Nickname,
	}, nil
}
