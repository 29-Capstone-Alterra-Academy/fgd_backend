package request

import (
	"fgd/core/verify"
)

type Verify struct {
	Code string `json:"code"`
}

func (r *Verify) ToDomain() verify.Domain {
	return verify.Domain{
		Code: r.Code,
	}
}

type PasswordReset struct {
	NewPassword string `json:"new_password"`
}
