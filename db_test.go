package main

import (
	"casdoordemo/err"
	"testing"
)

var dummy_uuid []string = []string{
	"b014cd21-2f20-450f-9ec6-e16dd159f798",
	"1a4dd704-da32-4fe2-b0e6-760a0e73dd0e",
}

var sets = []struct {
	name   string
	input  UserInput
	output error
}{
	{"no input", UserInput{}, &err.InvalidInputError{}},
	{"no sso id", UserInput{Name: "NoSSO"}, &err.InvalidInputError{}},
	{"no name", UserInput{SsoId: dummy_uuid[0]}, &err.InvalidInputError{}},
	{"input with leading and trailing spaces", UserInput{SsoId: dummy_uuid[0] + "    ", Name: "  Valid   "}, nil},
	{"duplicate sso id", UserInput{SsoId: dummy_uuid[0], Name: "AnotherValid"}, nil},
	{"valid input", UserInput{SsoId: dummy_uuid[1], Name: "Valid"}, nil},
}

func TestCreateUser(t *testing.T) {
	db := connectDB()
	for _, tt := range sets {
		t.Run(tt.name, func(t *testing.T) {
			if s := CreateUser(db, &tt.input); s != nil && s.Error() != tt.output.Error() {
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
