package hitokoto

import (
	"fmt"
	"testing"

	"github.com/mariomang/hitokoto/constants"
	"github.com/mariomang/hitokoto/op"
)

// TestDoForLoginSuccess This test case tests the successful login scenario
func TestDoForLoginSuccess(t *testing.T) {
	e := NewExecutor("token")

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

	fmt.Printf("resp: %v", resp)
}

// TestDoForLoginFailed This test case tests the failed login scenario
func TestDoForLoginFailed(t *testing.T) {
	e := NewExecutor("token")

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

	fmt.Printf("resp: %v", resp)
}

// TestDoForLoginSuccess This test case tests the successful register scenario
func TestDoForRegisterSuccess(t *testing.T) {
	e := NewExecutor("token")

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

// TestDoForLoginFailed This test case tests the failed register scenario
func TestDoForRegisterFailed(t *testing.T) {
	e := NewExecutor("token")

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

	fmt.Printf("resp: %v", resp)
}
