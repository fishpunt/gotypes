package datetime

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type NullDateTime struct {
	DateTime
}

func NewNullDatetime(src *time.Time) NullDateTime {
	if src == nil {
		now := time.Now()
		src = &now
	}

	dt := NullDateTime{}
	dt.Time = *src
	dt.Valid = true

	return dt
}

// MarshalJSON
func (dt NullDateTime) MarshalJSON() ([]byte, error) {
	if !dt.Valid {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("%q", dt.Time.Format(dateTimeLayoutAlt))), nil
}

// Value
// Value implements the driver Valuer interface.
func (dt NullDateTime) Value() (driver.Value, error) {
	if !dt.Valid {
		return nil, nil
	}
	return dt.Time, nil
}
