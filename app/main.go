package main

import (
	"fgd/app/config"
	"fgd/app/middleware"
	"fgd/app/routes"
	"fgd/helper/mail"
	"fmt"
	"log"
	"os"
	"time"

	replyCtrl "fgd/controllers/reply"
	threadCtrl "fgd/controllers/thread"
	topicCtrl "fgd/controllers/topic"
	userCtrl "fgd/controllers/user"
	verifyCtrl "fgd/controllers/verify"

	authCore "fgd/core/auth"
	replyCore "fgd/core/reply"
	threadCore "fgd/core/thread"
	topicCore "fgd/core/topic"
	userCore "fgd/core/user"
	verifyCore "fgd/core/verify"

	factory "fgd/drivers"
	_authRepo "fgd/drivers/databases/auth"
	_replyRepo "fgd/drivers/databases/reply"
	_threadRepo "fgd/drivers/databases/thread"
	_topicRepo "fgd/drivers/databases/topic"
	_userRepo "fgd/drivers/databases/user"
	_verifyRepo "fgd/drivers/databases/verify"
	persistence_driver "fgd/drivers/persistence"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func migrate(c *gorm.DB) {
	c.AutoMigrate(
		&_authRepo.Auth{},
		&_replyRepo.Reply{},
		&_threadRepo.Thread{},
		&_topicRepo.Topic{},
		&_userRepo.User{},
		&_verifyRepo.Verify{},
	)
}

func main() {
	conf := config.InitializeConfig()
	dbConf := persistence_driver.PersistenceConfig{
		Username: conf.DB_USERNAME,
		Password: conf.DB_PASSWORD,
		Host:     conf.DB_HOST,
		Port:     conf.DB_PORT,
		Database: conf.DB_NAME,
	}

	jwtConf := middleware.JWTConfig{
		Secret:        conf.JWT_SECRET,
		AccessExpiry:  time.Hour * 8,
		RefreshExpiry: time.Hour * 24 * 7,
	}


	mailHelper, err := mail.NewMailHelper(conf.MAIL_AT, conf.MAIL_RT, conf.MAIL_CLIENT, conf.MAIL_SECRET, conf.MAIL_REDIRECT)
	if err != nil {
		// TODO Handle this better
		os.Exit(1)
	}

	dbConn := dbConf.InitPersistenceDB()
	migrate(dbConn)

	authRepo := factory.NewAuthRepository(dbConn)
	userRepo := factory.NewUserRepository(dbConn)
	topicRepo := factory.NewTopicRepository(dbConn)
	threadRepo := factory.NewThreadRepository(dbConn)
	replyRepo := factory.NewReplyRepository(dbConn)
	verifyRepo := factory.NewVerifyRepository(dbConn)

	authUsecase := authCore.InitAuthUsecase(authRepo)
	userUsecase := userCore.InitUserUsecase(authUsecase, userRepo, &jwtConf)
	topicUsecase := topicCore.InitTopicUsecase(topicRepo, userUsecase)
	threadUsecase := threadCore.InitThreadUsecase(threadRepo, topicUsecase, userUsecase)
	replyUsecae := replyCore.InitReplyUsecase(replyRepo, userUsecase)
	verifyUsecase := verifyCore.InitVerifyUsecase(verifyRepo, *mailHelper)

	replyController := replyCtrl.InitReplyController(replyUsecae)
	threadController := threadCtrl.InitThreadController(threadUsecase)
	topicController := topicCtrl.InitTopicController(authUsecase, topicUsecase, userUsecase)
	userController := userCtrl.InitUserController(authUsecase, userUsecase, verifyUsecase, conf)
	verifyController := verifyCtrl.InitVerifyController(userUsecase, verifyUsecase)

	e := echo.New()

	routesConf := routes.Controllers{
		JWTMiddleware:    jwtConf.Init(),
		ReplyController:  *replyController,
		ThreadController: *threadController,
		TopicController:  *topicController,
		UserController:   *userController,
		VerifyController: *verifyController,
	}

	routesConf.Register(e)

	log.Fatal(e.Start(fmt.Sprintf("%s:%s", conf.HOST, conf.PORT)))
}