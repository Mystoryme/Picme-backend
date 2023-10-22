package payload

import "picme-backend/types/enum"

type ProfileInfo struct {
	Id        *uint64 `json:"id"`
	Username  *string `json:"username"`
	Bio       *string `json:"bio"`
	Contact   *string `json:"contact"`
	AvatarUrl *string `json:"avatarUrl"`
}

type ProfilePostResponse struct {
	PostId        *uint64 `json:"postId"`
	OwnerId       *uint64 `json:"ownerId"`
	OwnerUsername *string `json:"ownerUsername"`
	Caption       *string `json:"caption"`
	ImageUrl      *string `json:"imageUrl"`
	LikeCount     *int64  `json:"likeCount"`
	CommentCount  *int64  `json:"commentCount"`
}

type ProfileQuery struct {
	SortBy *enum.SortBy `query:"sortBy" validate:"omitempty,oneof=date like"`
}

type DonateBody struct {
	UserId       *uint64 `json:"userId"`
	DonateAmount *uint64 `json:"donateAmount"`
}
