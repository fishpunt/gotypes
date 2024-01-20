package datetime

import (
	"sync"
	"time"
)

const (
	defaultDatetimeLayout string = time.RFC3339
)

var (
	mtx                  sync.Mutex
	datetimeLayoutOutput string
	datetimeLayoutInputs map[string]string
)

func init() {
	datetimeLayoutOutput = defaultDatetimeLayout

	datetimeLayoutInputs = make(map[string]string)
	datetimeLayoutInputs[defaultDatetimeLayout] = defaultDatetimeLayout
	datetimeLayoutInputs["2006-01-02T15:04:05"] = "2006-01-02T15:04:05"
	datetimeLayoutInputs["2006-01-02T15:04:05Z07:00"] = "2006-01-02T15:04:05Z07:00"
}

// SetOutputLayout
func SetOutputLayout(layout string) {
	mtx.Lock()
	defer mtx.Unlock()
	datetimeLayoutOutput = layout
}

// OutputLayout
func OutputLayout() string {
	return datetimeLayoutOutput
}

// SetInputLayout
func SetInputLayout(layout ...string) {
	mtx.Lock()
	defer mtx.Unlock()
	for _, v := range layout {
		datetimeLayoutInputs[v] = v
	}
}

// InputLayout
func InputLayout() []string {
	var layouts []string
	for k := range datetimeLayoutInputs {
		layouts = append(layouts, k)
	}
	return layouts
}
