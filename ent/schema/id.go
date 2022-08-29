package schema

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
)

func GenerateID(prefix string) func() string {
	return func() string {
		return fmt.Sprintf("%s_%s", prefix, strings.ReplaceAll(uuid.New().String(), "-", ""))
	}
}
