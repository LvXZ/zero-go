package util

import (
	"fmt"
	"testing"
)

func TestGetSign(t *testing.T) {
	sign := DefaultGetSign("68ac168263931e1b314fca91917a80b7", "40dbaaf3aa8b0e18b7afac1e0f0b5be4")
	fmt.Println("sign = " + sign)
}
