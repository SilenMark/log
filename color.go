package log

import (
	"fmt"
)

type ANSICode uint8

func (c ANSICode) Render(v string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m  ", uint8(c), v)
}
