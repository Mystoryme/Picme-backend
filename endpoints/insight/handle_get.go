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
	wg.Add(8)

	// Use goroutines for concurrent database queries
	go func() {
		defer wg.Done()
		views := 0
		mod.DB.Raw("SELECT COUNT(*) FROM insights WHERE triggee_id = ? AND insight_type = 'view' AND created_at >= CURDATE() - INTERVAL 1 WEEK", l.Id).Scan(&views)
		lastWeek.Views = int64(views)
	}()

	go func() {
		defer wg.Done()
		likes := 0
		mod.DB.Raw("SELECT COUNT(*) FROM insights WHERE triggee_id = ? AND insight_type = 'like'  AND created_at >= CURDATE() - INTERVAL 1 WEEK", l.Id).Scan(&likes)
		lastWeek.Likes = int64(likes)
	}()

	go func() {
		defer wg.Done()
		views := 0
		mod.DB.Raw("SELECT COUNT(*) FROM insights WHERE triggee_id = ? AND insight_type = 'view' AND created_at >= CURDATE() - INTERVAL 2 WEEK AND created_at <= CURDATE() - INTERVAL 1 WEEK", l.Id).Scan(&views)
		last2Week.Views = int64(views)
	}()

	go func() {
		defer wg.Done()
		likes := 0
		mod.DB.Raw("SELECT COUNT(*) FROM insights WHERE triggee_id = ? AND insight_type = 'like' AND created_at >= CURDATE() - INTERVAL 2 WEEK AND created_at <= CURDATE() - INTERVAL 1 WEEK", l.Id).Scan(&likes)
		last2Week.Likes = int64(likes)
	}()

	go func() {
		defer wg.Done()
		views := 0
		mod.DB.Raw("SELECT COUNT(*) FROM insights WHERE triggee_id = ? AND insight_type = 'view' AND created_at >= CURDATE() - INTERVAL 3 WEEK AND created_at <= CURDATE() - INTERVAL 2 WEEK", l.Id).Scan(&views)
		last3Week.Views = int64(views)
	}()

	go func() {
		defer wg.Done()
		likes := 0
		mod.DB.Raw("SELECT COUNT(*) FROM insights WHERE triggee_id = ? AND insight_type = 'like' AND created_at >= CURDATE() - INTERVAL 3 WEEK AND created_at <= CURDATE() - INTERVAL 2 WEEK", l.Id).Scan(&likes)
		last3Week.Likes = int64(likes)
	}()

	go func() {
		defer wg.Done()
		views := 0
		mod.DB.Raw("SELECT COUNT(*) FROM insights WHERE triggee_id = ? AND insight_type = 'view' AND created_at >= CURDATE() - INTERVAL 4 WEEK AND created_at < CURDATE() - INTERVAL 3 WEEK", l.Id).Scan(&views)
		last4Week.Views = int64(views)
	}()

	go func() {
		defer wg.Done()
		likes := 0
		mod.DB.Raw("SELECT COUNT(*) FROM insights WHERE triggee_id = ? AND insight_type = 'like' AND created_at >= CURDATE() - INTERVAL 4 WEEK AND created_at < CURDATE() - INTERVAL 3 WEEK", l.Id).Scan(&likes)
		last4Week.Likes = int64(likes)
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
