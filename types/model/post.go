package model

import "picme-backend/types/table"

type PostWithCount struct {
	table.Post
	LikeCount    *int64
	CommentCount *int64
}

type BookmarkPostWithCount struct {
	table.PostBookMark
	ImageUrl     *string
	LikeCount    *int64
	CommentCount *int64
}
