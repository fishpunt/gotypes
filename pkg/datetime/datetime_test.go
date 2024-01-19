package datetime

import (
	"encoding/xml"
	"testing"
	"time"
)

func TestDateTimeUnmarshalXML(t *testing.T) {
	// Test case 1: Valid DateTime
	xmlData := `<DateTime>2022-03-07T00:00:00Z</DateTime>`
	dt := DateTime{}
	err := xml.Unmarshal([]byte(xmlData), &dt)
	if err != nil {
		t.Fatal(err)
	}
	expectedTime := time.Date(2022, time.March, 7, 0, 0, 0, 0, time.UTC)
	if !dt.Valid || dt.Time != expectedTime {
		t.Errorf("Expected DateTime: %v, but got: %v", expectedTime, dt.Time)
	}

	// Test case 2: Empty DateTime
	emptyXMLData := `<DateTime></DateTime>`
	emptyDt := DateTime{}
	err = xml.Unmarshal([]byte(emptyXMLData), &emptyDt)
	if err != nil {
		t.Fatal(err)
	}
	if emptyDt.Valid {
		t.Errorf("Expected DateTime to be invalid, but it is valid")
	}

	// Test case 3: Invalid DateTime
	invalidXMLData := `<DateTime>invalid</DateTime>`
	invalidDt := DateTime{}
	err = xml.Unmarshal([]byte(invalidXMLData), &invalidDt)
	expectedError := "parsing time \"invalid\" as \"2006-01-02T15:04:05Z07:00\": cannot parse \"invalid\" as \"2006\""
	if err == nil || err.Error() != expectedError {
		t.Errorf("Expected error: %s, but got: %v", expectedError, err)
	}
	if invalidDt.Valid {
		t.Errorf("Expected DateTime to be invalid, but it is valid")
	}
}

func TestDateTimeUnmarshalJson(t *testing.T) {
	// Test case 1: Valid DateTime
	jsonData := `"2022-03-07T00:00:00Z"`
	dt := DateTime{}
	err := dt.UnmarshalJSON([]byte(jsonData))
	if err != nil {
		t.Fatal(err)
	}
	expectedTime := time.Date(2022, time.March, 7, 0, 0, 0, 0, time.UTC)
	if !dt.Valid || dt.Time != expectedTime {
		t.Errorf("Expected DateTime: %v, but got: %v", expectedTime, dt.Time)
	}

	// Test case 2: Empty DateTime
	emptyJsonData := `""`
	emptyDt := DateTime{}
	err = emptyDt.UnmarshalJSON([]byte(emptyJsonData))
	if err != nil {
		t.Fatal(err)
	}
	if emptyDt.Valid {
		t.Errorf("Expected DateTime to be invalid, but it is valid")
	}

	// Test case 3: Invalid DateTime
	invalidJsonData := `"invalid"`
	invalidDt := DateTime{}
	err = invalidDt.UnmarshalJSON([]byte(invalidJsonData))
	expectedError := "parsing time \"invalid\" as \"2006-01-02T15:04:05Z07:00\": cannot parse \"invalid\" as \"2006\""
	if err == nil || err.Error() != expectedError {
		t.Errorf("Expected error: %s, but got: %v", expectedError, err)
	}
	if invalidDt.Valid {
		t.Errorf("Expected DateTime to be invalid, but it is valid")
	}
}
