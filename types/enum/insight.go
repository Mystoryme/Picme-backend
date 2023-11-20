package enum

import (
	"encoding/json"
	"fmt"
)

type Insight string

const (
	InsightView   Insight = "view"
	InsightLike   Insight = "like"
	InsightSearch Insight = "search"
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
	if Insight != InsightView && Insight != InsightLike && Insight != InsightSearch {
		return fmt.Errorf("invalid Insight enum value: %s", Insight)
	}

	*s = Insight

	return nil
}
