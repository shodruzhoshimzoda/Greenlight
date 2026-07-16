package data

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Runtime int

var InvalidRuntimeFormat = errors.New("invalid runtime format")

// This method make raadable and customizable runtime field. Example: 120 mins
func (r Runtime) MarshalJSON() ([]byte, error) {

	jsonValue := fmt.Sprintf("%d mins", r)

	quotedJSONValues := strconv.Quote(jsonValue)

	return []byte(quotedJSONValues), nil

}

func (r *Runtime) UnmarshalJSON(b []byte) error {
	unQuotedJSONValues, err := strconv.Unquote(string(b))
	if err != nil {
		return InvalidRuntimeFormat
	}
	parts := strings.Split(unQuotedJSONValues, " ")
	if len(parts) != 2 || parts[1] != "mins" {
		return InvalidRuntimeFormat
	}

	i, err := strconv.ParseInt(parts[0], 10, 32)
	if err != nil {
		return InvalidRuntimeFormat
	}

	*r = Runtime(i)

	return nil
}
