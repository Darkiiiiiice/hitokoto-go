package op

import (
	"net/url"

	"github.com/mariomang/hitokoto/constants"
	"github.com/valyala/fastjson"
)

type LoginRequest struct {
	Email    string
	Password string
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

type LoginCommand struct {
	Request  *LoginRequest
	Response *LoginResponse
	Api      constants.API
}

func NewLoginCommand(req *LoginRequest) *LoginCommand {
	return &LoginCommand{
		Request: req,
		Api:     constants.APILogin,
	}
}

func (lc *LoginCommand) Values() url.Values {
	var values = url.Values{}

	if lc.Request == nil {
		return values
	}

	values.Add("email", lc.Request.Email)
	values.Add("password", lc.Request.Password)

	return values
}

func (lc *LoginCommand) Parse(data []byte) error {
	fastjson.Parse(string(data))
	return nil
}

func (lc *LoginCommand) API() *constants.API {
	return &lc.Api
}
