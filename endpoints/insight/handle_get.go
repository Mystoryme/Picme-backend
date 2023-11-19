package postEndpoint

import (
	// ... (import statements)

	mod "picme-backend/modules"
	"picme-backend/types/misc"
	"picme-backend/types/payload"
	"picme-backend/types/response"
	"picme-backend/utils/text"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func GetHandler(c *fiber.Ctx) error {
	// Parse user claims
	l := c.Locals("l").(*jwt.Token).Claims.(*misc.UserClaim)

	// Parse query
	query := new(payload.PostQuery)
	if err := c.QueryParser(query); err != nil {
		return response.Error(false, "Unable to parse body", err)
	}

	// Validate query
	if err := text.Validator.Struct(query); err != nil {
		return err
	}

	var (
		lastWeek  payload.InsightObject
		last2Week payload.InsightObject
		last3Week payload.InsightObject
		last4Week payload.InsightObject
	)

	var wg sync.WaitGroup
	wg.Add(4)

	// Use goroutines for concurrent database queries
	go func() {
		defer wg.Done()
		mod.DB.Raw("SELECT views, likes FROM posts WHERE triggee_id = ? AND created_at >= CURDATE() - INTERVAL 1 WEEK AND created_at < CURDATE()", l.Id).Scan(&lastWeek)
	}()

	go func() {
		defer wg.Done()
		mod.DB.Raw("SELECT views, likes FROM posts WHERE triggee_id = ? AND created_at >= CURDATE() - INTERVAL 2 WEEK AND created_at < CURDATE() - INTERVAL 1 WEEK", l.Id).Scan(&last2Week)
	}()

	go func() {
		defer wg.Done()
		mod.DB.Raw("SELECT views, likes FROM posts WHERE triggee_id = ? AND created_at >= CURDATE() - INTERVAL 3 WEEK AND created_at < CURDATE() - INTERVAL 2 WEEK", l.Id).Scan(&last3Week)
	}()

	go func() {
		defer wg.Done()
		mod.DB.Raw("SELECT views, likes FROM posts WHERE triggee_id = ? AND created_at >= CURDATE() - INTERVAL 4 WEEK AND created_at < CURDATE() - INTERVAL 3 WEEK", l.Id).Scan(&last4Week)
	}()

	// Wait for all goroutines to finish
	wg.Wait()

	return c.JSON(response.Info(map[string]interface{}{
		"insight": []payload.InsightObject{
			lastWeek,
			last2Week,
			last3Week,
			last4Week,
		},
	}))
}
