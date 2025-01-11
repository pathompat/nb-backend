package types

import (
	"database/sql/driver"
	"fmt"
	"strings"
)

type IntArray []int

func (a *IntArray) Scan(src interface{}) error {
	if src == nil {
		*a = []int{}
		return nil
	}

	srcStr, ok := src.(string)
	if !ok {
		return fmt.Errorf("cannot convert %T to IntArray", src)
	}

	srcStr = strings.Trim(srcStr, "{}")
	if srcStr == "" {
		*a = []int{}
		return nil
	}

	elements := strings.Split(srcStr, ",")
	result := make([]int, len(elements))
	for i, elem := range elements {
		var value int
		if _, err := fmt.Sscanf(elem, "%d", &value); err != nil {
			return fmt.Errorf("error parsing element '%s': %v", elem, err)
		}
		result[i] = value
	}

	*a = result
	return nil
}

func (a IntArray) Value() (driver.Value, error) {
	if len(a) == 0 {
		return "{}", nil
	}

	elements := make([]string, len(a))
	for i, value := range a {
		elements[i] = fmt.Sprintf("%d", value)
	}
	return fmt.Sprintf("{%s}", strings.Join(elements, ",")), nil
}
