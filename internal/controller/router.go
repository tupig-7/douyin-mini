package controller

import (
	"douyin_service/global"
	v1 "douyin_service/internal/controller/api/v1"
	"douyin_service/internal/middleware"
	"douyin_service/pkg/limiter"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var methodLimiters = limiter.NewMethodLimiter().AddBuckets(limiter.LimiterBucketRule{
	Key:          "/auth",
	FillInterval: time.Second * 10,
	Capacity:     20,
	Quantum:      20,
})

func NewRouter() *gin.Engine {
	r := gin.New()
	if global.ServerSetting.RunMode == "debug" {
		r.Use(gin.Logger())
		r.Use(gin.Recovery())
	} else {
		r.Use(middleware.AccessLog())
		r.Use(middleware.Recovery())
	}
	r.Use(middleware.Cors())
	r.Use(middleware.RateLimiter(methodLimiters))
	r.Use(middleware.ContextTimeout(60 * time.Second))
	r.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))

	user := v1.NewUser()
	publish := v1.NewPublish()
	// msg := v1.NewMsg()
	msg := v1.NewMsg()
	cmt := v1.NewComment()
	apiv1 := r.Group("/douyin/")
	apiv1.Use()
	{
		// user
		apiv1.POST("/user/login/", user.Login)
		apiv1.POST("/user/register/", user.Register)
		apiv1.GET("/user/", user.Get)

		// message
		apiv1.GET("/message/chat", msg.Chat)
		apiv1.POST("/message/action", msg.Action)

		// comment
		apiv1.GET("/comment/list/", cmt.List)
		// publish
		apiv1.POST("/publish/action/", publish.Action)
		apiv1.POST("/publish/list/", publish.List)

	}

	return r
}
