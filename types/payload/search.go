package payload

type SearchQuery struct {
	Username *string `query:"username"`
}

type ViewProfileBody struct {
	UserId *uint64 `json:"userId"`
}
