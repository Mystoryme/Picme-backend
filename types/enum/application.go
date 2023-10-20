package enum

import (
	"encoding/json"
	"fmt"
)

type Application string

const (
	ApplicationProcreate       Application = "procreate"
	ApplicationIbisPaintX      Application = "ibis_paintX"
	ApplicationClipStudioPaint Application = "clip_studio_paint"
	ApplicationBlender         Application = "blender"
	ApplicationPhotoshop       Application = "photoshop"
	ApplicationOther           Application = "other "
)

func (s *Application) UnmarshalJSON(data []byte) error {
	var val string
	if err := json.Unmarshal(data, &val); err != nil {
		return err
	}
	if err := json.Unmarshal(data, &val); err != nil {
		return err
	}

	Application := Application(val)
	if Application != ApplicationProcreate && Application != ApplicationIbisPaintX && Application != ApplicationClipStudioPaint && Application != ApplicationBlender && Application != ApplicationPhotoshop && Application != ApplicationOther {
		return fmt.Errorf("invalid Application enum value: %s", Application)
	}

	*s = Application

	return nil
}
