package model

import "picme-backend/types/table"

type PostWithCount struct {
	table.Post
	LikeCount    *int64
	CommentCount *int64
}
