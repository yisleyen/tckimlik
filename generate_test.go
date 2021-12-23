package tckimlik

import "testing"

func TestGenerate(t *testing.T) {
	tests := []struct {
		name string
		data string
	}{
		{
			name: "Generate TC Identity and Validate",
			data: Generate(),
		},
	}
	for _, test := range tests {
		valid, err := Validate(test.data)
		if !valid || err != nil {
			t.Errorf("Generate() = %v, want %v", valid, test.data)
		}
	}
}
