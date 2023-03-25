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

		if e.Status != -1 {
			t.Errorf("Status is not correct: %v", e.Status)
		}
	}

	fmt.Printf("resp: %v", resp)
}
