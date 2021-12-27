package tckimlik

import "testing"

func TestValidate(t *testing.T) {
	tests := []struct {
		name string
		tckimlik string
		isValid bool
	}{
		{
			name: "Valid",
			tckimlik: "18888554478",
			isValid: true,
		},
		{
			name: "Invalid length",
			tckimlik: "188885544785",
			isValid: false,
		},
		{
			name: "Invalid character",
			tckimlik: "1888855447A",
			isValid: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			isValid, err := Validate(tc.tckimlik)
			if (err != nil) {
				t.Errorf("Validate() error = %v", err)
				return
			}
			if isValid != tc.isValid {
				t.Errorf("Validate() isValid = %v", isValid)
			}
		})
	}
}
