package routes

import (
	"fgd/app/middleware"
	"fgd/controllers/reply"
	"fgd/controllers/report"
	"fgd/controllers/thread"
	"fgd/controllers/topic"
	"fgd/controllers/user"
	"fgd/controllers/verify"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

type Controllers struct {
	JWTMiddleware    echoMiddleware.JWTConfig
	ReportController report.ReportController
	ReplyController  reply.ReplyController
	ThreadController thread.ThreadController
	TopicController  topic.TopicController
	UserController   user.UserController
	VerifyController verify.VerifyController
}

func (c *Controllers) Register(e *echo.Echo) {
	e.Use(echoMiddleware.Logger(), echoMiddleware.Recover(), echoMiddleware.CORS(), echoMiddleware.Static("/"))
	jwtMiddleware := echoMiddleware.JWTWithConfig(c.JWTMiddleware)

	e.Static("/", "files")

	e.POST("/login", c.UserController.Login)
	e.POST("/register", c.UserController.Register)
	e.POST("/reset", c.VerifyController.RequestForgetPassword)
	e.GET("/user/check", c.UserController.CheckAvailibility)
	e.POST("/refresh_token", c.UserController.RefreshToken)

	e.GET("/profile/verify", c.VerifyController.RequestEmailVerification, jwtMiddleware)
	e.POST("/profile/verify", c.VerifyController.SubmitEmailVerification, jwtMiddleware)
	e.GET("/profile", c.UserController.GetProfile, jwtMiddleware)
	e.PUT("/profile", c.UserController.UpdateProfile, jwtMiddleware)
	e.GET("/user/:userId", c.UserController.GetPublicProfile)
	e.POST("/user/:userId/report", c.ReportController.ReportUser, jwtMiddleware)
	e.GET("/user/:userId/follow", c.UserController.Follow, jwtMiddleware)
	e.GET("/user/:userId/unfollow", c.UserController.Unfollow, jwtMiddleware)

	e.GET("/topic", c.TopicController.GetTopics)
	e.POST("/topic", c.TopicController.CreateTopic, jwtMiddleware)
	e.GET("/topic/check", c.TopicController.CheckAvailibility, jwtMiddleware)
	e.GET("/topic/:topicId", c.TopicController.GetTopicDetails, jwtMiddleware)
	e.PUT("/topic/:topicId", c.TopicController.UpdateTopic, jwtMiddleware, middleware.ModeratorValidation)
	// e.GET("/topic/:topicId/moderator", c.TopicController.GetModerators, jwtMiddleware)
	e.POST("/topic/:topicId/report", c.ReportController.ReportTopic, jwtMiddleware)
	e.POST("/topic/:topicId/modrequest", c.TopicController.RequestPromotion, jwtMiddleware)
	e.GET("/topic/:topicId/subscribe", c.TopicController.Subscribe, jwtMiddleware)
	e.GET("/topic/:topicId/subscribe", c.TopicController.Unsubscribe, jwtMiddleware)
	e.POST("/topic/:topicId/thread", c.ThreadController.CreateThread, jwtMiddleware)

	e.GET("/thread", c.ThreadController.GetThreads)
	e.GET("/thread/:threadId", c.ThreadController.GetThread)
	e.PUT("/thread/:threadId", c.ThreadController.UpdateThread, jwtMiddleware)
	e.POST("/thread/:threadId/report", c.ReportController.ReportThread, jwtMiddleware)
	e.DELETE("/thread/:threadId", c.ThreadController.DeleteThread, jwtMiddleware)
	e.POST("/thread/:threadId/like", c.ThreadController.LikeThread, jwtMiddleware)
	e.DELETE("/thread/:threadId/unlike", c.ThreadController.UndoUnlikeThread, jwtMiddleware)
	e.POST("/thread/:threadId/unlike", c.ThreadController.UnlikeThread, jwtMiddleware)
	e.DELETE("/thread/:threadId/like", c.ThreadController.UndoLikeThread, jwtMiddleware)

	e.POST("/thread/:threadId/reply", c.ReplyController.CreateForThread, jwtMiddleware)
	e.GET("/reply/:replyId/reply", c.ReplyController.CreateForReply, jwtMiddleware)
	// e.GET("/reply/:replyId", c.ReplyController.GetChilds, jwtMiddleware)
	e.PUT("/reply/:replyId", c.ReplyController.UpdateReply, jwtMiddleware)
	e.POST("/reply/:replyId/report", c.ReportController.ReportReply, jwtMiddleware)
	e.POST("/reply/:replyId/like", c.ReplyController.LikeReply, jwtMiddleware)
	e.DELETE("/reply/:replyId/like", c.ReplyController.UndoLikeReply, jwtMiddleware)
	e.POST("/reply/:replyId/unlike", c.ReplyController.UnlikeReply, jwtMiddleware)
	e.DELETE("/reply/:replyId/unlike", c.ReplyController.UndoUnlikeReply, jwtMiddleware)

	e.GET("/topic/:topicId/reports", c.ReportController.GetTopicScopeReports, jwtMiddleware, middleware.ModeratorValidation)

	e.GET("/moderation/user/ban_request", c.ReportController.GetUserReports, jwtMiddleware, middleware.AdminValidation)
	e.PUT("/moderation/user/ban_request/:id", c.ReportController.ApproveUserReport, jwtMiddleware, middleware.AdminValidation)
	e.DELETE("/moderation/user/ban_request/:id", c.ReportController.RemoveUserReport, jwtMiddleware, middleware.AdminValidation)

	e.GET("/moderation/topic/ban_request", c.ReportController.GetTopicReports, jwtMiddleware, middleware.AdminValidation)
	e.PUT("/moderation/topic/ban_request/:id", c.ReportController.ApproveTopicReport, jwtMiddleware, middleware.AdminValidation)
	e.DELETE("/moderation/topic/ban_request/:id", c.ReportController.RemoveTopicReport, jwtMiddleware, middleware.AdminValidation)

	e.GET("/moderation/thread/ban_request", c.ReportController.GetThreadReports, jwtMiddleware, middleware.AdminValidation)
	e.PUT("/moderation/thread/ban_request/:id", c.ReportController.ApproveThreadReport, jwtMiddleware, middleware.AdminValidation)
	e.DELETE("/moderation/thread/ban_request/:id", c.ReportController.RemoveThreadReport, jwtMiddleware, middleware.AdminValidation)

	e.GET("/moderation/reply/ban_request", c.ReportController.GetReplyReports, jwtMiddleware, middleware.AdminValidation)
	e.PUT("/moderation/reply/ban_request/:id", c.ReportController.ApproveReplyReport, jwtMiddleware, middleware.AdminValidation)
	e.DELETE("/moderation/reply/ban_request/:id", c.ReportController.RemoveReplyReport, jwtMiddleware, middleware.AdminValidation)

	e.GET("/report/reason", c.ReportController.GetReasons, jwtMiddleware)
	e.POST("/report/reason", c.ReportController.AddReason, jwtMiddleware, middleware.AdminValidation)
	e.DELETE("/report/reason", c.ReportController.DeleteReason, jwtMiddleware, middleware.AdminValidation)
}
