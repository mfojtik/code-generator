package util

import (
	"fmt"
	"strings"
)

func PluralExceptionListToMapOrDie(pluralExceptions []string) map[string]string {
	pluralExceptionMap := make(map[string]string, len(pluralExceptions))
	for i := range pluralExceptions {
		parts := strings.Split(pluralExceptions[i], ":")
		if len(parts) != 2 {
			panic(fmt.Sprintf("invalid plural exception definition: %s", pluralExceptions[i]))
		}
		pluralExceptionMap[parts[0]] = parts[1]
	}
	return pluralExceptionMap
}
