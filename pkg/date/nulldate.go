package date

import (
	"database/sql/driver"
	"fmt"
)

type NullDate struct {
	Date
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
