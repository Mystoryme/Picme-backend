package enum

import (
	"encoding/json"
	"fmt"
)

type Category string

const (
	CategoryPainting Category = "painting"
	CategoryDrawing  Category = "drawing_&_illustration"
	CategoryMixMedia Category = "mixed_media_&_collage"
	CategoryGraphic  Category = "digital_/_graphic"
	CategoryOther    Category = "other"
)

func (s *Category) UnmarshalJSON(data []byte) error {
	var val string
	if err := json.Unmarshal(data, &val); err != nil {
		return err
	}

	category := Category(val)
	if category != CategoryPainting && category != CategoryDrawing && category != CategoryMixMedia && category != CategoryGraphic && category != CategoryOther {
		return fmt.Errorf("invalid category enum value: %s", category)
	}

	*s = category

	return nil
}
