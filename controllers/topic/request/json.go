package request

import "fgd/core/topic"

type NewTopic struct {
	Name         string `json:"name"`
	ProfileImage string `json:"profile_image"`
	Description  string `json:"description"`
	Rules        string `json:"rules"`
}

func (r *NewTopic) ToDomain() *topic.Domain {
	return &topic.Domain{
		Name:         r.Name,
		ProfileImage: &r.ProfileImage,
		Description:  r.Description,
		Rules:        &r.Rules,
	}
}
