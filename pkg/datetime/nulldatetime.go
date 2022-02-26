package datetime

import "database/sql/driver"

type NullDateTime struct {
	DateTime
}

// MarshalJSON
func (dt NullDateTime) MarshalJSON() ([]byte, error) {
	if !dt.Valid {
		return []byte("null"), nil
	}
	return []byte(dt.String()), nil
}

// Value
// Value implements the driver Valuer interface.
func (dt NullDateTime) Value() (driver.Value, error) {
	if !dt.Valid {
		return nil, nil
	}
	return dt.Time, nil
}
