package request

import "fgd/core/thread"

type Thread struct {
	Title   string `json:"title"`
	Content string `json:"content,omitempty"`
	Image1  string `json:"image_1,omitempty"`
	Image2  string `json:"image_2,omitempty"`
	Image3  string `json:"image_3,omitempty"`
	Image4  string `json:"image_4,omitempty"`
	Image5  string `json:"image_5,omitempty"`
}

func (r *Thread) ToDomain() *thread.Domain {
	return &thread.Domain{
		Image1:  &r.Image1,
		Image2:  &r.Image2,
		Image3:  &r.Image3,
		Image4:  &r.Image4,
		Image5:  &r.Image5,
		Title:   r.Title,
		Content: r.Content,
	}
}
