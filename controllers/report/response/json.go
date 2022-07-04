package response

import (
	"fgd/core/report"
	"time"
)

type User struct {
	ID           uint    `json:"id"`
	Username     string  `json:"username"`
	ProfileImage *string `json:"profile_image"`
}

type Topic struct {
	ID           uint    `json:"id"`
	Name         string  `json:"name"`
	ProfileImage *string `json:"profile_image"`
}

type Reason struct {
	ID     uint   `json:"id"`
	Detail string `json:"detail"`
}

type Thread struct {
	ID      uint    `json:"id"`
	Title   string  `json:"title"`
	Content *string `json:"content"`
	Image1  *string `json:"image_1"`
	Image2  *string `json:"image_2"`
	Image3  *string `json:"image_3"`
	Image4  *string `json:"image_4"`
	Image5  *string `json:"image_5"`
}

type Reply struct {
	ID      uint    `json:"id"`
	Content string  `json:"content"`
	Image   *string `json:"image"`
}

type Response struct {
	ID        uint      `json:"id"`
	Reporter  User      `json:"reporter"`
	Reason    Reason    `json:"reason"`
	Suspect   *User     `json:"suspect,omitempty"`
	Topic     *Topic    `json:"topic,omitempty"`
	Thread    *Thread   `json:"thread,omitempty"`
	Reply     *Reply    `json:"reply,omitempty"`
	Reviewed  *bool     `json:"reviewed,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}

func FromDomain(reportDomain *report.Domain, scope string) interface{} {
	if scope == "user" {
		return Response{
			Reporter: User{
				ID:           reportDomain.ReporterID,
				Username:     reportDomain.ReporterName,
				ProfileImage: reportDomain.ReporterProfileImage,
			},
			Reason: Reason{
				ID:     reportDomain.ReasonID,
				Detail: reportDomain.ReasonDetail,
			},
			Suspect: &User{
				ID:           *reportDomain.SuspectID,
				Username:     *reportDomain.SuspectUsername,
				ProfileImage: reportDomain.SuspectProfileImage,
			},
			CreatedAt: reportDomain.CreatedAt,
		}
	} else if scope == "topic" {
		return Response{
			Reporter: User{
				ID:           reportDomain.ReporterID,
				Username:     reportDomain.ReporterName,
				ProfileImage: reportDomain.ReporterProfileImage,
			},
			Reason: Reason{
				ID:     reportDomain.ReasonID,
				Detail: reportDomain.ReasonDetail,
			},
			Topic: &Topic{
				ID:           *reportDomain.TopicID,
				Name:         *reportDomain.TopicName,
				ProfileImage: reportDomain.TopicProfileImage,
			},
			CreatedAt: reportDomain.CreatedAt,
		}
	} else if scope == "thread" {
		return Response{
			Reporter: User{
				ID:           reportDomain.ReporterID,
				Username:     reportDomain.ReporterName,
				ProfileImage: reportDomain.ReporterProfileImage,
			},
			Reason: Reason{
				ID:     reportDomain.ReasonID,
				Detail: reportDomain.ReasonDetail,
			},
			Topic: &Topic{
				ID:           *reportDomain.TopicID,
				Name:         *reportDomain.TopicName,
				ProfileImage: reportDomain.TopicProfileImage,
			},
			Thread: &Thread{
				ID:      *reportDomain.ThreadID,
				Title:   *reportDomain.ThreadTitle,
				Content: reportDomain.ThreadContent,
				Image1:  reportDomain.ThreadImage1,
				Image2:  reportDomain.ThreadImage2,
				Image3:  reportDomain.ThreadImage3,
				Image4:  reportDomain.ThreadImage4,
				Image5:  reportDomain.ThreadImage5,
			},
			Reviewed:  reportDomain.Reviewed,
			CreatedAt: reportDomain.CreatedAt,
		}
	} else if scope == "reply" {
		return Response{
			Reporter: User{
				ID:           reportDomain.ReporterID,
				Username:     reportDomain.ReporterName,
				ProfileImage: reportDomain.ReporterProfileImage,
			},
			Reason: Reason{
				ID:     reportDomain.ReasonID,
				Detail: reportDomain.ReasonDetail,
			},
			Topic: &Topic{
				ID:           *reportDomain.TopicID,
				Name:         *reportDomain.TopicName,
				ProfileImage: reportDomain.TopicProfileImage,
			},
			Reply: &Reply{
				ID:      *reportDomain.ReplyID,
				Content: *reportDomain.ReplyContent,
				Image:   reportDomain.ReplyImage,
			},
			Reviewed:  reportDomain.Reviewed,
			CreatedAt: reportDomain.CreatedAt,
		}
	}

	return nil
}

func FromDomains(reportDomains *[]report.Domain, scope string) interface{} {
	responses := []Response{}
	if scope == "user" {
		for _, reportDomain := range *reportDomains {
			responses = append(responses, Response{
				Reporter: User{
					ID:           reportDomain.ReporterID,
					Username:     reportDomain.ReporterName,
					ProfileImage: reportDomain.ReporterProfileImage,
				},
				Reason: Reason{
					ID:     reportDomain.ReasonID,
					Detail: reportDomain.ReasonDetail,
				},
				Suspect: &User{
					ID:           *reportDomain.SuspectID,
					Username:     *reportDomain.SuspectUsername,
					ProfileImage: reportDomain.SuspectProfileImage,
				},
				CreatedAt: reportDomain.CreatedAt,
			})
		}
	} else if scope == "topic" {
		for _, reportDomain := range *reportDomains {
			responses = append(responses, Response{
				Reporter: User{
					ID:           reportDomain.ReporterID,
					Username:     reportDomain.ReporterName,
					ProfileImage: reportDomain.ReporterProfileImage,
				},
				Reason: Reason{
					ID:     reportDomain.ReasonID,
					Detail: reportDomain.ReasonDetail,
				},
				Topic: &Topic{
					ID:           *reportDomain.TopicID,
					Name:         *reportDomain.TopicName,
					ProfileImage: reportDomain.TopicProfileImage,
				},
				CreatedAt: reportDomain.CreatedAt,
			})
		}
	} else if scope == "thread" {
		for _, reportDomain := range *reportDomains {
			responses = append(responses, Response{
				Reporter: User{
					ID:           reportDomain.ReporterID,
					Username:     reportDomain.ReporterName,
					ProfileImage: reportDomain.ReporterProfileImage,
				},
				Reason: Reason{
					ID:     reportDomain.ReasonID,
					Detail: reportDomain.ReasonDetail,
				},
				Topic: &Topic{
					ID:           *reportDomain.TopicID,
					Name:         *reportDomain.TopicName,
					ProfileImage: reportDomain.TopicProfileImage,
				},
				Thread: &Thread{
					ID:      *reportDomain.ThreadID,
					Title:   *reportDomain.ThreadTitle,
					Content: reportDomain.ThreadContent,
					Image1:  reportDomain.ThreadImage1,
					Image2:  reportDomain.ThreadImage2,
					Image3:  reportDomain.ThreadImage3,
					Image4:  reportDomain.ThreadImage4,
					Image5:  reportDomain.ThreadImage5,
				},
				Reviewed:  reportDomain.Reviewed,
				CreatedAt: reportDomain.CreatedAt,
			})
		}
	} else if scope == "reply" {
		for _, reportDomain := range *reportDomains {
			responses = append(responses, Response{
				Reporter: User{
					ID:           reportDomain.ReporterID,
					Username:     reportDomain.ReporterName,
					ProfileImage: reportDomain.ReporterProfileImage,
				},
				Reason: Reason{
					ID:     reportDomain.ReasonID,
					Detail: reportDomain.ReasonDetail,
				},
				Topic: &Topic{
					ID:           *reportDomain.TopicID,
					Name:         *reportDomain.TopicName,
					ProfileImage: reportDomain.TopicProfileImage,
				},
				Reply: &Reply{
					ID:      *reportDomain.ReplyID,
					Content: *reportDomain.ReplyContent,
					Image:   reportDomain.ReplyImage,
				},
				Reviewed:  reportDomain.Reviewed,
				CreatedAt: reportDomain.CreatedAt,
			})
		}
	} else if scope == "reason" {
		for _, reportDomain := range *reportDomains {
			responses = append(responses, Response{
				Reason: Reason{
					ID:     reportDomain.ReasonID,
					Detail: reportDomain.ReasonDetail,
				},
			})
		}
	}

	return responses
}
