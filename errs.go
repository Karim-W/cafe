package cafe

import "fmt"

var (
	Err_REQUIRED_KEY_MISSING = fmt.Errorf("required key missing")
	Err_KEY_IS_REQUIRED      = fmt.Errorf("key is required")
)

const (
	str_UNREGISTERED      string = " is not a registered key"
	str_NON_MATCHED_FETCH string = "%s is registered as a %s, but you are trying to fetch it as a %s"
)

func buildRequiredKeyMissing(key string) error {
	return fmt.Errorf("required key %s missing", key)
}
