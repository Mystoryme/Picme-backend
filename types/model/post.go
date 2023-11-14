package model

import "picme-backend/types/table"

type PostWithCount struct {
	table.Post
	LikeCount    *int64
	CommentCount *int64
	Liked        *int64
	Booked       *int64
	IsLiked      bool `json:"isLiked"`
	IsBooked     bool `json:"isBooked"`
}

type BookmarkPostWithCount struct {
	table.PostBookMark
	ImageUrl     *string
	LikeCount    *int64
	CommentCount *int64
}

type BoostPostwithImage struct {
	table.PostBoost
	ImageUrl *string
}
