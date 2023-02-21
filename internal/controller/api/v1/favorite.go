package v1

import (
	"douyin_service/global"
	"douyin_service/internal/service"
	"douyin_service/pkg/app"
	"douyin_service/pkg/errcode"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Favorite struct { }

func NewFavorite() Favorite  {
	return  Favorite{}
}

func (f *Favorite) Action(c *gin.Context) {
	param := service.FavoriteActionRequest{}
	response := app.NewResponse(c)
	var res service.FavoriteActionResponse
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
		err = svc.CreateFavorite(&param)
		if err != nil {
			global.Logger.Errorf("svc.CreateFavorite errs: %v", err)
			response.ToErrorResponse(errcode.ErrorActionFavoriteFail)
			return
		}
		res.StatusCode = 0
		res.StatusMsg = "喜欢"
		response.ToResponse(res)
		return
	} else if param.ActionType == 2 {
		err = svc.CancelFavorite(&param)
		if err != nil {
			global.Logger.Errorf("svc.DeleteComment errs: %v", err)
			response.ToErrorResponse(errcode.ErrorActionFavoriteFail)
			return
		}

		res.StatusCode = 0
		res.StatusMsg = "取消点赞"
		response.ToResponse(res)
		return
	}
}

func (f *Favorite) List(c *gin.Context)  {
	param := service.FavoriteListRequest{}
	response := app.NewResponse(c)
	var res service.FavoriteListResponse
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
	fvtList, err := svc.FavoriteList(&param)
	resp := &service.FavoriteListResponse{}
	if err != nil {
		res.StatusCode = errcode.ErrorListCommentFail.Code()
		res.StatusMsg =  errcode.ErrorListCommentFail.Msg()
		response.ToResponse(res)
		return
	}

	fvtList.StatusCode = 0
	fvtList.StatusMsg = "获取喜欢列表成功"
	resp = &fvtList
	response.ToResponse(resp)
}