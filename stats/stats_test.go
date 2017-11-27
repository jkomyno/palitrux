package stats

import "testing"

func TestToMegaBytes(t *testing.T) {
	tests := []struct {
		value    uint64
		expected float64
	}{
		{1024, 0},
		{1024 * 1024, 1},
		{1024 * 1024 * 10, 10},
		{1024 * 1024 * 100, 100},
		{1024 * 1024 * 250, 250},
	}

	for _, test := range tests {
		val := toMegaBytes(test.value)
		if val != test.expected {
			t.Errorf("Invalid param: %#v != %#v", val, test.expected)
		}
	}
}
