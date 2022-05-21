package machine_monitor

import (
	"net"
	"time"
)

func CheckPortBind(ipPort string) bool {
	conn, err := net.DialTimeout("tcp", ipPort, 3*time.Second)
	if err != nil || conn == nil {
		return false
	}

	return true
}
