package main

import (
	"fmt"
	"net/http/httptest"
	"testing"

	auth "github.com/casdoor/casdoor-go-sdk/casdoorsdk"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

var callbackUrl = "http://localhost:9000/callback"

func init() {
	auth.InitConfig(CasdoorEndpoint, ClientId, ClientSecret, JwtPublicKey, CasdoorOrganization, CasdoorApplication)
}

func TestCallBackHandler(t *testing.T) {
	app := fiber.New()

	db = connectDB()
	db.AutoMigrate(&User{})

	app.Use(cors.New())
	app.Use(recover.New())

	app.Get("/callback", SignIn)

	t.Run("No query", func(t *testing.T) {
		resp, _ := app.Test(httptest.NewRequest("GET", fmt.Sprintf("%s", callbackUrl), nil))
		if resp.StatusCode != fiber.StatusUnauthorized {
			t.Errorf("expected %d got %d", fiber.StatusUnauthorized, resp.StatusCode)
		}
	})
	// t.Run("Invalid query", func(t *testing.T) {})
	// t.Run("code missing", func(t *testing.T) {})
	// t.Run("state missing", func(t *testing.T) {})
	// t.Run("used code", func(t *testing.T) {})
}
