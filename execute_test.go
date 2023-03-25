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
