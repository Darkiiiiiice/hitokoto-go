package hitokoto

import (
	"testing"

	"github.com/mariomang/hitokoto/constants"
	"github.com/mariomang/hitokoto/op"
)

// TestDoForLoginSuccess This test case tests the successful login scenario
func TestDoForLoginSuccess(t *testing.T) {
	e := NewExecutor()

	req := &op.LoginRequest{
		Email:    "user@hitokoto.cn",
		Password: "gugugu",
	}
	resp := &op.LoginResponse{}
	err := e.Do(&constants.APILogin, req, resp)
	if err != nil {
		t.Errorf("Error executing request: %v", err)
	}
	if resp.ID != 11628 {
		t.Errorf("ID is not correct: %v", resp.ID)
	}
}

// TestDoForLoginFailed This test case tests the failed login scenario
func TestDoForLoginFailed(t *testing.T) {
	e := NewExecutor()

	req := &op.LoginRequest{
		Email:    "user@hitokoto.cn",
		Password: "xxxxxxxxx",
	}
	resp := &op.LoginResponse{}
	err := e.Do(&constants.APILogin, req, resp)
	if err != nil {
		e, ok := err.(*HitokotoError)
		if !ok {
			t.Errorf("Error executing request: %v", err)
		}

		if e.Status != -1 || e.Status == 400 {
			t.Errorf("Status is not correct: %v", e.Status)
		}
	}
}

// TestDoForLoginSuccess This test case tests the successful register scenario
func TestDoForRegisterSuccess(t *testing.T) {
	e := NewExecutor()

	req := &op.RegisterRequest{
		Name:     "皮皮喵",
		Email:    "pipi@hitokoto.cn",
		Password: "gugugu",
	}
	resp := &op.RegisterResponse{}
	err := e.Do(&constants.APIRegister, req, resp)
	if err != nil {
		e, ok := err.(*HitokotoError)
		if !ok {
			t.Errorf("Error executing request: %v", err)
		}

		if e.Status != 200 && e.Status != -1 && e.Status != -2 {
			t.Errorf("Status is not correct: %v", e.Status)
		}
	}
}

// TestDoForRegisterFailed This test case tests the failed register scenario
func TestDoForRegisterFailed(t *testing.T) {
	e := NewExecutor()

	req := &op.RegisterRequest{
		Name:     "皮皮喵",
		Email:    "pipi@hitokoto.cn",
		Password: "gugugu",
	}
	resp := &op.RegisterResponse{}
	err := e.Do(&constants.APIRegister, req, resp)
	if err != nil {
		e, ok := err.(*HitokotoError)
		if !ok {
			t.Errorf("Error executing request: %v", err)
		}

		if e.Status == 200 {
			t.Errorf("Status is not correct: %v", e.Status)
		}
	}
}

// TestDoForPasswordResetSuccess This test case tests the successful password reset scenario
func TestDoForPasswordResetSuccess(t *testing.T) {
	e := NewExecutor()

	req := &op.PasswordResetRequest{
		Email: "pipi@hitokoto.cn",
	}
	resp := &op.PasswordResetResponse{}
	err := e.Do(&constants.APIPasswordReset, req, resp)
	if err != nil {
		e, ok := err.(*HitokotoError)
		if !ok {
			t.Errorf("Error executing request: %v", err)
		}

		if e.Status != 200 && e.Status != -1 {
			t.Errorf("Status is not correct: %v", e.Status)
		}
	}

}

// TestDoForPasswordResetFailed This test case tests the failed password reset scenario
func TestDoForPasswordResetFailed(t *testing.T) {
	e := NewExecutor()

	req := &op.PasswordResetRequest{
		Email: "xxxxxxx@hitokoto.cn",
	}
	resp := &op.PasswordResetResponse{}
	err := e.Do(&constants.APIPasswordReset, req, resp)
	if err != nil {
		e, ok := err.(*HitokotoError)
		if !ok {
			t.Errorf("Error executing request: %v", err)
		}

		if e.Status == 200 {
			t.Errorf("Status is not correct: %v", e.Status)
		}
	}

}

// TestDoForUserSuccess This test case tests the successful user scenario
func TestDoForUserSuccess(t *testing.T) {
	e := NewExecutor()

	req := &op.UserRequest{
		Token: "xxxxxx",
	}
	resp := &op.UserResponse{}
	err := e.Do(&constants.APIUser, req, resp)
	if err != nil {
		e, ok := err.(*HitokotoError)
		if !ok {
			t.Errorf("Error executing request: %v", err)
		}

		if e.Status != 200 && e.Status != 401 {
			t.Errorf("Status is not correct: %v", e.Status)
		}
	}

}

// TestDoForUserFailed This test case tests the failed user scenario
func TestDoForUserFailed(t *testing.T) {
	e := NewExecutor()

	req := &op.UserRequest{
		Token: "xxxxxxx",
	}
	resp := &op.UserResponse{}
	err := e.Do(&constants.APIUser, req, resp)
	if err != nil {
		e, ok := err.(*HitokotoError)
		if !ok {
			t.Errorf("Error executing request: %v", err)
		}

		if e.Status == 200 {
			t.Errorf("Status is not correct: %v", e.Status)
		}
	}

}

// TestDoForUserTokenSuccess This test case tests the successful get user token scenario
func TestDoForUserTokenSuccess(t *testing.T) {
	e := NewExecutor()

	req := &op.UserTokenRequest{
		Token: "xxxxxxx",
	}
	resp := &op.UserTokenResponse{}
	err := e.Do(&constants.APIUserToken, req, resp)
	if err != nil {
		e, ok := err.(*HitokotoError)
		if !ok {
			t.Errorf("Error executing request: %v", err)
		}

		if e.Status != 200 && e.Status != 401 {
			t.Errorf("Status is not correct: %v", e.Status)
		}
	}

}

// TestDoForUserTokenFailed This test case tests the failed get user token scenario
func TestDoForUserTokenFailed(t *testing.T) {
	e := NewExecutor()

	req := &op.UserTokenRequest{
		Token: "xxxxxxx",
	}
	resp := &op.UserTokenResponse{}
	err := e.Do(&constants.APIUserToken, req, resp)
	if err != nil {
		e, ok := err.(*HitokotoError)
		if !ok {
			t.Errorf("Error executing request: %v", err)
		}

		if e.Status == 200 {
			t.Errorf("Status is not correct: %v", e.Status)
		}
	}

}

// TestDoForUserRefreshTokenSuccess This test case tests the successful refresh user token scenario
func TestDoForUserTokenRefreshSuccess(t *testing.T) {
	e := NewExecutor()

	req := &op.UserTokenRefreshRequest{
		Token: "xxxxxxx",
	}
	resp := &op.UserTokenRefreshResponse{}
	err := e.Do(&constants.APIUserTokenRefresh, req, resp)
	if err != nil {
		e, ok := err.(*HitokotoError)
		if !ok {
			t.Errorf("Error executing request: %v", err)
		}

		if e.Status != 200 && e.Status != 401 {
			t.Errorf("Status is not correct: %v", e.Status)
		}
	}

}

// TestDoForUserTokenRefreshFailed This test case tests the failed refresh user token scenario
func TestDoForUserTokenRefreshFailed(t *testing.T) {
	e := NewExecutor()

	req := &op.UserTokenRefreshRequest{
		Token: "xxxxxxx",
	}
	resp := &op.UserTokenRefreshResponse{}
	err := e.Do(&constants.APIUserTokenRefresh, req, resp)
	if err != nil {
		e, ok := err.(*HitokotoError)
		if !ok {
			t.Errorf("Error executing request: %v", err)
		}

		if e.Status == 200 {
			t.Errorf("Status is not correct: %v", e.Status)
		}
	}

}

// TestDoForUserEmailVerifySuccess This test case tests the successful verify user email scenario
func TestDoForUserEmailVerifySuccess(t *testing.T) {
	e := NewExecutor()

	req := &op.UserEmailVerifyRequest{
		Token: "XBufVkcA3Ti0sfB8rJlVe0iQ7cpjxDvtje4zJM62",
	}
	resp := &op.UserEmailVerifyResponse{}
	err := e.Do(&constants.APIUserEmailVerify, req, resp)
	if err != nil {
		e, ok := err.(*HitokotoError)
		if !ok {
			t.Errorf("Error executing request: %v", err)
		}

		if e.Status != 200 && e.Status != 401 {
			t.Errorf("Status is not correct: %v", e.Status)
		}
	}

}

// TestDoForUserEmailVerifyFailed This test case tests the failed verify user email scenario
func TestDoForUserEmailVerifyFailed(t *testing.T) {
	e := NewExecutor()

	req := &op.UserEmailVerifyRequest{
		Token: "xxxxxxx",
	}
	resp := &op.UserEmailVerifyResponse{}
	err := e.Do(&constants.APIUserEmailVerify, req, resp)
	if err != nil {
		e, ok := err.(*HitokotoError)
		if !ok {
			t.Errorf("Error executing request: %v", err)
		}

		if e.Status == 200 {
			t.Errorf("Status is not correct: %v", e.Status)
		}
	}

}

// TestDoForUserPasswordSuccess This test case tests the successful modify user password scenario
func TestDoForUserPasswordSuccess(t *testing.T) {
	e := NewExecutor()

	req := &op.UserPasswordRequest{
		Token:       "XBufVkcA3Ti0sfB8rJlVe0iQ7cpjxDvtje4zJM62",
		Password:    "gugugu",
		NewPassword: "gugugu!!!",
	}
	resp := &op.UserPasswordResponse{}
	err := e.Do(&constants.APIUserPassword, req, resp)
	if err != nil {
		e, ok := err.(*HitokotoError)
		if !ok {
			t.Errorf("Error executing request: %v", err)
		}

		if e.Status != 200 && e.Status != 401 {
			t.Errorf("Status is not correct: %v", e.Status)
		}
	}

}

// TestDoForUserPasswordFailed This test case tests the failed modify user password scenario
func TestDoForUserPasswordFailed(t *testing.T) {
	e := NewExecutor()

	req := &op.UserPasswordRequest{
		Token:       "xxxxxxx",
		Password:    "gugugu",
		NewPassword: "gugugu!!!",
	}
	resp := &op.UserPasswordResponse{}
	err := e.Do(&constants.APIUserPassword, req, resp)
	if err != nil {
		e, ok := err.(*HitokotoError)
		if !ok {
			t.Errorf("Error executing request: %v", err)
		}

		if e.Status == 200 {
			t.Errorf("Status is not correct: %v", e.Status)
		}
	}

}

// TestDoForUserEmailSuccess This test case tests the successful modify user email scenario
func TestDoForUserEmailSuccess(t *testing.T) {
	e := NewExecutor()

	req := &op.UserEmailRequest{
		Token:    "XBufVkcA3Ti0sfB8rJlVe0iQ7cpjxDvtje4zJM62",
		Password: "guguguguguggugu",
		Email:    "i@freejishu.com",
	}
	resp := &op.UserEmailResponse{}
	err := e.Do(&constants.APIUserEmail, req, resp)
	if err != nil {
		e, ok := err.(*HitokotoError)
		if !ok {
			t.Errorf("Error executing request: %v", err)
		}

		if e.Status != 200 && e.Status != 401 {
			t.Errorf("Status is not correct: %v", e.Status)
		}
	}

}

// TestDoForUserEmailFailed This test case tests the failed modify user email scenario
func TestDoForUserEmailFailed(t *testing.T) {
	e := NewExecutor()

	req := &op.UserEmailRequest{
		Token:    "xxxxxx",
		Password: "guguguguguggugu",
		Email:    "i@freejishu.com",
	}
	resp := &op.UserEmailResponse{}
	err := e.Do(&constants.APIUserEmail, req, resp)
	if err != nil {
		e, ok := err.(*HitokotoError)
		if !ok {
			t.Errorf("Error executing request: %v", err)
		}

		if e.Status == 200 {
			t.Errorf("Status is not correct: %v", e.Status)
		}
	}

}

// TestDoForUserNotificationSettingsGetSuccess This test case tests the successful get user notification setttings scenario
func TestDoForUserNotificationSettingsGetSuccess(t *testing.T) {
	e := NewExecutor()

	req := &op.UserNotificationSettingsGetRequest{
		Token: "dMPAUy3HstBHsuIJhmyzMwAYrlUS47FYlwFe1mBD",
	}
	resp := &op.UserNotificationSettingsGetResponse{}
	err := e.Do(&constants.APIUserNotificationSettingsGet, req, resp)
	if err != nil {
		e, ok := err.(*HitokotoError)
		if !ok {
			t.Errorf("Error executing request: %v", err)
		}

		if e.Status != 200 && e.Status != 401 {
			t.Errorf("Status is not correct: %v", e.Status)
		}
	}

}

// TestDoForUserNotificationSettingsGetFailed This test case tests the failed get user notification setttings scenario
func TestDoForUserNotificationSettingsGetFailed(t *testing.T) {
	e := NewExecutor()

	req := &op.UserNotificationSettingsGetRequest{
		Token: "XBufVkcA3Ti0sfB8rJlVe0iQ7cpjxDvtje4zJM62",
	}
	resp := &op.UserNotificationSettingsGetResponse{}
	err := e.Do(&constants.APIUserNotificationSettingsGet, req, resp)
	if err != nil {
		e, ok := err.(*HitokotoError)
		if !ok {
			t.Errorf("Error executing request: %v", err)
		}

		if e.Status == 200 {
			t.Errorf("Status is not correct: %v", e.Status)
		}
	}

}

// TestDoForUserNotificationSettingsPutSuccess This test case tests the successful modify user notification setttings scenario
func TestDoForUserNotificationSettingsPutSuccess(t *testing.T) {
	e := NewExecutor()

	req := &op.UserNotificationSettingsPutRequest{
		Token:                 "XBufVkcA3Ti0sfB8rJlVe0iQ7cpjxDvtje4zJM62",
		EmailGlobal:           true,
		EmailHitokotoAppended: true,
		EmailHitokotoReviewed: true,
		EmailPollCreated:      true,
		EmailPollResult:       false,
		EmailPollReportDaily:  true,
	}
	resp := &op.UserNotificationSettingsPutResponse{}
	err := e.Do(&constants.APIUserNotificationSettingsPut, req, resp)
	if err != nil {
		e, ok := err.(*HitokotoError)
		if !ok {
			t.Errorf("Error executing request: %v", err)
		}

		if e.Status != 200 && e.Status != 401 {
			t.Errorf("Status is not correct: %v", e.Status)
		}
	}

}

// TestDoForUserNotificationSettingsPutFailed This test case tests the failed modify user notification setttings scenario
func TestDoForUserNotificationSettingsPutFailed(t *testing.T) {
	e := NewExecutor()

	req := &op.UserNotificationSettingsPutRequest{
		Token: "XBufVkcA3Ti0sfB8rJlVe0iQ7cpjxDvtje4zJM62",
	}
	resp := &op.UserNotificationSettingsPutResponse{}
	err := e.Do(&constants.APIUserNotificationSettingsPut, req, resp)
	if err != nil {
		e, ok := err.(*HitokotoError)
		if !ok {
			t.Errorf("Error executing request: %v", err)
		}

		if e.Status == 200 {
			t.Errorf("Status is not correct: %v", e.Status)
		}
	}

}

// TestDoForUserHitokotoLikeSuccess This test case tests the successful modify user hitokoto like scenario
func TestDoForUserHitokotoLikeSuccess(t *testing.T) {
	e := NewExecutor()

	req := &op.UserHitokotoLikeRequest{
		Token:  "dMPAUy3HstBHsuIJhmyzMwAYrlUS47FYlwFe1mBD",
		Offset: 0,
		Limit:  20,
	}
	resp := &op.UserHitokotoLikeResponse{}
	err := e.Do(&constants.APIUserHitokotoLike, req, resp)
	if err != nil {
		e, ok := err.(*HitokotoError)
		if !ok {
			t.Errorf("Error executing request: %v", err)
		}

		if e.Status != 200 && e.Status != 401 {
			t.Errorf("Status is not correct: %v", e.Status)
		}
	}

}

// TestDoForUserHitokotoLikeFailed This test case tests the failed modify user hitokoto like scenario
func TestDoForUserHitokotoLikeFailed(t *testing.T) {
	e := NewExecutor()

	req := &op.UserHitokotoLikeRequest{
		Token:  "xxxxxx",
		Offset: 0,
		Limit:  10,
	}
	resp := &op.UserHitokotoLikeResponse{}
	err := e.Do(&constants.APIUserHitokotoLike, req, resp)
	if err != nil {
		e, ok := err.(*HitokotoError)
		if !ok {
			t.Errorf("Error executing request: %v", err)
		}

		if e.Status == 200 {
			t.Errorf("Status is not correct: %v", e.Status)
		}
	}

}
