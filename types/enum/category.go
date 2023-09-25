package enum

import (
	"encoding/json"
	"fmt"
)

type Category string

const (
	CategoryPainting Category = "painting"
	CategoryDrawing  Category = "drawing"
)

func (s *Category) UnmarshalJSON(data []byte) error {
	var val string
	if err := json.Unmarshal(data, &val); err != nil {
		return err
	}

	category := Category(val)
	if category != CategoryPainting && category != CategoryDrawing {
		return fmt.Errorf("invalid category enum value: %s", category)
	}

	*s = category

	return nil
}
