package user

import (
	"fgd/app/config"
	"fgd/app/middleware"
	"fgd/core/auth"
	"fgd/helper/crypt"
	"fgd/helper/format"
	stringHelper "fgd/helper/string"
	"fmt"
	"strings"
)

type userUsecase struct {
	config         config.Config
	userRepository Repository
	authUsecase    auth.Usecase
	jwtAuth        *middleware.JWTConfig
}

func (uc *userUsecase) GetFollowers(userId int) ([]Domain, error) {
	return uc.userRepository.GetFollowers(userId)
}

func (uc *userUsecase) GetFollowing(userId int) ([]Domain, error) {
	return uc.userRepository.GetFollowing(userId)
}

func (uc *userUsecase) CreateToken(username string, email string, password string) (middleware.CustomToken, error) {
	var userDomain Domain
	var err error

	if username == "" {
		userDomain, err = uc.userRepository.GetUserByEmail(email)
	} else {
		userDomain, err = uc.userRepository.GetUserByUsername(username)
	}

	if err != nil {
		return middleware.CustomToken{}, fmt.Errorf("record not found")
	}

	if password != userDomain.Password {
		if !crypt.ValidateHash(password, userDomain.Password) {
			return middleware.CustomToken{}, fmt.Errorf("password mismatch")
		}
	}

	moderatedTopic, _ := uc.userRepository.GetModeratedTopic(userDomain.ID)

	token, err := uc.jwtAuth.GenerateToken(userDomain.ID, userDomain.Role == "admin", *moderatedTopic.ModeratedTopic)
	if err != nil {
		return middleware.CustomToken{}, fmt.Errorf("error creating token: %v", err)
	}

	err = uc.authUsecase.StoreAuth(userDomain.ID, token)

	return token, err
}

func (uc *userUsecase) CheckUserAvailibility(username string) (bool, error) {
	return uc.userRepository.CheckUserAvailibility(username)
}

func (uc *userUsecase) CreateUser(data *Domain) (Domain, error) {
	var err error

	if data.Username == "" {
		data.Username = stringHelper.GenerateRandomUsername()
	}

	if len(data.Password) < 8 {
		return Domain{}, fmt.Errorf("error: password must be at least 8 character")
	}
	data.Password, err = crypt.CreateHash(data.Password)
	if err != nil {
		return Domain{}, err
	}

	data.Username = strings.ToLower(data.Username)
	data.Email = strings.ToLower(data.Email)

	newUser, err := uc.userRepository.CreateUser(data)
	return newUser, err
}

func (uc *userUsecase) FollowUser(userId, targetId int) error {
	return uc.userRepository.FollowUser(userId, targetId)
}

func (uc *userUsecase) GetPersonalProfile(userId int) (Domain, error) {
	profile, err := uc.userRepository.GetPersonalProfile(userId)
	if err != nil {
		return Domain{}, err
	}

	format.FormatImageLink(uc.config, profile.ProfileImage)

	return profile, nil
}

func (uc *userUsecase) GetProfileByID(userId int) (Domain, error) {
	return uc.userRepository.GetProfileByID(userId)
}

func (uc *userUsecase) GetUsers(limit int, offset int) ([]Domain, error) {
	users, err := uc.userRepository.GetUsers(limit, offset)
	if err != nil {
		return []Domain{}, err
	}

	for _, user := range users {
		format.FormatImageLink(uc.config, user.ProfileImage)
	}

	return users, nil
}

func (uc *userUsecase) UnfollowUser(userId, targetId int) error {
	return uc.userRepository.UnfollowUser(userId, targetId)
}

func (uc *userUsecase) UpdatePassword(newPassword string, userId int) error {
	hashedPass, err := crypt.CreateHash(newPassword)
	if err != nil {
		return err
	}

	err = uc.userRepository.UpdatePassword(hashedPass, userId)

	return err
}

func (uc *userUsecase) UpdatePersonalProfile(data *Domain, userId int) (Domain, error) {
	updatedProfile, err := uc.userRepository.UpdatePersonalProfile(data, userId)
	if err != nil {
		return Domain{}, err
	}

	format.FormatImageLink(uc.config, updatedProfile.ProfileImage)

	return updatedProfile, nil
}

func InitUserUsecase(ac auth.Usecase, r Repository, conf config.Config, jwtConf *middleware.JWTConfig) Usecase {
	return &userUsecase{
		userRepository: r,
		authUsecase:    ac,
		jwtAuth:        jwtConf,
	}
}
