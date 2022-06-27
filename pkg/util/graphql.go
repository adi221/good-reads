package util

import "strconv"

// ConvGQLStringToUint converts ID GraphQL interface object to unsigned integer
func ConvGQLStringToUint(id interface{}) (uint, bool) {
	str, ok := id.(string)
	if !ok {
		return 0, ok
	}
	ui64, err := strconv.ParseUint(str, 10, 32)
	if err != nil {
		return 0, false
	}
	return uint(ui64), true
}

// GetGQLUintParameter converts graphql paramater (int/string) type into uint pointer
func GetGQLUintParameter(value interface{}) *uint {
	switch v := value.(type) {
	case int:
		i := uint(v)
		return &i
	case string:
		ui64, err := strconv.ParseUint(v, 10, 32)
		if err != nil {
			return nil
		}
		s := uint(ui64)
		return &s
	default:
		return nil
	}
}

// GetGQLStringParameter converts GraphQL parameter to string pointer
func GetGQLStringParameter(value interface{}) *string {
	switch value.(type) {
	case string:
		str := value.(string)
		return &str
	default:
		return nil
	}
}
