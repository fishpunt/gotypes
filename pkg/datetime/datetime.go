package datetime

import (
	"database/sql/driver"
	"encoding/xml"
	"fmt"
	"strings"
	"time"
)

/**
 * DateTime type
 */

type DateTime struct {
	Time  time.Time
	Valid bool
}

// New
func NewDateTime(src *time.Time) DateTime {
	if src == nil {
		now := time.Now()
		src = &now
	}

	dt := DateTime{
		Time:  *src,
		Valid: true,
	}

	return dt
}

// String
// String returns the time in the custom format
func (dt DateTime) String() string {
	if !dt.Valid {
		return ""
	}

	return dt.Time.Format(*datetimeLayoutOutput)
}

// UnmarshalXML
func (dt *DateTime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string
	d.DecodeElement(&v, &start)
	if v == "" {
		dt.Valid = false
		return nil
	}

	result, err := dt.parseTime(v)
	if err != nil {
		return err
	}
	if result.IsZero() {
		dt.Valid = false
		return nil
	}

	dt.Time = result
	dt.Valid = true

	return nil
}

// MarshalXML
func (dt DateTime) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	v := dt.String()
	return e.EncodeElement(v, start)
}

// UnmarshalJSON
func (dt *DateTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)

	if s == "" {
		dt.Valid = false
		return nil
	}

	nt, err := dt.parseTime(s)
	if err != nil {
		return err
	}

	if nt.IsZero() {
		dt.Valid = false
		return nil
	}

	dt.Time = nt
	dt.Valid = true

	return nil
}

// MarshalJSON
func (dt DateTime) MarshalJSON() ([]byte, error) {
	if !dt.Valid {
		return []byte("\"\""), nil
	}
	return []byte(fmt.Sprintf("%q", dt.String())), nil
}

// Scan
// Scan implements the Scanner interface.
func (dt *DateTime) Scan(value interface{}) error {
	dt.Time, dt.Valid = value.(time.Time)
	return nil
}

// Value
// Value implements the driver Valuer interface.
func (dt DateTime) Value() (driver.Value, error) {
	if !dt.Valid {
		return nil, nil
	}
	return dt.Time, nil
}

// ParseTime from different layouts
func (dt DateTime) parseTime(src string) (time.Time, error) {
	var result time.Time

	var err error
	var firstError error
	for _, v := range datetimeLayoutInputs {
		result, err = time.Parse(v, src)
		if err == nil {
			return result, nil
		}
		if v == *datetimeLayoutOutput {
			firstError = err
		}
	}

	if err != nil && firstError == nil {
		return result, err
	}

	return result, firstError
}

func (dt DateTime) DateEqual(src time.Time) bool {
	if !dt.Valid {
		return false
	}

	y1, m1, d1 := dt.Time.Date()
	y2, m2, d2 := src.Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}
