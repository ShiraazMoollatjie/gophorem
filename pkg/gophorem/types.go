package gophorem

import (
	"fmt"
	"time"
)

// EmptyTime is a type where the time can be an empty string.
type EmptyTime struct {
	time.Time
}

// UnmarshalJSON is a copy of the time.Time UnmarshalJSON function
// with an extra check for an empty string.
func (m *EmptyTime) UnmarshalJSON(data []byte) error {
	if string(data) == "null" || string(data) == `""` {
		return nil
	}

	tt, err := time.Parse(fmt.Sprintf(`"%s"`, time.RFC3339), string(data))
	if err != nil {
		return err
	}

	m.Time = tt
	return nil
}
