package handlers

import (
	"html/template"

	"github.com/gofiber/fiber/v2"
	"github.com/mrrizkin/finteligo/third_party/vite"
)

func (h *Handlers) AdminUI(c *fiber.Ctx) error {
	viteHTMLTag := vite.ReactRefresh() + vite.Entry("resources/assets/scripts/main.tsx")

	return c.Render("views/index", fiber.Map{
		"viteHTMLTag": template.HTML(viteHTMLTag),
	})
}
