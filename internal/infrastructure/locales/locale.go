package locales

import (
	"fmt"

	"github.com/gofiber/contrib/fiberi18n/v2"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/text/language"
)

func New() fiber.Handler {
	return fiberi18n.New(&fiberi18n.Config{
		RootPath:        "internal/infrastructure/locales/l10n",
		AcceptLanguages: []language.Tag{language.EuropeanPortuguese, language.English},
		DefaultLanguage: language.EuropeanPortuguese,
	})
}

func Localize(ctx *fiber.Ctx, key string, fallback string) string {
	return fiberi18n.MustLocalize(ctx, key)
}

func LocalizeHeader(ctx *fiber.Ctx, key string, fallback string) string {
	return fiberi18n.MustLocalize(ctx, fmt.Sprintf("%s-header", key))
}
