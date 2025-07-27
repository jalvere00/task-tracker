package cli

import (
	"strconv"
)

func parseTaskID(arg string) (int, error) {
	return strconv.Atoi(arg)
}
