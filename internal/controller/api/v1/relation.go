package v1

import (
	"douyin_service/global"
	"douyin_service/internal/service"
	"douyin_service/pkg/app"
	"douyin_service/pkg/errcode"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Relation struct { }

func NewRelation() Relation {
	return Relation{}
}

func (r *Relation) Action(c *gin.Context)  {
	param := service.FollowActionRequest{}
	response := app.NewResponse(c)
	var res service.FollowActionResponse
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	valid, tokenErr := app.ValidToken(param.Token, errcode.SkipCheckUserID)
	if !valid {
		global.Logger.Errorf("app.ValidToken errs: %v", tokenErr)
		res.StatusCode = tokenErr.Code()
		res.StatusMsg = tokenErr.Msg()
		response.ToResponse(res)
		return
	}

	// 从token中获取user_id
	claims, err := app.ParseToken(param.Token)
	if err != nil {
		global.Logger.Errorf("app.ParseToken: %v", err)
		response.ToErrorResponse(errcode.ErrorActionFail)
		return
	}
	userId, _ := strconv.Atoi(claims.Audience)
	svc := service.New(c.Request.Context())

	param.UserId = uint(userId)
	if param.ActionType == 1 {
		err = svc.CreateFollow(&param)
		if err != nil {
			global.Logger.Errorf("svc.CreateFollow errs: %v", err)
			response.ToErrorResponse(errcode.ErrorFollowActionFail)
			return
		}
		res.StatusCode = 0
		res.StatusMsg = "关注成功"
		response.ToResponse(res)
		return
	} else if param.ActionType == 2 {
		err = svc.CancelFollow(&param)
		if err != nil {
			global.Logger.Errorf("svc.CancelFollow errs: %v", err)
			response.ToErrorResponse(errcode.ErrorFollowActionFail)
			return
		}

		res.StatusCode = 0
		res.StatusMsg = "取消关注"
		response.ToResponse(res)
		return
	}
}

func (r *Relation) FollowList(c *gin.Context) {
	param := service.FollowListRequest{}
	response := app.NewResponse(c)
	var res service.FollowListResponse
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	valid, tokenErr := app.ValidToken(param.Token, errcode.SkipCheckUserID)
	if !valid {
		global.Logger.Errorf("app.ValidToken errs: %v", tokenErr)
		res.StatusCode = tokenErr.Code()
		res.StatusMsg = tokenErr.Msg()
		response.ToResponse(res)
		return
	}

	svc := service.New(c.Request.Context())
	flwList, err := svc.FollowList(&param)
	resp := &service.FollowListResponse{}
	if err != nil {
		res.StatusCode = errcode.ErrorListFollowFail.Code()
		res.StatusMsg = errcode.ErrorListFollowFail.Msg()
		response.ToResponse(res)
		return
	}

	flwList.StatusCode = 0
	flwList.StatusMsg = "获取关注列表成功"
	resp = &flwList
	response.ToResponse(resp)
}

func (r *Relation) FollowerList(c *gin.Context) {
	param := service.FollowerListRequest{}
	response := app.NewResponse(c)
	var res service.FollowerListResponse
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	valid, tokenErr := app.ValidToken(param.Token, errcode.SkipCheckUserID)
	if !valid {
		global.Logger.Errorf("app.ValidToken errs: %v", tokenErr)
		res.StatusCode = tokenErr.Code()
		res.StatusMsg = tokenErr.Msg()
		response.ToResponse(res)
		return
	}

	svc := service.New(c.Request.Context())
	flwList, err := svc.FollowerList(&param)
	resp := &service.FollowerListResponse{}
	if err != nil {
		res.StatusCode = errcode.ErrorListFollowerFail.Code()
		res.StatusMsg = errcode.ErrorListFollowerFail.Msg()
		response.ToResponse(res)
		return
	}

	flwList.StatusCode = 0
	flwList.StatusMsg = "获取粉丝列表成功"
	resp = &flwList
	response.ToResponse(resp)
}

func (r *Relation) FriendList(c *gin.Context) {
	param := service.FriendListRequest{}
	response := app.NewResponse(c)
	var res service.FollowerListResponse
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	valid, tokenErr := app.ValidToken(param.Token, errcode.SkipCheckUserID)
	if !valid {
		global.Logger.Errorf("app.ValidToken errs: %v", tokenErr)
		res.StatusCode = tokenErr.Code()
		res.StatusMsg = tokenErr.Msg()
		response.ToResponse(res)
		return
	}

	svc := service.New(c.Request.Context())
	frdList, err := svc.FriendList(&param)
	resp := &service.FriendListResponse{}
	if err != nil {
		res.StatusCode = errcode.ErrorListFriendFail.Code()
		res.StatusMsg = errcode.ErrorListFriendFail.Msg()
		response.ToResponse(res)
		return
	}

	frdList.StatusCode = 0
	frdList.StatusMsg = "获取好友列表成功"
	resp = &frdList
	response.ToResponse(resp)
}

