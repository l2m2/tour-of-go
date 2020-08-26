package main

import (
	"fmt"
	"strconv"
	"strings"
)

type IPAddr [4]byte

// TODO: 给 IPAddr 添加一个 "String() string" 方法
func (v IPAddr) String() string {
	// return fmt.Sprintf("%d.%d.%d.%d", v[0], v[1], v[2], v[3])
	s := make([]string, 0, len(v))
	for _, i := range v {
		s = append(s, strconv.Itoa(int(i)))
	}
	return strings.Join(s, ".")
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
