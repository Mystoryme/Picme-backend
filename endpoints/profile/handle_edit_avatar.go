package profileEndpoint

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/minio/minio-go/v7"
	mod "picme-backend/modules"
	"picme-backend/types/misc"
	"picme-backend/types/payload"
	"picme-backend/types/response"
	"picme-backend/types/table"
	"picme-backend/utils/text"
)

func AvatarHandler(c *fiber.Ctx) error {

	// * Parse user claims
	l := c.Locals("l").(*jwt.Token).Claims.(*misc.UserClaim)

	// * Parse body
	body := new(payload.EditAvatarBody)
	if err := c.BodyParser(body); err != nil {
		return response.Error(false, "Unable to parse body", err)
	}

	// * Validate body
	if err := text.Validator.Struct(body); err != nil {
		return err
	}

	// * Parse file form
	// Note: file is a *multipart.FileHeader instance
	fileHeader, err := c.FormFile("file")
	if err != nil {
		return response.Error(false, "File not found", err)
	}

	// * Convert multipart.FileHeader to File (which satisfies io.Reader)
	// Note: Since file is a *multipart.FileHeader instance
	// and minio.PutObject() requires a io.Reader instance
	file, err := fileHeader.Open()
	if err != nil {
		return response.Error(false, "Failed to open file", err)
	}

	// * Generate filename

	filename := text.GenerateRandomString(12) + "_" + fileHeader.Filename
	imageUrl := "https://" + mod.Conf.BucketEndpoint + "/" + mod.Conf.BucketName + "/" + filename

	// * Upload file to minio
	info, err := mod.Minio.PutObject(
		c.Context(),
		"cs-picme",
		filename,
		file,
		fileHeader.Size,
		minio.PutObjectOptions{ContentType: fileHeader.Header.Get("Content-Type")},
	)
	if err != nil {
		return response.Error(false, "Failed to upload file", err)
	}

	// * Dump uploaded info
	spew.Dump(info)

	//query user by userid
	var user table.User
	if result := mod.DB.Where("id = ?", l.Id).First(&user); result.Error != nil {
		return response.Error(false, "Unable to fetch user profile", result.Error)
	}

	user.AvatarUrl = &imageUrl

	if result := mod.DB.Save(&user); result.Error != nil {
		return response.Error(false, "Unable to update user avatar", result.Error)
	}

	return c.JSON(response.Info("Successfully edit avatar!"))
}
