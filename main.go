package clientForm3Api

import (
	"fmt"
	"unsafe"
)

var a *int64

func main() {
	a := 5
	//Convert the integer to a uintptr type
	ptrVal := uintptr(a)
	//Convert the uintptr to a Pointer type
	ptr := unsafe.Pointer(ptrVal)

	//Get the string pointer by address
	stringPtr := (*string)(ptr)

	//Get the value at that pointer
	newData := *stringPtr

	//Got it:
	fmt.Println(newData)
}
