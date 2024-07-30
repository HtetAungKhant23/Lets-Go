package api_test

import (
	"testing"

	"github.com/htetaungkhant/go/crypto/api"
)

func TestApiCall(t *testing.T) {
	_, err := api.GetRate("")
	if err == nil {
		t.Error("No error found.")
	}
}
