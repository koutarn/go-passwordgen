package generate

import (
	"fmt"
	"testing"

	generate "github.com/koutarn/go-passwordgen/pkg"
)

func TestGeneratePassword正常系(t *testing.T){
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!@#$%^&*()0123456789"
	var password,_ = generate.GeneratePassword(500,charset)
	fmt.Println(password)
}

func TestGeneratePassword文字列0(t *testing.T){
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!@#$%^&*()0123456789"
	var password,_ = generate.GeneratePassword(0,charset)
	fmt.Println(password)
}