package payload

import (
	"picme-backend/types/enum"
)

type CreatePostBody struct {
	Caption     *string           `form:"caption" validate:"required"`
	Category    *enum.Category    `form:"category" validate:"required"`
	Application *enum.Application `form:"application" validate:"required"`
}

type PostResponse struct {
	PostId        *uint64           `json:"postId"`
	OwnerId       *uint64           `json:"ownerId"`
	OwnerUsername *string           `json:"ownerUsername"`
	Caption       *string           `json:"caption"`
	ImageUrl      *string           `json:"imageUrl"`
	Application   *enum.Application `json:"application"`
	LikeCount     *int64            `json:"likeCount"`
	CommentCount  *int64            `json:"commentCount"`
	IsLiked       *bool             `json:"isLiked"`
}

type PostQuery struct {
	Category *enum.Category `query:"category" validate:"omitempty,oneof=painting drawing mixedmedia digital other"`
}

type CreateLikeBody struct {
	PostId *uint64 `json:"postId"`
}

type UnLikeBody struct {
	PostId *uint64 `json:"postId"`
}

type CreateBookmarkBody struct {
	PostId *uint64 `json:"postId"`
}

type UnBookmarkBody struct {
	PostId *uint64 `json:"postId"`
}

type ViewBody struct {
	PostId *uint64 `json:"postId"`
}

type BoostBody struct {
	PostId   *uint64 `json:"postId"`
	BoostDay *int    `json:"boostDay"`
}

type CreateDonateBody struct {
	PostId       *uint64 `json:"postId"`
	DonateAmount *uint64 `json:"donateAmount"`
}

type DeletePostRequest struct {
	PostId *uint64 `json:"postId"`
}
