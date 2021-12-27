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
		isValid, err := Validate(test.data)
		if !isValid || err != nil {
			t.Errorf("Generate() = %v, %v", isValid, test.data)
		}
	}
}
