package datetime

import (
	"encoding/json"
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

	// Test case 4: Valid DateTime
	xmlData = `<DateTime>2015-02-02T11:04:57.497</DateTime>`
	dt = DateTime{}
	err = xml.Unmarshal([]byte(xmlData), &dt)
	if err != nil {
		t.Fatal(err)
	}
	expectedTime = time.Date(2015, time.February, 2, 11, 04, 57, 497000000, time.UTC)
	if !dt.Valid || dt.Time != expectedTime {
		t.Errorf("Expected DateTime: %v, but got: %v", expectedTime, dt.Time)
	}
}

func TestDateTimeUnmarshalJson(t *testing.T) {
	type testJsonModel struct {
		CreationDate DateTime `json:"CreationDate"`
	}

	// Test case 1: Valid DateTime
	jsonData := `{"CreationDate": "2022-03-07T00:00:00Z"}`
	dt := testJsonModel{}
	err := json.Unmarshal([]byte(jsonData), &dt)
	if err != nil {
		t.Fatal(err)
	}
	expectedTime := time.Date(2022, time.March, 7, 0, 0, 0, 0, time.UTC)
	if !dt.CreationDate.Valid || dt.CreationDate.Time != expectedTime {
		t.Errorf("Expected DateTime: %v, but got: %v", expectedTime, dt.CreationDate.Time)
	}

	// Test case 2: Empty DateTime
	emptyJsonData := `{"CreationDate": ""}`
	emptyDt := testJsonModel{}
	err = json.Unmarshal([]byte(emptyJsonData), &emptyDt)
	if err != nil {
		t.Fatal(err)
	}
	if emptyDt.CreationDate.Valid {
		t.Errorf("Expected DateTime to be invalid, but it is valid")
	}

	// Test case 3: Invalid DateTime
	invalidJsonData := `"invalid"`
	invalidDt := testJsonModel{}
	err = json.Unmarshal([]byte(invalidJsonData), &invalidDt)
	// expectedError := "parsing time \"invalid\" as \"2006-01-02T15:04:05Z07:00\": cannot parse \"invalid\" as \"2006\""
	expectedError := "json: cannot unmarshal string into Go value of type datetime.testJsonModel"
	if err == nil || err.Error() != expectedError {
		t.Errorf("Expected error: %s, but got: %v", expectedError, err)
	}
	if invalidDt.CreationDate.Valid {
		t.Errorf("Expected DateTime to be invalid, but it is valid")
	}

	// Test case 4: Valid DateTime
	jsonData = `{"CreationDate": "2015-02-02T11:04:57.497"}`
	dt = testJsonModel{}
	err = json.Unmarshal([]byte(jsonData), &dt)
	if err != nil {
		t.Fatal(err)
	}
	expectedTime = time.Date(2015, time.February, 2, 11, 04, 57, 497000000, time.UTC)
	if !dt.CreationDate.Valid || dt.CreationDate.Time != expectedTime {
		t.Errorf("Expected DateTime: %v, but got: %v", expectedTime, dt.CreationDate.Time)
	}

	// Test case 5: Null DateTime
	nullJsonData := `{"CreationDate": null}`
	nullDt := testJsonModel{}
	err = json.Unmarshal([]byte(nullJsonData), &nullDt)
	expectedError = "parsing time \"null\" as \"2006-01-02T15:04:05Z07:00\": cannot parse \"null\" as \"2006\""
	if err == nil || err.Error() != expectedError {
		t.Errorf("Expected error: %s, but got: %v", expectedError, err)
	}
	if nullDt.CreationDate.Valid {
		t.Errorf("Expected DateTime to be invalid, but it is valid")
	}
}

// Test MarshalJSON
func TestDateTimeMarshalJSON(t *testing.T) {
	type testJsonModel struct {
		CreationDate DateTime `json:"CreationDate"`
	}

	// Test case 1: Valid DateTime
	dt := testJsonModel{
		CreationDate: DateTime{
			Time:  time.Date(2022, time.March, 7, 0, 0, 0, 0, time.UTC),
			Valid: true,
		},
	}
	expectedResult := `{"CreationDate":"2022-03-07T00:00:00Z"}`
	result, err := json.Marshal(dt)
	if err != nil {
		t.Fatal(err)
	}
	if string(result) != expectedResult {
		t.Errorf("Expected result: %s, but got: %s", expectedResult, result)
	}

	// Test case 2: Invalid DateTime
	invalidDt := testJsonModel{
		CreationDate: DateTime{
			Valid: false,
		},
	}
	expectedResult = `{"CreationDate":""}`
	result, err = json.Marshal(invalidDt)
	if err != nil {
		t.Fatal(err)
	}
	if string(result) != expectedResult {
		t.Errorf("Expected result: %s, but got: %s", expectedResult, result)
	}

}
