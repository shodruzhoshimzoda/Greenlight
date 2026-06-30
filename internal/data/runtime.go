package data

import (
	"fmt"
	"strconv"
)





type Runtime int


// This method make raadable and customizable runtime field. Example: 120 mins
func (r Runtime) MarshalJSON() ([]byte, error) {

	jsonValue := fmt.Sprintf("%d mins", r)

	quotedJSONValues := strconv.Quote(jsonValue)


	return []byte(quotedJSONValues), nil


}