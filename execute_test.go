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

// TestDoForUserHitokotoLikeSuccess This test case tests the successful user hitokoto like scenario
func TestDoForUserHitokotoLikeSuccess(t *testing.T) {
	e := NewExecutor()

	req := &op.UserHitokotoLikeRequest{
		Token:  "XBufVkcA3Ti0sfB8rJlVe0iQ7cpjxDvtje4zJM62",
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

// TestDoForUserHitokotoLikeFailed This test case tests the failed user hitokoto like scenario
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

// TestDoForUserHitokotoSummarySuccess This test case tests the successful user hitokoto summary scenario
func TestDoForUserHitokotoSummarySuccess(t *testing.T) {
	e := NewExecutor()

	req := &op.UserHitokotoSummaryRequest{
		Token: "XBufVkcA3Ti0sfB8rJlVe0iQ7cpjxDvtje4zJM62",
	}
	resp := &op.UserHitokotoSummaryResponse{}
	err := e.Do(&constants.APIUserHitokotoSummary, req, resp)
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

// TestDoForUserHitokotoSummaryFailed This test case tests the failed user hitokoto summary scenario
func TestDoForUserHitokotoSummaryFailed(t *testing.T) {
	e := NewExecutor()

	req := &op.UserHitokotoSummaryRequest{
		Token: "xxxx",
	}
	resp := &op.UserHitokotoSummaryResponse{}
	err := e.Do(&constants.APIUserHitokotoSummary, req, resp)
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

// TestDoForUserHitokotoHistorySuccess This test case tests the successful user hitokoto history scenario
func TestDoForUserHitokotoHistorySuccess(t *testing.T) {
	e := NewExecutor()

	req := &op.UserHitokotoHistoryRequest{
		Token:  "XBufVkcA3Ti0sfB8rJlVe0iQ7cpjxDvtje4zJM62",
		Offset: 0,
		Limit:  20,
	}
	resp := &op.UserHitokotoHistoryResponse{}
	err := e.Do(&constants.APIUserHitokotoHistory, req, resp)
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

// TestDoForUserHitokotoHistoryFailed This test case tests the failed user hitokoto history scenario
func TestDoForUserHitokotoHistoryFailed(t *testing.T) {
	e := NewExecutor()

	req := &op.UserHitokotoHistoryRequest{
		Token: "xxxx",
	}
	resp := &op.UserHitokotoHistoryResponse{}
	err := e.Do(&constants.APIUserHitokotoHistory, req, resp)
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

// TestDoForUserHitokotoHistoryPendingSuccess This test case tests the successful user hitokoto history pending scenario
func TestDoForUserHitokotoHistoryPendingSuccess(t *testing.T) {
	e := NewExecutor()

	req := &op.UserHitokotoHistoryPendingRequest{
		Token:  "XBufVkcA3Ti0sfB8rJlVe0iQ7cpjxDvtje4zJM62",
		Offset: 0,
		Limit:  20,
	}
	resp := &op.UserHitokotoHistoryPendingResponse{}
	err := e.Do(&constants.APIUserHitokotoHistoryPending, req, resp)
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

// TestDoForUserHitokotoHistoryPendingFailed This test case tests the failed user hitokoto history scenario
func TestDoForUserHitokotoHistoryPendingFailed(t *testing.T) {
	e := NewExecutor()

	req := &op.UserHitokotoHistoryPendingRequest{
		Token: "xxxx",
	}
	resp := &op.UserHitokotoHistoryPendingResponse{}
	err := e.Do(&constants.APIUserHitokotoHistoryPending, req, resp)
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

// TestDoForUserHitokotoHistoryRefuseSuccess This test case tests the successful user hitokoto history refuse scenario
func TestDoForUserHitokotoHistoryRefuseSuccess(t *testing.T) {
	e := NewExecutor()

	req := &op.UserHitokotoHistoryRefuseRequest{
		Token:  "XBufVkcA3Ti0sfB8rJlVe0iQ7cpjxDvtje4zJM62",
		Offset: 0,
		Limit:  20,
	}
	resp := &op.UserHitokotoHistoryRefuseResponse{}
	err := e.Do(&constants.APIUserHitokotoHistoryRefuse, req, resp)
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

// TestDoForUserHitokotoHistoryRefuseFailed This test case tests the failed user hitokoto history refuse scenario
func TestDoForUserHitokotoHistoryRefuseFailed(t *testing.T) {
	e := NewExecutor()

	req := &op.UserHitokotoHistoryRefuseRequest{
		Token: "xxxx",
	}
	resp := &op.UserHitokotoHistoryRefuseResponse{}
	err := e.Do(&constants.APIUserHitokotoHistoryRefuse, req, resp)
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

// TestDoForUserHitokotoHistoryAcceptSuccess This test case tests the successful user hitokoto history accept scenario
func TestDoForUserHitokotoHistoryAcceptSuccess(t *testing.T) {
	e := NewExecutor()

	req := &op.UserHitokotoHistoryAcceptRequest{
		Token:  "XBufVkcA3Ti0sfB8rJlVe0iQ7cpjxDvtje4zJM62",
		Offset: 0,
		Limit:  20,
	}
	resp := &op.UserHitokotoHistoryAcceptResponse{}
	err := e.Do(&constants.APIUserHitokotoHistoryAccept, req, resp)
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

// TestDoForUserHitokotoHistoryAcceptFailed This test case tests the failed user hitokoto history accept scenario
func TestDoForUserHitokotoHistoryAcceptFailed(t *testing.T) {
	e := NewExecutor()

	req := &op.UserHitokotoHistoryAcceptRequest{
		Token: "xxxx",
	}
	resp := &op.UserHitokotoHistoryAcceptResponse{}
	err := e.Do(&constants.APIUserHitokotoHistoryAccept, req, resp)
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

// TestDoForLikeGetSuccess This test case tests the successful get like scenario
func TestDoForLikeGetSuccess(t *testing.T) {
	e := NewExecutor()

	req := &op.LikeGetRequest{
		SentenceUuid: "34662c58-8eba-4757-a637-c7c11e9f537e",
	}
	resp := &op.LikeGetResponse{}
	err := e.Do(&constants.APILikeGet, req, resp)
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

// TestDoForLikeGetFailed This test case tests the failed get like scenario
func TestDoForLikeGetFailed(t *testing.T) {
	e := NewExecutor()

	req := &op.LikeGetRequest{
		SentenceUuid: "xxxxxx",
	}
	resp := &op.LikeGetResponse{}
	err := e.Do(&constants.APILikeGet, req, resp)
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

// TestDoForLikePostSuccess This test case tests the successful like post scenario
func TestDoForLikePostSuccess(t *testing.T) {
	e := NewExecutor()

	req := &op.LikePostRequest{
		SentenceUuid: "cc5d4eca-b4fb-4da8-aa1c-7f69d8cea9fb",
	}
	resp := &op.LikePostResponse{}
	err := e.Do(&constants.APILikePost, req, resp)
	if err != nil {
		e, ok := err.(*HitokotoError)
		if !ok {
			t.Errorf("Error executing request: %v", err)
		}

		if e.Status != 200 && e.Status != 401 && e.Status != -1 {
			t.Errorf("Status is not correct: %v", e.Status)
		}
	}

}

// TestDoForLikePostFailed This test case tests the failed like post scenario
func TestDoForLikePostFailed(t *testing.T) {
	e := NewExecutor()

	req := &op.LikePostRequest{
		SentenceUuid: "xxxxxx",
	}
	resp := &op.LikePostResponse{}
	err := e.Do(&constants.APILikePost, req, resp)
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

// TestDoForLikeCancelSuccess This test case tests the successful like cancel scenario
func TestDoForLikeCancelSuccess(t *testing.T) {
	e := NewExecutor()

	req := &op.LikeCancelRequest{
		Token:        "XBufVkcA3Ti0sfB8rJlVe0iQ7cpjxDvtje4zJM62",
		SentenceUuid: "cc5d4eca-b4fb-4da8-aa1c-7f69d8cea9fb",
	}
	resp := &op.LikeCancelResponse{}
	err := e.Do(&constants.APILikeCancel, req, resp)
	if err != nil {
		e, ok := err.(*HitokotoError)
		if !ok {
			t.Errorf("Error executing request: %v", err)
		}

		if e.Status != 200 && e.Status != 401 && e.Status != -1 {
			t.Errorf("Status is not correct: %v", e.Status)
		}
	}

}

// TestDoForLikeCancelFailed This test case tests the failed like cancel scenario
func TestDoForLikeCancelFailed(t *testing.T) {
	e := NewExecutor()

	req := &op.LikeCancelRequest{
		Token:        "xxxxxx",
		SentenceUuid: "xxxxxx",
	}
	resp := &op.LikeCancelResponse{}
	err := e.Do(&constants.APILikeCancel, req, resp)
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

// TestDoForMarkSuccess This test case tests the successful mark scenario
func TestDoForMarkSuccess(t *testing.T) {
	e := NewExecutor()

	req := &op.MarkRequest{
		Token: "XBufVkcA3Ti0sfB8rJlVe0iQ7cpjxDvtje4zJM62",
	}
	resp := &op.MarkResponse{}
	err := e.Do(&constants.APIMark, req, resp)
	if err != nil {
		e, ok := err.(*HitokotoError)
		if !ok {
			t.Errorf("Error executing request: %v", err)
		}

		if e.Status != 200 && e.Status != 401 && e.Status != -1 {
			t.Errorf("Status is not correct: %v", e.Status)
		}
	}

}

// TestDoMarkFailed This test case tests the failed mark scenario
func TestDoForMarkFailed(t *testing.T) {
	e := NewExecutor()

	req := &op.MarkRequest{
		Token: "xxxxxx",
	}
	resp := &op.MarkResponse{}
	err := e.Do(&constants.APIMark, req, resp)
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
