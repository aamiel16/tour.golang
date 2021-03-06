package main

import "fmt"

type IPAddr [4]byte

/*
* - To string function (equivalent to .toString() method in JS)
* - Must satisfy the built-in interface:
*     type Stringer interface {
*       String() string
*     } 
*/
func (ip IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
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
