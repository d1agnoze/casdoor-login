package main

import (
	"testing"

	"github.com/google/uuid"
)

var dummy_uuid []string = []string{
	"b014cd21-2f20-450f-9ec6-e16dd159f798",
}

var sets = []struct {
	name  string
	input UserInput
}{
	{"no input", UserInput{}},
	{"no sso id", UserInput{Name: "NoSSO"}},
	{"no name", UserInput{SsoId: dummy_uuid[0]}},
	{"input with leading and trailing spaces", UserInput{SsoId: dummy_uuid[0] + "    ", Name: "  Valid   "}},
	{"valid input", UserInput{SsoId: uuid.New().String(), Name: "Valid"}},
}

func TestCreateUser(t *testing.T) {
	db := connectDB()
	for _, tt := range sets {
		t.Run(tt.name, func(t *testing.T) {
			if s := CreateUser(db, &tt.input); s == nil {
				t.Errorf(">>> ERROR: GOT nil error")
			}
		})
	}
}
