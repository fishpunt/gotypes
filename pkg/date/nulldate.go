package date

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type NullDate struct {
	Date
}

func NewNullDate(src *time.Time) NullDate {
	if src == nil {
		now := time.Now()
		src = &now
	}

	dt := NullDate{}
	dt.Time = *src
	dt.Valid = true

	return dt
}

// MarshalJSON
func (dt NullDate) MarshalJSON() ([]byte, error) {
	if !dt.Valid {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("%q", dt.String())), nil
}

// Value
// Value implements the driver Valuer interface.
func (dt NullDate) Value() (driver.Value, error) {
	if !dt.Valid {
		return nil, nil
	}
	return dt.Time, nil
}
