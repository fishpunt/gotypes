package date

import "sync"

const (
	defaultDateLayout string = "2006-01-02"
)

var (
	mtx              sync.Mutex
	dateLayoutOutput *string
	dateLayoutInput  *string
)

func init() {
	defaultDateOutput := defaultDateLayout
	dateLayoutOutput = &defaultDateOutput

	defaultDateInput := defaultDateLayout
	dateLayoutInput = &defaultDateInput
}

// SetOutputLayout
func SetOutputLayout(layout string) {
	mtx.Lock()
	defer mtx.Unlock()
	dateLayoutOutput = &layout
}

// OutputLayout
func OutputLayout() string {
	return *dateLayoutOutput
}

// SetInputLayout
func SetInputLayout(layout string) {
	mtx.Lock()
	defer mtx.Unlock()
	dateLayoutInput = &layout
}

// InputLayout
func InputLayout() string {
	return *dateLayoutInput
}
