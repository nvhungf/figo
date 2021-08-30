package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nvhungf/figo/pkg/handler"
)

// Expenses route
func Expenses(app *fiber.App) {
	var h handler.ExpenseHandler
	r := app.Group("/expenses")
	r.Get("/", h.Index)
}
