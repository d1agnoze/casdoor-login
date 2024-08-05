package main

import (
	_ "embed"

	auth "github.com/casdoor/casdoor-go-sdk/casdoorsdk"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"gorm.io/gorm"
)

var CasdoorEndpoint = "http://localhost:8000"
var ClientId = "b63455e4474ea776e1e6"
var ClientSecret = "3b9ebe3b22556398acad0dd3b627cda5fcdc0f2c"
var CasdoorOrganization = "nvstudio"
var CasdoorApplication = "demo"

//go:embed token_jwt_key.pem
var JwtPublicKey string

var db *gorm.DB

func init() {
	auth.InitConfig(CasdoorEndpoint, ClientId, ClientSecret, JwtPublicKey, CasdoorOrganization, CasdoorApplication)
}

func main() {
	app := fiber.New()
	db = connectDB()
	db.AutoMigrate(&User{})

	app.Use(cors.New())
	app.Use(logger.New())
	app.Use(recover.New())

	app.Get("/callback", SignIn)
	app.Get("/", func(c *fiber.Ctx) error { return c.SendString("ok") })

	app.Listen(":9000")
}
