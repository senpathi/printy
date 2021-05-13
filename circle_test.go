package printy

import "testing"

func TestPrintInCircle(t *testing.T) {
	sampleStr := `This is a sample string, again this is a sample string `
	PrintInCircle(sampleStr)
}

func TestRadius(t *testing.T) {
	samples := []int{16, 18, 51, 53, 158, 159, 236, 238, 328, 329, 442, 444, 570, 571, 717, 719, 872, 873}
	radius := []int{1, 2, 2, 3, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9, 10, 10, 11}

	for i, sample := range samples {
		r := dd(sample)
		if r < radius[i] {
			t.Errorf(`radius expected %d, received %d for area %d`, radius[i], r, sample)
		}
	}
}
