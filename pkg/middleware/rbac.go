package middleware

import "github.com/gofiber/fiber/v3"

func RbacRequired(roleRequired []string) fiber.Handler {
	return func(c fiber.Ctx) error {
		role, ok := c.Locals("role").(string)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON("Role not found")
		}

		for _, j := range roleRequired {
			if j == role {
				return c.Next()
			}
		}
		return c.Status(fiber.StatusForbidden).JSON("Permission denied")
	}
}
