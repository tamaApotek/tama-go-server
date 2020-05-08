package datestr

import "regexp"

// DateStr represent date string YYYY-MM-DD
type DateStr string

// IsValid check whether date is valid YYYY-MM-DD
func (ds *DateStr) IsValid() bool {
	r, _ := regexp.Compile(`([12]\d{3}-(0[1-9]|1[0-2])-(0[1-9]|[12]\d|3[01]))`)
	matched := r.MatchString(string(*ds))

	return matched
}
