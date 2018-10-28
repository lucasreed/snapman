package tags

import (
	"errors"
	"regexp"
)

// Validate is used to make sure the tags passed in as an argument are valid
func Validate(t *string) error {
	tag := *t
	validPre := regexp.MustCompile(`^(aws:|AWS:)`)
	validTag := regexp.MustCompile(`^([a-zA-Z0-9+-=._:/]{1,128})=([a-zA-Z0-9+-=._:/]{1,256})$`)

	if !validPre.MatchString(tag[0:4]) && validTag.MatchString(tag) {
		return nil
	}
	return errors.New("Not a valid tag key=value combination")
}
