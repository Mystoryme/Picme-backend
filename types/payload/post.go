package payload

import "picme-backend/types/enum"

type CreatePostBody struct {
	ImageUrl    *string           `json:"imageUrl" validate:"required"`
	Caption     *string           `json:"caption" validate:"required"`
	Category    *enum.Category    `json:"category" validate:"required"`
	Application *enum.Application `json:"application" validate:"required"`
}

type PostResponse struct {
	PostId        *uint64 `json:"postId"`
	OwnerId       *uint64 `json:"ownerId"`
	OwnerUsername *string `json:"ownerUsername"`
	Caption       *string `json:"caption"`
	ImageUrl      *string `json:"imageUrl"`
	LikeCount     *int64  `json:"likeCount"`
	CommentCount  *int64  `json:"commentCount"`
}

type PostQuery struct {
	Category *enum.Category `query:"category" validate:"omitempty,oneof=painting drawing mixedmedia digital other"`
}

type CreateLikeBody struct {
	PostId *uint64 `json:"postId"`
}
