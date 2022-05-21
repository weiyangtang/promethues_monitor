package machine_monitor

import (
	"fmt"
	"testing"
)

func TestCheckPortBind(t *testing.T) {
	fmt.Println(CheckPortBind("localhost:6379"))
}
