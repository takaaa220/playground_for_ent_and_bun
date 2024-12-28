package blob_bun

import "testing"

func TestCreateUser(t *testing.T) {
	tests := map[string]struct {
	}{
		"Test CreateUser": {},
	}
	for name, _ := range tests {
		t.Run(name, func(t *testing.T) {
			CreateUser()
		})
	}
}
