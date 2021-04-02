package guid

import (
	"strings"

	"github.com/google/uuid"
)

func Gen() string {
	return strings.ReplaceAll(uuid.NewString(), "-", "")
}
