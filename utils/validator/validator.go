package validator

import (
	"regexp"
)

func CheckPhoneNumber(number string) bool {
	str := `^((8|\+7)[\- ]?)?(\(?\d{3}\)?[\- ]?)?[\d\- ]{7,10}$`
	re := regexp.MustCompile(str)

	return re.MatchString(number)
}
