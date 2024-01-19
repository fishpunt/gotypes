package date

import (
	"encoding/xml"
	"testing"
	"time"
)

func TestDateUnmarshalXML(t *testing.T) {
	// Test case 1: Valid DateTime
	dateStr := "2022-03-07"
	xmlStr := "<Date>" + dateStr + "</Date>"

	dt := &Date{}
	err := xml.Unmarshal([]byte(xmlStr), dt)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %s", err)
	}

	expectedDate, _ := time.Parse("2006-01-02", dateStr)
	if !dt.Valid || dt.Time != expectedDate {
		t.Errorf("Unexpected result. Expected: %s, Got: %s", expectedDate, dt.Time)
	}

	// Test case 2: Empty Date
	emptyXMLStr := "<Date></Date>"
	emptyDt := &Date{}
	err = xml.Unmarshal([]byte(emptyXMLStr), emptyDt)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %s", err)
	}
	if emptyDt.Valid {
		t.Errorf("Expected Date to be invalid, but it is valid")
	}

	// Test case 3: Invalid Date
	invalidXMLStr := "<Date>invalid</Date>"
	invalidDt := &Date{}
	err = xml.Unmarshal([]byte(invalidXMLStr), invalidDt)
	expectedError := "parsing time \"invalid\" as \"2006-01-02\": cannot parse \"invalid\" as \"2006\""
	if err == nil || err.Error() != expectedError {
		t.Errorf("Expected error: %s, but got: %v", expectedError, err)
	}
	if invalidDt.Valid {
		t.Errorf("Expected Date to be invalid, but it is valid")
	}
}

func TestDateUnmarshalJson(t *testing.T) {
	// Test case 1: Valid DateTime
	dateStr := "2022-03-07"
	jsonStr := "\"" + dateStr + "\""

	dt := &Date{}
	err := dt.UnmarshalJSON([]byte(jsonStr))
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON: %s", err)
	}

	expectedDate, _ := time.Parse("2006-01-02", dateStr)
	if !dt.Valid || dt.Time != expectedDate {
		t.Errorf("Unexpected result. Expected: %s, Got: %s", expectedDate, dt.Time)
	}

	// Test case 2: Empty Date
	emptyJSONStr := "\"\""
	emptyDt := &Date{}
	err = emptyDt.UnmarshalJSON([]byte(emptyJSONStr))
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON: %s", err)
	}
	if emptyDt.Valid {
		t.Errorf("Expected Date to be invalid, but it is valid")
	}

	// Test case 3: Invalid Date
	invalidJSONStr := "\"invalid\""
	invalidDt := &Date{}
	err = invalidDt.UnmarshalJSON([]byte(invalidJSONStr))
	expectedError := "parsing time \"invalid\" as \"2006-01-02\": cannot parse \"invalid\" as \"2006\""
	if err == nil || err.Error() != expectedError {
		t.Errorf("Expected error: %s, but got: %v", expectedError, err)
	}
	if invalidDt.Valid {
		t.Errorf("Expected Date to be invalid, but it is valid")
	}
}
