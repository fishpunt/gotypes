package datetime

import (
	"encoding/json"
	"testing"
	"time"
)

func TestNullDateTimeMarshalJSON(t *testing.T) {
	// Test case 1: Valid NullDateTime
	dt := NullDateTime{
		DateTime: DateTime{
			Time:  time.Date(2022, time.March, 7, 0, 0, 0, 0, time.UTC),
			Valid: true,
		},
	}
	expectedJSON := `"` + dt.Time.Format(defaultDatetimeLayout) + `"`
	jsonData, err := json.Marshal(dt)
	if err != nil {
		t.Fatal(err)
	}
	if string(jsonData) != expectedJSON {
		t.Errorf("Expected JSON: %s, but got: %s", expectedJSON, string(jsonData))
	}

	// Test case 2: Invalid NullDateTime
	invalidDt := NullDateTime{
		DateTime: DateTime{
			Valid: false,
		},
	}
	expectedNullJSON := "null"
	invalidJSONData, err := json.Marshal(invalidDt)
	if err != nil {
		t.Fatal(err)
	}
	if string(invalidJSONData) != expectedNullJSON {
		t.Errorf("Expected JSON: %s, but got: %s", expectedNullJSON, string(invalidJSONData))
	}
}
