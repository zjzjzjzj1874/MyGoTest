package base

import (
	"context"
	"fmt"
	"testing"
)

func TestVirtualMemoryWithContext(t *testing.T) {

	meminfo, err := VirtualMemoryWithContext(context.Background())

	fmt.Println(err)
	fmt.Println(meminfo)
}
