package datetime

import (
	"database/sql/driver"
	"encoding/xml"
	"fmt"
	"strings"
	"time"
)

const (
	dateTimeLayout       = time.RFC3339
	dateTimeLayoutAlt    = "2006-01-02T15:04:05"
	dateTimeLayoutAlt2   = "2006-01-02T15:04Z07:00"
	dateTimeEncodeLayout = time.RFC3339
	DateTimeDefaultValue = "1970-01-01T00:00:00"
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

	return dt.Time.Format(dateTimeLayout)
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
		return fmt.Errorf("datetime UnmarshalXML error: %s", err)
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
		return fmt.Errorf("datetime MarshalXML error: %s", err)
	}

	dt.Time = nt
	dt.Valid = true

	return nil
}

// MarshalJSON
func (dt DateTime) MarshalJSON() ([]byte, error) {
	if !dt.Valid {
		return []byte(""), nil
	}
	return []byte(fmt.Sprintf("%q", dt.Time.Format(dateTimeLayoutAlt))), nil
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

	result, err := time.Parse(dateTimeLayout, src)
	if err != nil {
		// With Timezone
		err = nil
		result, err = time.Parse(dateTimeLayoutAlt, src)
		if err != nil {
			err = nil
			result, err = time.Parse(dateTimeLayoutAlt2, src)
			if err != nil {
				return result, fmt.Errorf("datetime failed to parse time. (error: %s)", err)
			}
		}
	}

	return result, nil
}

func ParseTime(src string) (time.Time, error) {

	result, err := time.Parse(dateTimeLayout, src)
	if err != nil {
		// With Timezone
		err = nil
		result, err = time.Parse(dateTimeLayoutAlt, src)
		if err != nil {
			err = nil
			result, err = time.Parse(dateTimeLayoutAlt2, src)
			if err != nil {
				return result, fmt.Errorf("datetime failed to parse time. (error: %s)", err)
			}
		}
	}

	return result, nil
}

func (dt DateTime) DateEqual(src time.Time) bool {
	if !dt.Valid {
		return false
	}

	y1, m1, d1 := dt.Time.Date()
	y2, m2, d2 := src.Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}
