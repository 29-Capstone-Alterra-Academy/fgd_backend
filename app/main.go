package main

import (
	"fgd/app/config"
	"fgd/app/middleware"
	"fgd/app/routes"
	"fgd/helper/mail"
	"fgd/helper/storage"
	"fmt"
	"log"
	"time"

	moderatorCtrl "fgd/controllers/moderator"
	replyCtrl "fgd/controllers/reply"
	reportCtrl "fgd/controllers/report"
	searchCtrl "fgd/controllers/search"
	threadCtrl "fgd/controllers/thread"
	topicCtrl "fgd/controllers/topic"
	userCtrl "fgd/controllers/user"
	verifyCtrl "fgd/controllers/verify"

	authCore "fgd/core/auth"
	moderatorCore "fgd/core/moderator"
	replyCore "fgd/core/reply"
	reportCore "fgd/core/report"
	searchCore "fgd/core/search"
	threadCore "fgd/core/thread"
	topicCore "fgd/core/topic"
	userCore "fgd/core/user"
	verifyCore "fgd/core/verify"

	factory "fgd/drivers"
	cache_driver "fgd/drivers/cache"
	_moderatorRepo "fgd/drivers/databases/moderator"
	_replyRepo "fgd/drivers/databases/reply"
	_reportRepo "fgd/drivers/databases/report"
	_searchRepo "fgd/drivers/databases/search"
	_threadRepo "fgd/drivers/databases/thread"
	_topicRepo "fgd/drivers/databases/topic"
	_userRepo "fgd/drivers/databases/user"
	_verifyRepo "fgd/drivers/databases/verify"
	persistence_driver "fgd/drivers/persistence"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func migrate(c *gorm.DB) error {
	c.AutoMigrate(
		&_moderatorRepo.ModeratorRequest{},
		&_reportRepo.ReportReason{},
		&_replyRepo.Reply{},
		&_searchRepo.SearchHistory{},
		&_threadRepo.Thread{},
		&_topicRepo.Topic{},
		&_userRepo.User{},
		&_verifyRepo.Verify{},
		&_reportRepo.UserReport{},
		&_reportRepo.TopicReport{},
		&_reportRepo.ThreadReport{},
		&_reportRepo.ReplyReport{},
	)

	err := c.SetupJoinTable(&_userRepo.User{}, "UserReports", &_reportRepo.UserReport{})
	if err != nil {
		return err
	}
	err = c.SetupJoinTable(&_topicRepo.Topic{}, "TopicReports", &_reportRepo.TopicReport{})
	if err != nil {
		return err
	}
	err = c.SetupJoinTable(&_threadRepo.Thread{}, "ThreadReports", &_reportRepo.ThreadReport{})
	if err != nil {
		return err
	}
	err = c.SetupJoinTable(&_replyRepo.Reply{}, "ReplyReports", &_reportRepo.ReplyReport{})
	if err != nil {
		return err
	}

	return nil
}

func main() {
	conf := config.InitializeConfig()

	cacheConf := cache_driver.CacheConfig{
		Username: conf.CACHE_USERNAME,
		Password: conf.CACHE_PASSWORD,
		Host:     conf.CACHE_HOST,
		Port:     conf.CACHE_PORT,
	}

	dbConf := persistence_driver.PersistenceConfig{
		Username: conf.DB_USERNAME,
		Password: conf.DB_PASSWORD,
		Host:     conf.DB_HOST,
		Port:     conf.DB_PORT,
		Database: conf.DB_NAME,
	}

	storageHelper := storage.NewStorageHelper(conf)

	storageErr := storageHelper.InitializeStaticDirectory()
	if storageErr != nil {
		log.Fatal(storageErr)
	}

	mailHelper, err := mail.NewMailHelper(conf.MAIL_AT, conf.MAIL_RT, conf.MAIL_CLIENT, conf.MAIL_SECRET, conf.MAIL_REDIRECT)
	if err != nil {
		// TODO Handle this better
		log.Fatal(err)
	}

	cacheConn := cacheConf.InitCacheDB()
	dbConn := dbConf.InitPersistenceDB()
	migrateErr := migrate(dbConn)
	if migrateErr != nil {
		log.Fatal(migrateErr)
	}

	authRepo := factory.NewAuthRepository(cacheConn)
	authUsecase := authCore.InitAuthUsecase(authRepo)

	jwtConf := middleware.JWTConfig{
		AuthUsecase:   authUsecase,
		Secret:        conf.JWT_SECRET,
		AccessExpiry:  time.Hour * 8,
		RefreshExpiry: time.Hour * 24 * 7,
	}

	userRepo := factory.NewUserRepository(dbConn)
	topicRepo := factory.NewTopicRepository(dbConn)
	threadRepo := factory.NewThreadRepository(dbConn)
	reportRepo := factory.NewReportRepository(dbConn)
	replyRepo := factory.NewReplyRepository(dbConn)
	moderatorRepo := factory.NewModeratorRepository(dbConn)
	searchRepo := factory.NewSearchRepository(dbConn)
	verifyRepo := factory.NewVerifyRepository(dbConn)

	userUsecase := userCore.InitUserUsecase(authUsecase, userRepo, conf, &jwtConf)
	topicUsecase := topicCore.InitTopicUsecase(topicRepo, userUsecase, conf)
	threadUsecase := threadCore.InitThreadUsecase(threadRepo, topicUsecase, userUsecase, conf)
	reportUsecae := reportCore.InitReportUsecase(reportRepo, conf)
	replyUsecae := replyCore.InitReplyUsecase(replyRepo, userUsecase, conf)
	moderatorUsecase := moderatorCore.InitModeratorUsecase(moderatorRepo, conf)
	searchUsecase := searchCore.InitSearchUsecase(searchRepo)
	verifyUsecase := verifyCore.InitVerifyUsecase(verifyRepo, *mailHelper)

	moderatorController := moderatorCtrl.InitModeratorController(moderatorUsecase)
	reportController := reportCtrl.InitReportController(reportUsecae)
	replyController := replyCtrl.InitReplyController(replyUsecae, storageHelper)
	searchController := searchCtrl.InitSearchController(searchUsecase, threadUsecase, topicUsecase, userUsecase)
	threadController := threadCtrl.InitThreadController(threadUsecase, storageHelper)
	topicController := topicCtrl.InitTopicController(authUsecase, topicUsecase, userUsecase, storageHelper)
	userController := userCtrl.InitUserController(authUsecase, userUsecase, verifyUsecase, conf, jwtConf, storageHelper)
	verifyController := verifyCtrl.InitVerifyController(userUsecase, verifyUsecase)

	e := echo.New()

	routesConf := routes.Controllers{
		JWTMiddleware:       jwtConf.Init(),
		ModeratorController: *moderatorController,
		ReportController:    *reportController,
		ReplyController:     *replyController,
		SearchController:    *searchController,
		ThreadController:    *threadController,
		TopicController:     *topicController,
		UserController:      *userController,
		VerifyController:    *verifyController,
	}

	routesConf.Register(e)

	log.Fatal(e.Start(fmt.Sprintf("%s:%s", conf.HOST, conf.PORT)))
}
