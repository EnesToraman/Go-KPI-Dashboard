package utils

import (
	"errors"
	"os"
	"os/exec"
	"testing"
)

// TestCheckError calls utils.CheckError with a new error, tests the process.
// Check out: https://talks.golang.org/2014/testing.slide#23
func TestCheckError(t *testing.T) {
	err := errors.New("error")
	if os.Getenv("BE_ERROR") == "1" {
		CheckError(err)
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=TestCheckError")
	cmd.Env = append(os.Environ(), "BE_ERROR=1")
	err = cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatalf("Process ran with err %v, want exit status 1", err)
}
