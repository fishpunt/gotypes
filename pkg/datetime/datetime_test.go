package datetime

import (
	"encoding/xml"
	"log"
	"testing"
)

func TestDateTimeParseTime(t *testing.T) {
	ts := TestStruct{}
	err := xml.Unmarshal([]byte(data), &ts)
	if err != nil {
		t.Fatal(err)
	}

	// _, err = dt.parseTime("2022-03-07T00:00:00")
	// if err != nil {
	// 	t.Fatal(err)
	// }
	log.Printf("result: %s; dateTime => %+v", ts.DateTime.String(), ts.DateTime)
}

func TestDateTimeParseEmptyTime(t *testing.T) {
	ts := TestStruct{}
	err := xml.Unmarshal([]byte(emptyData), &ts)
	if err != nil {
		t.Fatal(err)
	}

	log.Printf("result: %s; dateTime => %+v", ts.DateTime.String(), ts.DateTime)
}

var data = `<TestStruct><DateTime>2022-03-07T00:00:00</DateTime></TestStruct>`
var emptyData = `<TestStruct><DateTime></DateTime></TestStruct>`

type TestStruct struct {
	DateTime DateTime `xml:"DateTime"`
}
