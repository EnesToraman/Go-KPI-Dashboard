package utils

import (
	"testing"
)

// TestHashPassword calls utils.HashPassword with a password, checks
// for any error.
func TestHashPassword(t *testing.T) {
	res, err := HashPassword("password")
	if err != nil {
		t.Fatalf(`HashPassword("password") = %v, %q`, res, err)
	}
}

// TestCheckPassword calls utils.HashPassword, creates a hashed password,
// then calls utils.CheckPassword and compares, checks for any error.
func TestCheckPassword(t *testing.T) {
	res, _ := HashPassword("password")
	err := CheckPassword("password", res)
	if err != nil {
		t.Fatalf(`CheckPassword("password", res) = %v`, err)
	}
}
