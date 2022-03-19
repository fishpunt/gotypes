package datetime

import "testing"

func TestDateTimeParseTime(t *testing.T) {
	_, err := ParseTime("2022-03-07T00:00:00")
	if err != nil {
		t.Fatal(err)
	}
}
