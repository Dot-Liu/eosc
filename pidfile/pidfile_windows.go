package pidfile // import "github.com/docker/docker/pkg/pidfile"

import (
	"golang.org/x/sys/windows"
)

const (
	processQueryLimitedInformation = 0x1000

	stillActive = 259
)

func ProcessExists(pid int) bool {
	h, err := windows.OpenProcess(processQueryLimitedInformation, false, uint32(pid))
	if err != nil {
		return false
	}
	var c uint32
	err = windows.GetExitCodeProcess(h, &c)
	windows.Close(h)
	if err != nil {
		return c == stillActive
	}
	return true
}
