package date

import (
	"encoding/xml"
	"testing"
	"time"
)

func TestDateUnmarshalXML(t *testing.T) {
	// Test case 1: Valid Date
	dateStr := "2022-03-07"
	xmlStr := "<Date>" + dateStr + "</Date>"

	dt := &Date{}
	err := xml.Unmarshal([]byte(xmlStr), dt)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %s", err)
	}

	expectedDate, _ := time.Parse("2006-01-02", dateStr)
	if !dt.Valid || !dt.Time.Equal(expectedDate) {
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

	// Test case 4: Valid DateTime
	date2Str := "2023-01-04T15:36:22"
	xml2Str := "<Date>" + date2Str + "</Date>"

	dt2 := &Date{}
	err = xml.Unmarshal([]byte(xml2Str), dt2)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %s", err)
	}

	expected2Date, _ := time.Parse("2006-01-02T15:04:05", date2Str)
	if !dt2.Valid || dt2.Time.Format("2006-01-02") != expected2Date.Format("2006-01-02") {
		t.Errorf("Unexpected result. Expected: %s, Got: %s", expected2Date.Format("2006-01-02"), dt2.Time.Format("2006-01-02"))
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
		t.Fatalf("Expected Date to be invalid, but it is valid")
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
		t.Fatalf("Expected Date to be invalid, but it is valid")
	}

	// Test case 4: Valid DateTime
	date2Str := "2023-01-04T15:36:22"
	json2Str := "\"" + date2Str + "\""

	dt2 := &Date{}
	err = dt2.UnmarshalJSON([]byte(json2Str))
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON: %s", err)
	}

	expected2Date, _ := time.Parse("2006-01-02T15:04:05", date2Str)
	if !dt2.Valid || dt2.Time.Format("2006-01-02") != expected2Date.Format("2006-01-02") {
		t.Errorf("Unexpected result. Expected: %s, Got: %s", expected2Date.Format("2006-01-02"), dt2.Time.Format("2006-01-02"))
	}

	// Test case 5: Null DateTime
	nullJsonData := "null"
	nullDt := &Date{}
	err = nullDt.UnmarshalJSON([]byte(nullJsonData))
	expectedError = "parsing time \"null\" as \"2006-01-02\": cannot parse \"null\" as \"2006\""
	if err == nil || err.Error() != expectedError {
		t.Errorf("Expected error: %s, but got: %v", expectedError, err)
	}
	if nullDt.Valid {
		t.Fatalf("Expected Date to be invalid, but it is valid")
	}
}
