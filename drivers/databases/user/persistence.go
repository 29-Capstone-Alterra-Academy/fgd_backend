package user

import (
	"errors"
	"fgd/core/user"

	"gorm.io/gorm"
)

type persistenceUserRepository struct {
	Conn *gorm.DB
}

func (rp *persistenceUserRepository) GetFollowers(userId int) ([]user.Domain, error) {
	followers := []User{}
	res := rp.Conn.Table("user_follow").Where("following_id", userId).Select("ID", "Username", "ProfileImage").Find(&followers)
	if res.Error != nil {
		return []user.Domain{}, res.Error
	}

	userDomains := []user.Domain{}

	for _, follower := range followers {
		domain := follower.toDomain()
		userDomains = append(userDomains, domain)
	}

	return userDomains, nil
}

func (rp *persistenceUserRepository) GetFollowing(userId int) ([]user.Domain, error) {
	followings := []User{}
	res := rp.Conn.Table("user_follow").Where("user_id", userId).Select("ID", "Username", "ProfileImage").Find(&followings)
	if res.Error != nil {
		return []user.Domain{}, res.Error
	}

	userDomains := []user.Domain{}

	for _, follower := range followings {
		domain := follower.toDomain()
		userDomains = append(userDomains, domain)
	}

	return userDomains, nil
}

func (rp *persistenceUserRepository) GetUserByEmail(email string) (user.Domain, error) {
	user := User{}
	res := rp.Conn.Preload("UserRole").Omit("Following", "Notifications").Where("email = ?", email).Find(&user)
	return user.toDomain(), res.Error
}

func (rp *persistenceUserRepository) GetUserByUsername(username string) (user.Domain, error) {
	user := User{}
	res := rp.Conn.Preload("UserRole").Omit("Following", "Notifications").Where("username = ?", username).Find(&user)
	return user.toDomain(), res.Error
}

func (rp *persistenceUserRepository) CheckIsAdmin(userId int) (bool, error) {
	user := User{}
	res := rp.Conn.Preload("UserRole").Take("Type").Find(&user, userId)

	return user.Role == "admin", res.Error
}

func (rp *persistenceUserRepository) GetModeratedTopic(userId int) (user.Domain, error) {
	moderatedTopic := UserModeratedTopic{}
	res := rp.Conn.Raw("SELECT topic_id FROM topic_moderator WHERE user_id = ?", userId).Scan(&moderatedTopic)

	return moderatedTopic.toDomain(), res.Error
}

func (rp *persistenceUserRepository) CheckUserAvailibility(username string) (bool, error) {
	user := User{}

	err := rp.Conn.Where("username = ?", username).First(&user).Error

	return errors.Is(err, gorm.ErrRecordNotFound), err
}

func (rp *persistenceUserRepository) CreateUser(data *user.Domain) (user.Domain, error) {
	newUser := fromDomain(*data)

	newUser.Role = "user"
	res := rp.Conn.Create(&newUser)

	return newUser.toDomain(), res.Error
}

func (rp *persistenceUserRepository) FollowUser(userId int, targetId int) error {
	err := rp.Conn.
		Model(&User{}).
		Where("id = ?", userId).
		Association("Following").
		Append(&User{
			Model: gorm.Model{ID: uint(targetId)},
		})

	return err
}

func (rp *persistenceUserRepository) GetPersonalProfile(userId int) (user.Domain, error) {
	user := User{}
	res := rp.Conn.Take(&user, userId)

	return user.toDomain(), res.Error
}

func (rp *persistenceUserRepository) GetProfileByID(userId int) (user.Domain, error) {
	user := User{}
	res := rp.Conn.Find(&user, userId)

	return user.toDomain(), res.Error
}

func (rp *persistenceUserRepository) GetUsers(limit int, offset int) ([]user.Domain, error) {
	users := []User{}

	res := rp.Conn.Limit(limit).Offset(offset).Omit("Following", "Notifications").Find(&users)

	if res.Error != nil {
		return []user.Domain{}, res.Error
	}

	userDomains := []user.Domain{}
	for _, user := range users {
		userDomain := user.toDomain()
		var threadCount int64
		var followerCount int64

		rp.Conn.Table("threads").Where("author_id = ?", userDomain.ID).Count(&threadCount)
		userDomain.ThreadCount = int(threadCount)
		rp.Conn.Table("user_follow").Where("following_id = ?", userDomain.ID).Count(&followerCount)
		userDomain.FollowersCount = int(followerCount)

		userDomains = append(userDomains, userDomain)
	}

	return userDomains, nil
}

func (rp *persistenceUserRepository) UnfollowUser(userId int, targetId int) error {
	err := rp.Conn.
		Model(&User{}).
		Where("id = ?", userId).
		Association("Following").
		Delete(&User{
			Model: gorm.Model{ID: uint(targetId)},
		})

	return err
}

func (rp *persistenceUserRepository) UpdatePassword(hashedPassword string, userId int) error {
	res := rp.Conn.Model(&User{}).Where("id = ?", userId).Update("password", hashedPassword)

	return res.Error
}

func (rp *persistenceUserRepository) UpdatePersonalProfile(data *user.Domain, userId int) (user.Domain, error) {
	updatedUser := fromDomain(*data)
	res := rp.Conn.Save(&updatedUser)

	return updatedUser.toDomain(), res.Error
}

func (rp *persistenceUserRepository) UpdateProfileImage(data *user.Domain, userId int) error {
	res := rp.Conn.Model(&User{}).Where("id = ?", userId).Update("profile_image", data.ProfileImage)

	return res.Error
}

func InitPersistenceUserRepository(c *gorm.DB) user.Repository {
	return &persistenceUserRepository{
		Conn: c,
	}
}
