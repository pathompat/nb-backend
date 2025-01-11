package types

import (
	"database/sql/driver"
	"fmt"
	"strings"
)

type StrArray []string

func (a *StrArray) Scan(src interface{}) error {
	if src == nil {
		*a = []string{}
		return nil
	}

	srcStr, ok := src.(string)
	if !ok {
		return fmt.Errorf("cannot convert %T to StrArray", src)
	}

	srcStr = strings.Trim(srcStr, "{}")
	if srcStr == "" {
		*a = []string{}
		return nil
	}

	elements := strings.Split(srcStr, ",")
	result := make([]string, len(elements))
	for i, elem := range elements {
		var value string
		if _, err := fmt.Sscanf(elem, "%s", &value); err != nil {
			return fmt.Errorf("error parsing element '%s': %v", elem, err)
		}
		result[i] = value
	}

	*a = result
	return nil
}

func (a StrArray) Value() (driver.Value, error) {
	if len(a) == 0 {
		return "{}", nil
	}

	elements := make([]string, len(a))
	for i, value := range a {
		elements[i] = value
	}
	return fmt.Sprintf("{%s}", strings.Join(elements, ",")), nil
}
