package date

import (
	"database/sql/driver"
	"encoding/xml"
	"fmt"
	"strings"
	"time"
)

/**
 * Date type
 */

type Date struct {
	Time  time.Time
	Valid bool
}

// New
func NewDate(src *time.Time) Date {
	if src == nil {
		now := time.Now()
		src = &now
	}

	dt := Date{
		Time:  *src,
		Valid: true,
	}

	return dt
}

// String
// String returns the time in the custom format
func (dt Date) String() string {
	if !dt.Valid {
		return ""
	}

	return dt.Time.Format(*dateLayoutOutput)
}

// UnmarshalXML
func (dt *Date) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string
	d.DecodeElement(&v, &start)
	if v == "" {
		dt.Valid = false
		return nil
	}

	// result, err := time.Parse(*dateLayoutInput, v)
	result, err := dt.parseTime(v)
	if err != nil {
		return err
	}
	dt.Time = result
	dt.Valid = true

	return nil
}

// MarshalXML
func (dt Date) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	v := dt.String()
	return e.EncodeElement(v, start)
}

// UnmarshalJSON
func (dt *Date) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)
	if s == "" || s == "null" {
		dt.Valid = false
		return nil
	}

	// nt, err := time.Parse(*dateLayoutInput, s)
	nt, err := dt.parseTime(s)
	if err != nil {
		return err
	}

	dt.Time = nt
	dt.Valid = true

	return nil
}

// MarshalJSON
func (dt Date) MarshalJSON() ([]byte, error) {
	if !dt.Valid {
		return []byte("\"\""), nil
	}
	return []byte(fmt.Sprintf("%q", dt.String())), nil
}

// Scan
// Scan implements the Scanner interface.
func (dt *Date) Scan(value interface{}) error {
	dt.Time, dt.Valid = value.(time.Time)
	return nil
}

// Value
// Value implements the driver Valuer interface.
func (dt Date) Value() (driver.Value, error) {
	if !dt.Valid {
		return "", nil
	}
	return dt.String(), nil
}

func (dt Date) EqualTo(src time.Time) bool {
	if !dt.Valid {
		return false
	}

	y1, m1, d1 := dt.Time.Date()
	y2, m2, d2 := src.Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}

// ParseTime from different layouts
func (dt Date) parseTime(src string) (time.Time, error) {
	var result time.Time

	var err error
	var firstError error
	for _, v := range dateLayoutInputs {
		result, err = time.Parse(v, src)
		if err == nil {
			return result, nil
		}
		if v == *dateLayoutOutput {
			firstError = err
		}
	}

	if err != nil && firstError == nil {
		return result, err
	}

	return result, firstError
}
