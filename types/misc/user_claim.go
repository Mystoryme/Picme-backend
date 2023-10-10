package misc

type UserClaim struct {
	Id *uint64 `json:"id"`
}

func (T *UserClaim) Valid() error {
	return nil
}
