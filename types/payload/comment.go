package payload

type CreateCommentBody struct {
	PostId  *uint64 `json:"postId"`
	Message *string `json:"message"`
}

type CommentRequest struct {
	PostId *uint64 `json:"postId"`
}

type DeleteCommentRequest struct {
	Id *uint64 `json:"id"`
}

type CommentRespond struct {
	Id        *uint64 `json:"Id"`
	UserId    *uint64 `json:"userId"`
	Username  *string `json:"username"`
	PostId    *uint64 `json:"postId"`
	AvatarUrl *string `json:"avatarUrl"`
	Message   *string `json:"message"`
}
