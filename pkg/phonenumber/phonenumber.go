package phonenumber

import "strconv"

func IsValid(phoneNo string) bool {
	if len(phoneNo) != 11 {
		return false
	}
	if phoneNo[0:2] != "09" {
		return false
	}
	if _, error := strconv.Atoi(phoneNo[2:]); error == nil {
		return false
	}
	return true
}
