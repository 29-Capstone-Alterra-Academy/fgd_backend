package moderator

import "time"

type Domain struct {
	ID               uint
	UserID           uint
	Username         string
	UserProfileImage *string
	TopicID          uint
	TopicName        string
	CreatedAt        time.Time
}

type Usecase interface {
	ApplyPromotion(userId, topicId uint) (Domain, error)
	ApprovePromotion(promotionId uint) error
	GetPromotionRequest() ([]Domain, error)
	RejectPromotion(promotionId uint) error
	StepDown(userId, topicId uint) error
	RemoveModerator(userId, topicId uint) error
}

type Repository interface {
	ApplyPromotion(userId, topicId uint) (Domain, error)
	ApprovePromotion(promotionId uint) error
	GetPromotionRequest() ([]Domain, error)
	RejectPromotion(promotionId uint) error
	RemoveModerator(userId, topicId uint) error
}
