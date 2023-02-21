package v1

import (
	"douyin_service/global"
	"douyin_service/internal/service"
	"douyin_service/pkg/app"
	"douyin_service/pkg/errcode"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Comment struct { }

func NewComment() Comment {
	return Comment{}
}

func (cmt *Comment) Action(c *gin.Context)  {
	param := service.CommentActionRequest{}
	response := app.NewResponse(c)
	var res service.CommentActionResponse
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
		res, err = svc.CreateComment(&param)
		if err != nil {
			global.Logger.Errorf("svc.CreateComment errs: %v", err)
			response.ToErrorResponse(errcode.ErrorActionCommentFail)
			return
		}
		res.StatusCode = 0
		res.StatusMsg = "评论成功"
		response.ToResponse(res)
		return
	} else if param.ActionType == 2 {
		err = svc.DeleteComment(&param)
		if err != nil {
			global.Logger.Errorf("svc.DeleteComment errs: %v", err)
			response.ToErrorResponse(errcode.ErrorActionCommentFail)
			return
		}

		res.StatusCode = 0
		res.StatusMsg = "删除成功"
		response.ToResponse(res)
		return
	}
}

func (cmt *Comment) List(c *gin.Context)  {
	param := service.CommentListRequest{}
	response := app.NewResponse(c)
	var res service.CommentListResponse
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
	cmtList, err := svc.CommentList(&param)
	resp := &service.CommentListResponse{}
	if err != nil {
		res.StatusCode = errcode.ErrorListCommentFail.Code()
		res.StatusMsg =  errcode.ErrorListCommentFail.Msg()
		response.ToResponse(res)
		return
	}

	cmtList.StatusCode = 0
	cmtList.StatusMsg = "获取评论列表成功"
	resp = &cmtList
	response.ToResponse(resp)
}