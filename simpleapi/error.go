package simpleapi

import "encoding/json"

//Error custom errors detail
type Error struct {
	Detail interface{}
}

//Error must implement this Error func
func (e *Error) Error() string {
	dat, _ := json.Marshal(e.Detail)
	return string(dat)
}
