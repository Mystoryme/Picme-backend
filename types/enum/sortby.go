package enum

import (
	"encoding/json"
	"fmt"
)

type SortBy string

const (
	SortByDate SortBy = "date"
	SortByLike SortBy = "like"
)

func (s *SortBy) UnmarshalJSON(data []byte) error {
	var val string
	if err := json.Unmarshal(data, &val); err != nil {
		return err
	}

	SortBy := SortBy(val)
	if SortBy != SortByDate && SortBy != SortByLike {
		return fmt.Errorf("invalid SortBy enum value: %s", SortBy)
	}

	*s = SortBy

	return nil
}
