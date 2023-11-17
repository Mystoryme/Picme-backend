package enum

import (
	"encoding/json"
	"fmt"
)

type Notification string

const (
	NotificationComment    Notification = "comment"
	NotificationLike       Notification = "like"
	NotificationUserDonate Notification = "user_donate"
	NotificationPostDonate Notification = "post_donate"
)

func (s *Notification) UnmarshalJSON(data []byte) error {
	var val string
	if err := json.Unmarshal(data, &val); err != nil {
		return err
	}
	if err := json.Unmarshal(data, &val); err != nil {
		return err
	}

	Notification := Notification(val)
	if Notification != NotificationComment && Notification != NotificationLike && Notification != NotificationUserDonate && Notification != NotificationPostDonate {
		return fmt.Errorf("invalid Application enum value: %s", Notification)
	}

	*s = Notification

	return nil
}
