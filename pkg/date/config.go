package date

import (
	"sync"
	"time"
)

const (
	defaultDateLayout string = "2006-01-02"
)

var (
	mtx              sync.Mutex
	dateLayoutOutput *string
	dateLayoutInputs map[string]string
)

func init() {
	defaultDateOutput := defaultDateLayout
	dateLayoutOutput = &defaultDateOutput

	dateLayoutInputs = make(map[string]string)
	dateLayoutInputs[defaultDateLayout] = defaultDateLayout
	dateLayoutInputs["2006-01-02T15:04:05"] = "2006-01-02T15:04:05"
	dateLayoutInputs["2006-01-02T15:04:05Z07:00"] = "2006-01-02T15:04:05Z07:00"
	dateLayoutInputs["2006-01-02T15:04:05-07:00"] = "2006-01-02T15:04:05-07:00"
	dateLayoutInputs[time.RFC3339] = time.RFC3339
}

// SetOutputLayout
func SetOutputLayout(layout string) {
	mtx.Lock()
	defer mtx.Unlock()
	*dateLayoutOutput = layout
}

// OutputLayout
func OutputLayout() string {
	return *dateLayoutOutput
}

// SetInputLayout
func SetInputLayout(layout ...string) {
	mtx.Lock()
	defer mtx.Unlock()
	for _, v := range layout {
		dateLayoutInputs[v] = v
	}
}

// InputLayout
func InputLayout() []string {
	var layouts []string
	for k := range dateLayoutInputs {
		layouts = append(layouts, k)
	}
	return layouts
}
