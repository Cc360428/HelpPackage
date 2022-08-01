package eureka_client

import (
	"fmt"
	"testing"
)

func TestGetLocalIP(t *testing.T) {
	fmt.Println(GetLocalIP())
}
