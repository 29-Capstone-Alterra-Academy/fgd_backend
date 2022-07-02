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
	e.Use(echoMiddleware.Logger(), echoMiddleware.Recover(), echoMiddleware.CORS())
	jwtMiddleware := echoMiddleware.JWTWithConfig(c.JWTMiddleware)

	e.POST("/login", c.UserController.Login)
	e.POST("/register", c.UserController.Register)
	e.POST("/reset", c.VerifyController.RequestForgetPassword)
	e.GET("/user/check", c.UserController.CheckAvailibility)

	e.GET("/profile/verify", c.VerifyController.RequestEmailVerification, jwtMiddleware)
	e.POST("/profile/verify", c.VerifyController.SubmitEmailVerification, jwtMiddleware)
	e.GET("/profile", c.UserController.GetProfile, jwtMiddleware)
	e.PUT("/profile", c.UserController.UpdateProfile, jwtMiddleware)
	e.GET("/user/:userId", c.UserController.GetPublicProfile)
	e.GET("/user/:userId/follow", c.UserController.Follow, jwtMiddleware)
	e.GET("/user/:userId/unfollow", c.UserController.Unfollow, jwtMiddleware)

	e.GET("/topic", c.TopicController.GetTopics)
	e.POST("/topic", c.TopicController.CreateTopic, jwtMiddleware)
	e.POST("/topic/check", c.TopicController.CheckAvailibility, jwtMiddleware)
	e.GET("/topic/:topicId", c.TopicController.GetTopicDetails, jwtMiddleware)
	// e.GET("/topic/:topicId/moderator", c.TopicController.GetModerators, jwtMiddleware)
	e.POST("/topic/:topicId/modrequest", c.TopicController.RequestPromotion, jwtMiddleware)
	e.GET("/topic/:topicId/subscribe", c.TopicController.Subscribe, jwtMiddleware)
	e.GET("/topic/:topicId/subscribe", c.TopicController.Unsubscribe, jwtMiddleware)
	e.POST("/topic/:topicId/thread", c.ThreadController.CreateThread, jwtMiddleware)

	e.GET("/thread", c.ThreadController.GetThreads)
	e.GET("/thread/:threadId", c.ThreadController.GetThread)
	e.PUT("/thread/:threadId", c.ThreadController.UpdateThread, jwtMiddleware)
	e.DELETE("/thread/:threadId", c.ThreadController.DeleteThread, jwtMiddleware)
	e.POST("/thread/:threadId/like", c.ThreadController.LikeThread, jwtMiddleware)
	e.DELETE("/thread/:threadId/unlike", c.ThreadController.UndoUnlikeThread, jwtMiddleware)
	e.POST("/thread/:threadId/unlike", c.ThreadController.UnlikeThread, jwtMiddleware)
	e.DELETE("/thread/:threadId/like", c.ThreadController.UndoLikeThread, jwtMiddleware)

	e.POST("/thread/:threadId/reply", c.ReplyController.CreateForThread, jwtMiddleware)
	e.GET("/reply/:replyId/reply", c.ReplyController.CreateForReply, jwtMiddleware)
	// e.GET("/reply/:replyId", c.ReplyController.GetChilds, jwtMiddleware)
	e.PUT("/reply/:replyId", c.ReplyController.UpdateReply, jwtMiddleware)
	e.POST("/reply/:replyId/like", c.ReplyController.LikeReply, jwtMiddleware)
	e.DELETE("/reply/:replyId/like", c.ReplyController.UndoLikeReply, jwtMiddleware)
	e.POST("/reply/:replyId/unlike", c.ReplyController.UnlikeReply, jwtMiddleware)
	e.DELETE("/reply/:replyId/unlike", c.ReplyController.UndoUnlikeReply, jwtMiddleware)
}
