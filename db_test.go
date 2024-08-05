package main

import (
	"testing"
)

var dummy_uuid []string = []string{
	"b014cd21-2f20-450f-9ec6-e16dd159f798",
	"1a4dd704-da32-4fe2-b0e6-760a0e73dd0e",
}

var sets = []struct {
	name  string
	input UserInput
}{
	{"no input", UserInput{}},
	{"no sso id", UserInput{Name: "NoSSO"}},
	{"no name", UserInput{SsoId: dummy_uuid[0]}},
	{"input with leading and trailing spaces", UserInput{SsoId: dummy_uuid[0] + "    ", Name: "  Valid   "}},
	{"duplicate sso id", UserInput{SsoId: dummy_uuid[0], Name: "AnotherValid"}},
	{"valid input", UserInput{SsoId: dummy_uuid[1], Name: "Valid"}},
}

func TestCreateUser(t *testing.T) {
	db := connectDB()
	for _, tt := range sets {
		t.Run(tt.name, func(t *testing.T) {
			if s := CreateUser(db, &tt.input); s != nil {
				t.Errorf(">>> ERROR: GOT error: %s", s.Error())
			}
		})
	}

	defer func() {
		db.Exec("DELETE FROM users WHERE sso_id IN ('" + dummy_uuid[0] + "', '" + dummy_uuid[1] + "')")
		sql, _ := db.DB()
		sql.Close()
	}()
}
