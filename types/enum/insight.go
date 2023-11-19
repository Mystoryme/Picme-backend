package enum

import (
	"encoding/json"
	"fmt"
)

type Insight string

const (
	InsightView Insight = "view"
	InsightLike Insight = "like"
)

func (s *Insight) UnmarshalJSON(data []byte) error {
	var val string
	if err := json.Unmarshal(data, &val); err != nil {
		return err
	}
	if err := json.Unmarshal(data, &val); err != nil {
		return err
	}

	Insight := Insight(val)
	if Insight != InsightView && Insight != InsightLike {
		return fmt.Errorf("invalid Insight enum value: %s", Insight)
	}

	*s = Insight

	return nil
}
