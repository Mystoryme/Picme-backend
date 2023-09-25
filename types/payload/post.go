package payload

import "picme-backend/types/enum"

type CreatePostBody struct {
	Caption  *string        `json:"caption"`
	Category *enum.Category `json:"category"`
}

type PostResponse struct {
	PostId        *uint64 `json:"postId"`
	OwnerId       *uint64 `json:"ownerId"`
	OwnerUsername *string `json:"ownerUsername"`
	Caption       *string `json:"caption"`
	ImageUrl      *string `json:"imageUrl"`
	LikeCount     *uint64 `json:"likeCount"`
	CommentCount  *uint64 `json:"commentCount"`
}
