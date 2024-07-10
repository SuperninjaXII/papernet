package main
import(
  "github.com/gofiber/fiber/v2"
  "search/routes"
  "github.com/gofiber/template/html/v2"
)
import "search/config"
func init() {
  config.Database()
}

func main() {
  engine := html.New("./views", ".html")
  app := fiber.New(
    fiber.Config{
        Views: engine,
    },
    )
  app.Static("/", "./public")
  routes.Routes(app)
  app.Listen(":3000")
}