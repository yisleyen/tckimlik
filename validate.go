package tckimlik

import "errors"

func Validate(tcno string) (valid bool, err error) {
	if len(tcno)!=11 {
		err = errors.New("Invalid length")
		return
	}

	if tcno[0] == '0' {
		err = errors.New("Invalid TC Identity Number")
		return
	}

	return true, nil
}
