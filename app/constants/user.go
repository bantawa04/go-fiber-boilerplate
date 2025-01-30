package constants

import (
	"fmt"
)

type Gender string

const (
	Male   Gender = "male"
	Female Gender = "female"
	Other  Gender = "other"
)

func (r Gender) IsValidVal(val string) error {
	switch val {
	case "male":
		return nil
	case "female":
		return nil
	case "other":
		return nil
	default:
		return fmt.Errorf("invalid user relation status: %q", val)
	}
}
