package routes

import (
	// "fgd/app/middleware"
	"fgd/controllers/reply"
	"fgd/controllers/thread"
	"fgd/controllers/topic"
	"fgd/controllers/user"
	"fgd/controllers/verify"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

type Controllers struct {
	JWTMiddleware    echoMiddleware.JWTConfig
	ReplyController  reply.ReplyController
	ThreadController thread.ThreadController
	TopicController  topic.TopicController
	UserController   user.UserController
	VerifyController verify.VerifyController
}

func (c *Controllers) Register(e *echo.Echo) {
	e.Use(echoMiddleware.Logger(), echoMiddleware.Recover())

	public := e.Group("public")
	registered := e.Group("registered", echoMiddleware.JWTWithConfig(c.JWTMiddleware))
	// moderator := e.Group("moderator", echoMiddleware.JWTWithConfig(c.JWTMiddleware), middleware.ModeratorValidation)
	// admin := e.Group("admin", echoMiddleware.JWTWithConfig(c.JWTMiddleware), middleware.AdminValidation)

	public.POST("/login", c.UserController.Login)
	public.POST("/register", c.UserController.Register)
	public.POST("/reset", c.VerifyController.RequestForgetPassword)
	public.GET("/user/check", c.UserController.CheckAvailibility)

	registered.GET("/profile/verify", c.VerifyController.RequestEmailVerification)
	registered.POST("/profile/verify", c.VerifyController.SubmitEmailVerification)
	registered.GET("/profile", c.UserController.GetProfile)
	// registered.PUT("/profile", c.UserController.UpdateProfile)
	registered.PUT("/profile/image", c.UserController.UpdateProfileImage)
	registered.GET("/user/:userId", c.UserController.GetPublicProfile)
	registered.GET("/user/:userId/follow", c.UserController.Follow)
	registered.GET("/user/:userId/unfollow", c.UserController.Unfollow)

	public.GET("/topic", c.TopicController.GetTopics)
	registered.POST("/topic", c.TopicController.CreateTopic)
	registered.POST("/topic/check", c.TopicController.CheckAvailibility)
	registered.GET("/topic/:topicId", c.TopicController.GetTopicDetails)
	// registered.GET("/topic/:topicId/moderator", c.TopicController.GetModerators)
	registered.POST("/topic/:topicId/modrequest", c.TopicController.RequestPromotion)
	registered.GET("/topic/:topicId/subscribe", c.TopicController.Subscribe)
	registered.GET("/topic/:topicId/subscribe", c.TopicController.Unsubscribe)
	registered.POST("/topic/:topicId/thread", c.ThreadController.Create)

	// registered.GET("/thread", c.ThreadController.GetThreads)
	registered.PUT("/thread/:threadId", c.ThreadController.Update)
	registered.DELETE("/thread/:threadId", c.ThreadController.Delete)
	registered.POST("/thread/:threadId/like", c.ThreadController.Like)
	registered.DELETE("/thread/:threadId/unlike", c.ThreadController.UndoUnlike)
	registered.POST("/thread/:threadId/unlike", c.ThreadController.Unlike)
	registered.DELETE("/thread/:threadId/like", c.ThreadController.UndoLike)

	registered.POST("/thread/:threadId/reply", c.ReplyController.CreateForThread)
	registered.GET("/reply/:replyId/reply", c.ReplyController.CreateForReply)
	// registered.GET("/reply/:replyId", c.ReplyController.GetChilds)
	// registered.PUT("/reply/:replyId", c.ReplyController.Update)
	registered.POST("/reply/:replyId/like", c.ReplyController.Like)
	registered.DELETE("/reply/:replyId/like", c.ReplyController.UndoLike)
	registered.POST("/reply/:replyId/unlike", c.ReplyController.Unlike)
	registered.DELETE("/reply/:replyId/unlike", c.ReplyController.UndoUnlike)
}
