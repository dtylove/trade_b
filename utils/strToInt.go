package utils

import "strconv"

func StrToUint(strNumber string, value interface{}) (err error) {
	var number interface{}
	number, err = strconv.ParseUint(strNumber, 10, 64)

	switch v := number.(type) {
	case uint64:
		switch d := value.(type) {
		case *uint64:
			*d = v
		case *uint:
			*d = uint(v)
		case *uint16:
			*d = uint16(v)
		case *uint32:
			*d = uint32(v)
		case *uint8:
			*d = uint8(v)
		}
	}
	return
}
