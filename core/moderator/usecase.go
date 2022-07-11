package moderator

import (
	"fgd/app/config"
	"fgd/helper/format"
)

type moderatorUsecase struct {
	config              config.Config
	moderatorRepository Repository
}

func (uc *moderatorUsecase) ApplyPromotion(userId uint, topicId uint) (Domain, error) {
	newRequest, err := uc.moderatorRepository.ApplyPromotion(userId, topicId)
	if err != nil {
		return Domain{}, err
	}

	format.FormatImageLink(uc.config, newRequest.UserProfileImage)

	return newRequest, nil
}

func (uc *moderatorUsecase) ApprovePromotion(promotionId uint) error {
	// TODO Invalidate user jwt
	return uc.moderatorRepository.ApprovePromotion(promotionId)
}

func (uc *moderatorUsecase) GetPromotionRequest() ([]Domain, error) {
	reqs, err := uc.moderatorRepository.GetPromotionRequest()
	if err != nil {
		return []Domain{}, err
	}

	for _, req := range reqs {
		format.FormatImageLink(uc.config, req.UserProfileImage)
	}

	return reqs, nil
}

func (uc *moderatorUsecase) RejectPromotion(promotionId uint) error {
	return uc.moderatorRepository.RejectPromotion(promotionId)
}

func (uc *moderatorUsecase) RemoveModerator(userId uint, topicId uint) error {
	// TODO Invalidate user jwt
	return uc.moderatorRepository.RemoveModerator(userId, topicId)
}

func (uc *moderatorUsecase) StepDown(userId uint, topicId uint) error {
	// TODO Invalidate user jwt
	return uc.moderatorRepository.RemoveModerator(userId, topicId)
}

func InitModeratorUsecase(r Repository, conf config.Config) Usecase {
	return &moderatorUsecase{
		config:              conf,
		moderatorRepository: r,
	}
}
