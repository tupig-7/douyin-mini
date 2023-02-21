package v1

import (
	"douyin_service/global"
	"douyin_service/internal/service"
	"douyin_service/pkg/app"
	"douyin_service/pkg/errcode"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Publish struct{}

func NewPublish() Publish {
	return Publish{}
}

// 发布视频
func (p Publish) Action(c *gin.Context) {
	data, _ := c.FormFile("data")
	param := service.PublishActionRequest{
		Data:  data,
		Token: c.PostForm("token"),
		Title: c.PostForm("title"),
	}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	// token不合法
	valid, tokenErr := app.ValidToken(param.Token, errcode.SkipCheckUserID)
	if !valid {
		global.Logger.Errorf("app.ValidToken errs: %v", tokenErr)
		response.ToErrorResponse(tokenErr)
		return
	}

	// 从token中获取user_id
	claims, err := app.ParseToken(param.Token)
	if err != nil {
		global.Logger.Errorf("app.ParseToken: %v", err)
		response.ToErrorResponse(errcode.ErrorActionPublishFail)
		return
	}
	userId, _ := strconv.Atoi(claims.Audience) // 已经是int了还需要再强制转换吗

	// 发布视频
	var resp service.ResponseCommon
	svc := service.New(c.Request.Context())
	err = svc.PublishAction(param.Data, param.Token, param.Title, uint(userId))
	if err != nil {
		global.Logger.Errorf("svc.PublishAction err: %v", err)
		response.ToErrorResponse(errcode.ErrorActionPublishFail)
		return
	}

	resp.StatusCode = 0
	resp.StatusMsg = "视频发布成功"
	response.ToResponse(resp)
}
