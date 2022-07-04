package request

import "fgd/core/report"

type Reason struct {
	Detail string `json:"detail"`
}

func (r *Reason) ToDomain() *report.Domain {
	return &report.Domain{
		ReasonDetail: r.Detail,
	}
}
