package op

import (
	"net/url"
)

type PasswordResetRequest struct {
	Email string
}

func (r *PasswordResetRequest) FormatToValues() url.Values {
	var values = url.Values{}

	values.Add("email", r.Email)

	return values
}

type PasswordResetResponse struct {
}

func (r *PasswordResetResponse) Parse(data []byte) error {
	return nil
}
