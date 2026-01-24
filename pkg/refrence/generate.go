package refrence

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

func GenerateReference(prefix string) string {
	return fmt.Sprintf(
		"%s-%s-%s",
		prefix,
		time.Now().Format("20060102"),
		uuid.NewString(),
	)
}
