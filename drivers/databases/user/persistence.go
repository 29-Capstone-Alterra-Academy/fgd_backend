package user

import (
	"errors"
	"fgd/core/user"
	"fmt"
	"strings"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type persistenceUserRepository struct {
	Conn *gorm.DB
}

func (rp *persistenceUserRepository) GetModerators(topicId int) ([]user.Domain, error) {
	moderators := []User{}
	err := rp.Conn.Model(&User{}).Select("users.id", "users.username", "users.profile_image").Joins("left join topic_moderator on topic_moderator.user_id = users.id").Where("topic_moderator.topic_id = ?", topicId).Find(&moderators).Error
	if err != nil {
		return []user.Domain{}, err
	}

	userDomains := []user.Domain{}

	for _, follower := range moderators {
		userDomains = append(userDomains, follower.toDomain())
	}

	return userDomains, nil
}

func (rp *persistenceUserRepository) GetFollowers(userId int) ([]user.Domain, error) {
	followers := []User{}
	err := rp.Conn.Table("user_follow").Where("following_id = ?", userId).Select("ID", "Username", "ProfileImage").Find(&followers).Error
	if err != nil {
		return []user.Domain{}, err
	}

	userDomains := []user.Domain{}

	for _, follower := range followers {
		userDomains = append(userDomains, follower.toDomain())
	}

	return userDomains, nil
}

func (rp *persistenceUserRepository) GetFollowing(userId int) ([]user.Domain, error) {
	followings := []User{}
	err := rp.Conn.Table("user_follow").Where("user_id", userId).Select("ID", "Username", "ProfileImage").Find(&followings).Error
	if err != nil {
		return []user.Domain{}, err
	}

	userDomains := []user.Domain{}

	for _, follower := range followings {
		userDomains = append(userDomains, follower.toDomain())
	}

	return userDomains, nil
}

func (rp *persistenceUserRepository) GetUserByEmail(email string) (user.Domain, error) {
	user := User{}
	err := rp.Conn.Omit("Following", "Notifications").Where("email = ?", email).Find(&user).Error
	return user.toDomain(), err
}

func (rp *persistenceUserRepository) GetUserByUsername(username string) (user.Domain, error) {
	user := User{}
	err := rp.Conn.Omit("Following", "Notifications").Where("username = ?", username).Find(&user).Error
	return user.toDomain(), err
}

func (rp *persistenceUserRepository) CheckIsAdmin(userId int) (bool, error) {
	user := User{}
	err := rp.Conn.Take("Role").Find(&user, userId).Error

	return user.Role == "admin", err
}

func (rp *persistenceUserRepository) GetModeratedTopic(userId int) (user.Domain, error) {
	moderatedTopic := UserModeratedTopic{}
	topics := []struct {
		TopicID uint
	}{}
	res := rp.Conn.Table("topic_moderator").Where("user_id = ?", userId).Raw("SELECT topic_id FROM topic_moderator WHERE user_id = ?", userId).Scan(&topics)

	for _, topic := range topics {
		moderatedTopic.TopicID = append(moderatedTopic.TopicID, int(topic.TopicID))
	}

	return moderatedTopic.toDomain(), res.Error
}

func (rp *persistenceUserRepository) CheckUserAvailibility(username string) bool {
	user := User{}
	err := rp.Conn.Where("username = ?", username).Take(&user).Error

	return !errors.Is(err, gorm.ErrRecordNotFound)
}

func (rp *persistenceUserRepository) CheckEmailAvailibility(email string) bool {
	user := User{}
	err := rp.Conn.Where("email = ?", email).Take(&user).Error

	return !errors.Is(err, gorm.ErrRecordNotFound)
}

func (rp *persistenceUserRepository) CreateUser(data *user.Domain) (user.Domain, error) {
	tx := rp.Conn.Begin()

	newUser := fromDomain(*data)
	checkUsername := User{}
	checkEmail := User{}
	fetchUsernameErr := tx.Select("username").Where("UPPER(username) = UPPER(?)", newUser.Username).Take(&checkUsername).Error
	if fetchUsernameErr != nil && !errors.Is(fetchUsernameErr, gorm.ErrRecordNotFound) {
		return user.Domain{}, fmt.Errorf("error: username already in use")
	}
	fetchEmailErr := tx.Select("email").Where("email = ?", newUser.Email).Take(&checkEmail).Error
	if fetchEmailErr != nil && !errors.Is(fetchEmailErr, gorm.ErrRecordNotFound) {
		return user.Domain{}, fmt.Errorf("error: email already in use")
	}
	if newUser.Username == checkUsername.Username || newUser.Email == checkEmail.Email {
		return user.Domain{}, fmt.Errorf("error: username/email already in use")
	}

	err := tx.Omit(clause.Associations).Create(&newUser).Error
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if !strings.EqualFold(pgErr.Code, pgerrcode.UniqueViolation) {
				tx.Rollback()
				return user.Domain{}, err
			}
		}
	}

	err = tx.Commit().Error
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if !strings.EqualFold(pgErr.Code, pgerrcode.UniqueViolation) {
				tx.Rollback()
				return user.Domain{}, err
			}
		}
	}

	return newUser.toDomain(), nil
}

func (rp *persistenceUserRepository) FollowUser(userId int, targetId int) error {
	user := User{Model: gorm.Model{ID: uint(userId)}}
	targetUser := User{Model: gorm.Model{ID: uint(targetId)}}

	tx := rp.Conn.Begin()

	err := tx.
		Model(&user).
		Association("Following").
		Append(&targetUser)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit().Error
	if err != nil {
		return err
	}

	return nil
}

func (rp *persistenceUserRepository) GetPersonalProfile(userId int) (user.Domain, error) {
	user := User{}
	res := rp.Conn.Take(&user, userId)

	return user.toDomain(), res.Error
}

func (rp *persistenceUserRepository) GetProfileByID(userId int) (user.Domain, error) {
	tx := rp.Conn.Session(&gorm.Session{SkipDefaultTransaction: true})

	user := User{}
	err := tx.Preload(clause.Associations).Take(&user, userId).Error
	if err != nil {
		return user.toDomain(), err
	}

	userDomain := user.toDomain()

	var threadCount int64
	var followerCount int64
	var followingCount int64

	tx.Table("threads").Where("author_id = ?", userDomain.ID).Count(&threadCount)
	userDomain.ThreadCount = int(threadCount)
	tx.Table("user_follow").Where("following_id = ?", userDomain.ID).Count(&followerCount)
	userDomain.FollowersCount = int(followerCount)
	tx.Table("user_follow").Where("user_id = ?", userDomain.ID).Count(&followingCount)
	userDomain.FollowingCount = int(followingCount)

	return userDomain, nil
}

func (rp *persistenceUserRepository) GetUsers(limit int, offset int) ([]user.Domain, error) {
	tx := rp.Conn.Session(&gorm.Session{SkipDefaultTransaction: true})

	users := []User{}

	err := tx.Limit(limit).Offset(offset).Omit("Following", "Notifications").Find(&users).Error
	if err != nil {
		return []user.Domain{}, err
	}

	userDomains := []user.Domain{}
	for _, user := range users {
		userDomain := user.toDomain()
		var threadCount int64
		var followerCount int64

		tx.Table("threads").Where("author_id = ?", userDomain.ID).Count(&threadCount)
		userDomain.ThreadCount = int(threadCount)
		tx.Table("user_follow").Where("following_id = ?", userDomain.ID).Count(&followerCount)
		userDomain.FollowersCount = int(followerCount)

		userDomains = append(userDomains, userDomain)
	}

	return userDomains, nil
}

func (rp *persistenceUserRepository) GetUsersByKeyword(keyword string, limit int, offset int) ([]user.Domain, error) {
	tx := rp.Conn.Session(&gorm.Session{SkipDefaultTransaction: true})

	users := []User{}

	err := tx.Limit(limit).Offset(offset).Omit("Following", "Notifications").Where("UPPER(username) LIKE UPPER(?)", "%"+keyword+"%").Find(&users).Error
	if err != nil {
		return []user.Domain{}, err
	}

	userDomains := []user.Domain{}
	for _, user := range users {
		userDomain := user.toDomain()
		var followerCount int64

		tx.Table("user_follow").Where("following_id = ?", userDomain.ID).Count(&followerCount)
		userDomain.FollowersCount = int(followerCount)

		userDomains = append(userDomains, userDomain)
	}

	return userDomains, nil
}

func (rp *persistenceUserRepository) UnfollowUser(userId int, targetId int) error {
	user := User{Model: gorm.Model{ID: uint(userId)}}
	targetUser := User{Model: gorm.Model{ID: uint(targetId)}}

	tx := rp.Conn.Begin()

	err := tx.
		Model(&user).
		Association("Following").
		Delete(&targetUser)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit().Error
	if err != nil {
		return err
	}

	return nil
}

func (rp *persistenceUserRepository) UpdatePassword(hashedPassword string, userId int) error {
	tx := rp.Conn.Begin()
	err := tx.Model(&User{}).Where("id = ?", userId).Update("password", hashedPassword).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (rp *persistenceUserRepository) UpdatePersonalProfile(data *user.Domain, userId int) (user.Domain, error) {
	existingUser := User{}
	fetchResult := rp.Conn.Take(&existingUser, userId)
	if fetchResult.Error != nil {
		return user.Domain{}, fetchResult.Error
	}

	updatedUser := fromDomain(*data)

	if updatedUser.Username != "" {
		existingUser.Username = updatedUser.Username
	}
	if updatedUser.Bio != nil && *updatedUser.Bio != "" {
		existingUser.Bio = updatedUser.Bio
	}
	if updatedUser.BirthDate != nil {
		existingUser.BirthDate = updatedUser.BirthDate
	}
	if updatedUser.Gender != nil && *updatedUser.Gender != "" {
		existingUser.Gender = updatedUser.Gender
	}
	if updatedUser.ProfileImage != nil && *updatedUser.ProfileImage != "" {
		existingUser.ProfileImage = updatedUser.ProfileImage
	}

	tx := rp.Conn.Begin()

	err := tx.Save(&existingUser).Error
	if err != nil {
		tx.Rollback()
		return user.Domain{}, err
	}

	err = tx.Commit().Error
	if err != nil {
		return user.Domain{}, err
	}

	return existingUser.toDomain(), nil
}

func InitPersistenceUserRepository(c *gorm.DB) user.Repository {
	return &persistenceUserRepository{
		Conn: c,
	}
}
