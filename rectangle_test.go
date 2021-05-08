package printy

import "testing"

func TestPrintInRectangle(t *testing.T) {
	sampleStr := `This is a sample string, again this is a sample string `
	PrintInRectangle(sampleStr, 5, 3)
}
