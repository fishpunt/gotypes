package date

import (
	"database/sql/driver"
	"encoding/xml"
	"fmt"
	"strings"
	"time"
)

const (
	dateLayout       = "2006-01-02"
	dateEncodeLayout = "2006-01-02"
	dateDefaultValue = "1970-01-01"
)

/**
 * Date type
 */

type Date struct {
	Time  time.Time
	Valid bool
}

// New
func New(src *time.Time) Date {
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
	t := time.Time(dt.Time)
	return t.Format(dateLayout)
}

// UnmarshalXML
func (dt *Date) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string
	d.DecodeElement(&v, &start)
	result, err := time.Parse(dateLayout, v)
	if err != nil {
		return fmt.Errorf("date UnmarshalXML error: %s", err)
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
	nt, err := time.Parse(dateLayout, s)
	if err != nil {
		return fmt.Errorf("date UnmarshalJSON error: %s", err)
	}

	dt.Time = nt
	dt.Valid = true

	return nil
}

// MarshalJSON
func (dt Date) MarshalJSON() ([]byte, error) {
	if !dt.Valid {
		return []byte(""), nil
	}
	return []byte(dt.String()), nil
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
	return dt.Time, nil
}

func (dt Date) EqualTo(src time.Time) bool {
	if !dt.Valid {
		return false
	}

	y1, m1, d1 := dt.Time.Date()
	y2, m2, d2 := src.Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}
