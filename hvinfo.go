package main

import (
	"fmt"
	"github.com/fromanirh/cpuid"
)

func main() {
	fmt.Printf("VendorString=%s\n", cpuid.VendorIdentificatorString)
	fmt.Printf("HyperVSupport=%v\n", cpuid.HasHypervSupport())
}
