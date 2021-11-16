package validator

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidator(t *testing.T) {
	type user struct {
		Username string `validate:"required"`
		Email    string `validate:"required,email"`
		Age      int    `validate:"required,gte=18,lte=130"`
	}

	u := &user{
		Username: "",
		Email:    "xxxxx",
		Age:      10,
	}

	msg := Verify(u)
	fmt.Println(msg)
	assert.Len(t, msg, 3)
}

func TestVerifyField(t *testing.T) {
	email := "xx"
	msg := VerifyField(email, "required,email")
	fmt.Println(msg)
	assert.NotEmpty(t, msg)
	assert.Contains(t, msg, "邮箱")
}

func TestForm(t *testing.T) {
	form := map[string]string{
		"username": "",
		"email":    "xx",
	}

	rule := map[string]string{
		"username": "required",
		"email":    "required,email",
	}

	var ret []string
	for k, v := range rule {
		msg := VerifyField(form[k], v)
		if msg != "" {
			ret = append(ret, k+msg)
		}
	}

	fmt.Println(ret)
	assert.Len(t, ret, 2)
}
