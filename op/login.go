package op

import (
	"net/url"

	"github.com/valyala/fastjson"
)

type LoginRequest struct {
	Email    string
	Password string
}

func (lr *LoginRequest) FormatToValues() url.Values {
	var values = url.Values{}

	values.Add("email", lr.Email)
	values.Add("password", lr.Password)

	return values
}

type LoginResponse struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	Email           string `json:"email"`
	IsSuspended     int    `json:"is_suspended"`
	IsAdmin         int    `json:"is_admin"`
	IsReviewer      int    `json:"is_reviewer"`
	EmailVerifiedAt string `json:"email_verified_at"`
	CreatedFrom     string `json:"created_from"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
	Token           string `json:"token"`
}

func (lr *LoginResponse) Parse(data []byte) error {
	v, err := fastjson.ParseBytes(data)
	if err != nil {
		return err
	}

	list, err := v.Array()
	if err != nil {
		return err
	}
	if len(list) > 0 {
		d := list[0]
		lr.ID = d.GetInt("id")
		lr.Name = string(d.GetStringBytes("name"))
		lr.Email = string(d.GetStringBytes("email"))
		lr.IsSuspended = d.GetInt("is_suspended")
		lr.IsAdmin = d.GetInt("is_admin")
		lr.IsReviewer = d.GetInt("is_reviewer")
		lr.EmailVerifiedAt = string(d.GetStringBytes("email_verified_at"))
		lr.CreatedFrom = string(d.GetStringBytes("created_from"))
		lr.CreatedAt = string(d.GetStringBytes("created_at"))
		lr.UpdatedAt = string(d.GetStringBytes("updated_at"))
		lr.Token = string(d.GetStringBytes("token"))
	}

	return nil
}
