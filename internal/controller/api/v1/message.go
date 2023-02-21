package v1

import (
	"douyin_service/global"
	"douyin_service/internal/service"
	"douyin_service/pkg/app"
	"douyin_service/pkg/errcode"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Message struct { }

func NewMsg() Message {
	return Message{}
}

func (m Message) Action(c *gin.Context) {
	param := service.ChatRequest{}
	response := app.NewResponse(c)
	var res service.ChatResponse
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
	param.FromUserId = uint(userId)
	res, err = svc.MessageAction(&param)
	if err != nil {
		global.Logger.Errorf("svc.MessageAction errs: %v", err)
		response.ToErrorResponse(errcode.ErrorActionMessageFail)
		return
	}
	res.StatusCode = 0
	res.StatusMsg = "发送成功"
	response.ToResponse(res)
	return
}

func (m Message) Chat(c *gin.Context) {
	param := service.MessagesRequest{}
	response := app.NewResponse(c)
	var res service.MessagesResponse
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
	param.FromUserId = uint(userId)
	msgList, err := svc.MessageChat(&param)
	resp := &service.MessagesResponse{}
	if err != nil {
		res.StatusCode = errcode.ErrorListMessageFail.Code()
		res.StatusMsg = errcode.ErrorListMessageFail.Msg()
		response.ToResponse(res)
		return
	}
	msgList.StatusCode = 0
	msgList.StatusMsg = "获取消息记录成功"
	resp = &msgList
	response.ToResponse(resp)
}