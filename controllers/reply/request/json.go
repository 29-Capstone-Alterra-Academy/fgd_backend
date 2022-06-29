package request

import "fgd/core/reply"

type Reply struct {
	Content string `json:"content"`
	Image   string `json:"image,omitempty"`
}

func (r *Reply) ToDomain() *reply.Domain {
	return &reply.Domain{
		Image:   &r.Image,
		Content: r.Content,
	}
}
