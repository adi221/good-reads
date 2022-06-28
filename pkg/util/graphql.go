package util

import "strconv"

// ConvGQLParamaterToUint converts ID GraphQL empty interface object to unsigned integer
func ConvGQLParamaterToUint(t interface{}) (uint, bool) {
	switch t := t.(type) {
	case int:
		return uint(t), true
	case string:
		ui64, err := strconv.ParseUint(t, 10, 32)
		if err != nil {
			return 0, false
		}
		return uint(ui64), true
	default:
		return 0, false
	}
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
